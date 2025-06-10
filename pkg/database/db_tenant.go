package database

import (
	"fmt"
	// "log"
	"os"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	// "github.com/google/uuid"

	"gorm.io/driver/sqlite"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQLITE

func PrepareDB(uri string, userID string) error {
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