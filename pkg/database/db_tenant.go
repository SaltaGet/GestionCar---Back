package database

import (
	"fmt"
	// "log"
	"os"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	// "github.com/google/uuid"

	"gorm.io/driver/sqlite"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQLITE

func PrepareDB(uri string) error {
	// Si el archivo ya existe, no hacer nada
	filePath := filePathFromURI(uri)
	if _, err := os.Stat(filePath); err == nil {
		return nil // Ya existe, no hacer nada
	}

	// Crear la base de datos
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error inicializando db: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("no se pudo obtener la conexión de bajo nivel: %w", err)
	}
	defer sqlDB.Close()

	// Migrar las tablas necesarias
	if err := db.AutoMigrate(
		&models.Attendance{},
		&models.Client{},
		&models.Employee{},
		&models.Expense{},
		&models.Income{},
		&models.Member{},
		&models.MovementType{},
		&models.Permission{},
		&models.Product{},
		&models.PurchaseOrder{},
		&models.PurchaseProduct{},
		&models.ResumeExpense{},
		&models.ResumeIncome{},
		&models.Role{},
		&models.Service{},
		&models.Supplier{},
		&models.Vehicle{},
		// Agrega aquí tus modelos tenant-específicos
	); err != nil {
		_ = os.Remove(filePath)
		return fmt.Errorf("error al migrar tablas: %w", err)
	}

	// var count int64
	// db.Model(&models.Role{}).Count(&count)
	// var role models.Role
	// if count == 0 {
	// 	role = models.Role{ID: uuid.NewString(), Name: "admin"}
	// 	db.Create(&role)
	// } else {
	// 	db.Where("name = ?", "admin").First(&role)
	// }

	// db.Model(&models.Member{}).Count(&count)
	// if count == 0 {
	// 	db.Create(&models.Member{
	// 		ID: uuid.NewString(),
	// 		UserID: userID,
	// 		RoleID: role.ID,
	// 	})
	// }

	db.Create(&permissions)

	return nil
}

// // MYSQL
// func PrepareDB(dsn string, userID string) error {
// 	err := EnsureDatabaseExists(dsn)
// 	if err != nil {
// 		return fmt.Errorf("error al crear la db: %w", err)
// 	}

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return fmt.Errorf("error inicializando db: %w", err)
// 	}
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return fmt.Errorf("no se pudo obtener la conexión de bajo nivel: %w", err)
// 	}
// 	defer sqlDB.Close()

// 	if err := db.AutoMigrate(
// 		&models.Attendance{},
// 		&models.Client{},
// 		&models.Employee{},
// 		&models.Expense{},
// 		&models.Income{},
// 		&models.Member{},
// 		&models.MovementType{},
// 		&models.Permission{},
// 		&models.Product{},
// 		&models.PurchaseOrder{},
// 		&models.PurchaseProduct{},
// 		&models.ResumeExpense{},
// 		&models.ResumeIncome{},
// 		&models.Role{},
// 		&models.Service{},
// 		&models.Supplier{},
// 		&models.Vehicle{},
// 	); err != nil {
// 		return fmt.Errorf("error al migrar tablas: %w", err)
// 	}

// 	var count int64
// 	db.Model(&models.Role{}).Count(&count)

// 	var role models.Role
// 	if count == 0 {
// 		role = models.Role{ID: uuid.NewString(), Name: "admin"}
// 		db.Create(&role)
// 	} else {
// 		db.Where("name = ?", "admin").First(&role)
// 	}

// 	db.Model(&models.Member{}).Count(&count)
// 	if count == 0 {
// 		db.Create(&models.Member{
// 			ID:     uuid.NewString(),
// 			UserID: userID,
// 			RoleID: role.ID,
// 		})
// 	}

// 	return nil
// }


var permissions = []models.Permission{
	{ID: uuid.NewString(), Code: "VEV", Details: "Ver ventas", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VEA", Details: "Agregar venta", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VEE", Details: "Editar venta", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VED", Details: "Eliminar venta", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VEX", Details: "Acceso a Caja Rápida", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VAP", Details: "Agregar pago", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VEP", Details: "Editar pago", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VDP", Details: "Eliminar pago", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VMP", Details: "Modificar precios", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VMD", Details: "Modificar descuentos", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VAD", Details: "Agregar descuento", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VED2", Details: "Editar descuento", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VDD", Details: "Eliminar descuento", Group: "Ventas"},
	{ID: uuid.NewString(), Code: "VSC", Details: "Enviar email a cliente", Group: "Ventas"},

	{ID: uuid.NewString(), Code: "COV", Details: "Ver compras", Group: "Compras"},
	{ID: uuid.NewString(), Code: "COA", Details: "Agregar compra", Group: "Compras"},
	{ID: uuid.NewString(), Code: "COE", Details: "Editar compra", Group: "Compras"},
	{ID: uuid.NewString(), Code: "COD", Details: "Eliminar compra", Group: "Compras"},
	{ID: uuid.NewString(), Code: "AOV", Details: "Ver ajuste stock", Group: "Compras"},
	{ID: uuid.NewString(), Code: "AOA", Details: "Agregar ajuste", Group: "Compras"},
	{ID: uuid.NewString(), Code: "AOE", Details: "Editar ajuste", Group: "Compras"},
	{ID: uuid.NewString(), Code: "AOD", Details: "Eliminar ajuste", Group: "Compras"},
	{ID: uuid.NewString(), Code: "CAP", Details: "Agregar pago proveedor", Group: "Compras"},
	{ID: uuid.NewString(), Code: "CEP", Details: "Editar pago proveedor", Group: "Compras"},
	{ID: uuid.NewString(), Code: "CDP", Details: "Eliminar pago proveedor", Group: "Compras"},
	{ID: uuid.NewString(), Code: "CSP", Details: "Enviar email proveedor", Group: "Compras"},

	{ID: uuid.NewString(), Code: "EGA", Details: "Agregar egreso", Group: "Egresos/Ingresos"},
	{ID: uuid.NewString(), Code: "EGE", Details: "Editar egreso", Group: "Egresos/Ingresos"},
	{ID: uuid.NewString(), Code: "EGD", Details: "Eliminar egreso", Group: "Egresos/Ingresos"},
	{ID: uuid.NewString(), Code: "IGA", Details: "Agregar ingreso", Group: "Egresos/Ingresos"},
	{ID: uuid.NewString(), Code: "IGE", Details: "Editar ingreso", Group: "Egresos/Ingresos"},
	{ID: uuid.NewString(), Code: "IGD", Details: "Eliminar Ingreso", Group: "Egresos/Ingresos"},

	{ID: uuid.NewString(), Code: "CLV", Details: "Ver cliente", Group: "Clientes"},
	{ID: uuid.NewString(), Code: "CLA", Details: "Agregar cliente", Group: "Clientes"},
	{ID: uuid.NewString(), Code: "CLE", Details: "Editar cliente", Group: "Clientes"},
	{ID: uuid.NewString(), Code: "CLD", Details: "Eliminar cliente", Group: "Clientes"},

	{ID: uuid.NewString(), Code: "PRV", Details: "Ver proveedor", Group: "Proveedores"},
	{ID: uuid.NewString(), Code: "PRA", Details: "Agregar proveedor", Group: "Proveedores"},
	{ID: uuid.NewString(), Code: "PRE", Details: "Editar proveedor", Group: "Proveedores"},
	{ID: uuid.NewString(), Code: "PRD", Details: "Eliminar proveedor", Group: "Proveedores"},

	{ID: uuid.NewString(), Code: "PDV", Details: "Ver producto", Group: "Productos"},
	{ID: uuid.NewString(), Code: "PDA", Details: "Agregar producto", Group: "Productos"},
	{ID: uuid.NewString(), Code: "PDE", Details: "Editar producto", Group: "Productos"},
	{ID: uuid.NewString(), Code: "PDD", Details: "Eliminar producto", Group: "Productos"},
	{ID: uuid.NewString(), Code: "PDS", Details: "Añadir stock inicial", Group: "Productos"},
	{ID: uuid.NewString(), Code: "PDP", Details: "Ver precio compra", Group: "Productos"},
	{ID: uuid.NewString(), Code: "PDI", Details: "Editar IVA", Group: "Productos"},

	{ID: uuid.NewString(), Code: "ARV", Details: "Ver factura ARCA", Group: "ARCA"},
	{ID: uuid.NewString(), Code: "ARC", Details: "Crear factura ARCA", Group: "ARCA"},

	{ID: uuid.NewString(), Code: "CTX", Details: "Acceso módulo contable", Group: "Cuentas"},

	{ID: uuid.NewString(), Code: "INP", Details: "Reporte pagos/cobros", Group: "Informes"},
	{ID: uuid.NewString(), Code: "INC", Details: "Cuenta corriente", Group: "Informes"},
	{ID: uuid.NewString(), Code: "INI", Details: "Ingresos", Group: "Informes"},
	{ID: uuid.NewString(), Code: "INE", Details: "Egresos", Group: "Informes"},
	{ID: uuid.NewString(), Code: "INS", Details: "Stock", Group: "Informes"},
	{ID: uuid.NewString(), Code: "INR", Details: "Ranking", Group: "Informes"},
	{ID: uuid.NewString(), Code: "ING", Details: "Generales (ventas, compras, cajas)", Group: "Informes"},

	{ID: uuid.NewString(), Code: "USV", Details: "Ver usuarios", Group: "Usuarios"},
	{ID: uuid.NewString(), Code: "USA", Details: "Agregar usuarios", Group: "Usuarios"},
	{ID: uuid.NewString(), Code: "USE", Details: "Editar usuarios", Group: "Usuarios"},
	{ID: uuid.NewString(), Code: "USD", Details: "Eliminar usuarios", Group: "Usuarios"},

	{ID: uuid.NewString(), Code: "PAR", Details: "Ver resumen ejecutivo", Group: "Panel"},
	{ID: uuid.NewString(), Code: "PAV", Details: "Ver ventas en dashboard", Group: "Panel"},

	{ID: uuid.NewString(), Code: "RLV", Details: "Ver roles", Group: "Roles"},
	{ID: uuid.NewString(), Code: "RLA", Details: "Agregar roles", Group: "Roles"},
	{ID: uuid.NewString(), Code: "RLE", Details: "Editar roles", Group: "Roles"},
	{ID: uuid.NewString(), Code: "RLD", Details: "Eliminar roles", Group: "Roles"},

	{ID: uuid.NewString(), Code: "MCV", Details: "Ver marcas/categorías", Group: "Marcas/Categorías"},
	{ID: uuid.NewString(), Code: "MCA", Details: "Agregar marcas/categorías", Group: "Marcas/Categorías"},
	{ID: uuid.NewString(), Code: "MCE", Details: "Editar marcas/categorías", Group: "Marcas/Categorías"},
	{ID: uuid.NewString(), Code: "MCD", Details: "Eliminar marcas/categorías", Group: "Marcas/Categorías"},

	{ID: uuid.NewString(), Code: "CXX", Details: "Arqueo de caja", Group: "Caja Rápida"},
}
