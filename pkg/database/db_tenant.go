package database

import (
	// "log"
	"os"
	"strings"
	
	"fmt"
	
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	
	// "github.com/google/uuid"
	
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func PrepareDB(uri string) error {
	env := os.Getenv("ENV")

	var driver gorm.Dialector
	if env == "prod" {
		// MYSQL
		if err := ensureDatabaseExists(uri); err != nil {
			return fmt.Errorf("error al crear la base: %w", err)
		}
		driver = mysql.Open(uri)
	} else {
		// SQLITE
		filePath := filePathFromURI(uri)
		if _, err := os.Stat(filePath); err == nil {
			return nil
		}
		driver = sqlite.Open(uri)
	}

	// Conexión GORM
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		handleDBCreationError(env, uri)
		return fmt.Errorf("error al inicializar DB: %w", err)
	}

	// Bajo nivel
	sqlDB, err := db.DB()
	if err != nil {
		handleDBCreationError(env, uri)
		return fmt.Errorf("error al obtener conexión de bajo nivel: %w", err)
	}
	defer sqlDB.Close()

	// Migraciones
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
	); err != nil {
		handleDBCreationError(env, uri)
		return fmt.Errorf("error al migrar tablas: %w", err)
	}

	if err := db.Create(&permissions).Error; err != nil {
		handleDBCreationError(env, uri)
		return fmt.Errorf("error al migrar permisos: %w", err)
	}

	return nil
}

func handleDBCreationError(env, uri string) {
	if env == "prod" {
		_ = dropDatabase(uri)
	} else {
		_ = os.Remove(filePathFromURI(uri))
	}
}

// // SQLITE

// func PrepareDB(uri string) error {
// 	filePath := filePathFromURI(uri)
// 	if _, err := os.Stat(filePath); err == nil {
// 		return nil
// 	}

// 	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
// 	if err != nil {
// 		return models.ErrorResponse(500, "Error interno al crear la base de datos", err)
// 	}
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return models.ErrorResponse(500, "Error interno no se pudo obtener la conexión de bajo nivel", err)
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
// 		_ = os.Remove(filePath)
// 		return models.ErrorResponse(500, "Error interno al migrar la base de datos", err)
// 	}

// 	if err := db.Create(&permissions).Error; err != nil {
// 		_ = os.Remove(filePath)
// 		return models.ErrorResponse(500, "Error interno al migrar permisos base de datos", err)
// 	}

// 	return nil
// }

// // MYSQL
// func PrepareDB(dsn string) error {
// 	err := EnsureDatabaseExists(dsn)
// 	if err != nil {
// 		return fmt.Errorf("error al crear la db: %w", err)
// 	}

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		dropErr := dropDatabase(dsn)
// 		if dropErr != nil {
// 			return fmt.Errorf("error al crear permisos: %v; además falló al borrar la base: %v", err, dropErr)
// 		}
// 		return fmt.Errorf("error inicializando db: %w", err)
// 	}
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		dropErr := dropDatabase(dsn)
// 		if dropErr != nil {
// 			return fmt.Errorf("error al crear permisos: %v; además falló al borrar la base: %v", err, dropErr)
// 		}
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
// 		dropErr := dropDatabase(dsn)
// 		if dropErr != nil {
// 			return fmt.Errorf("error al crear permisos: %v; además falló al borrar la base: %v", err, dropErr)
// 		}
// 		return fmt.Errorf("error al migrar tablas: %w", err)
// 	}

// 	if err := db.Create(&permissions).Error; err != nil {
// 		dropErr := dropDatabase(dsn)
// 		if dropErr != nil {
// 			return fmt.Errorf("error al crear permisos: %v; además falló al borrar la base: %v", err, dropErr)
// 		}
// 		return models.ErrorResponse(500, "Error interno al migrar permisos base de datos", err)
// 	}
	
// 	return nil
// }

func dropDatabase(dsn string) error {
	dbName, err := extractDBName(dsn)
	if err != nil {
		return fmt.Errorf("no se pudo extraer el nombre de la base: %w", err)
	}

	baseDSN := removeDBFromDSN(dsn)
	sqlDB, err := sql.Open("mysql", baseDSN)
	if err != nil {
		return fmt.Errorf("no se pudo conectar al servidor MySQL: %w", err)
	}
	defer sqlDB.Close()

	_, err = sqlDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", dbName))
	if err != nil {
		return fmt.Errorf("error al ejecutar DROP DATABASE: %w", err)
	}
	return nil
}

func extractDBName(dsn string) (string, error) {
	beforeParams := strings.SplitN(dsn, "?", 2)[0]
	parts := strings.Split(beforeParams, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("formato de DSN inválido (no se encontró la base)")
	}
	return parts[1], nil
}

func removeDBFromDSN(dsn string) string {
	i := strings.Index(dsn, "/")
	if i == -1 {
		return dsn
	}

	paramStart := strings.Index(dsn[i:], "?")
	if paramStart != -1 {
		return dsn[:i+1] + dsn[i+paramStart:]
	}
	return dsn[:i+1]
}

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
