package database

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
	"github.com/google/uuid"

	//SQLITE
	// "gorm.io/driver/sqlite"

	//MYSQL
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"

	"github.com/hashicorp/golang-lru"
	"gorm.io/gorm"
)

// // SQLITE

// var (
// 	mainDB       *gorm.DB
// 	tenantDBs    *lru.Cache
// 	mu           sync.RWMutex
// 	dbExpiration = 30 * time.Minute
// )

// func ConnectDB(uri string) (*gorm.DB, error) {
// 	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDB.SetMaxOpenConns(50)
// 	sqlDB.SetMaxIdleConns(25)
// 	sqlDB.SetConnMaxLifetime(3 * time.Hour)
// 	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

// 	db.AutoMigrate(
// 		&models.User{},
// 		&models.Tenant{},
// 		&models.UserTenant{},
// 	)

// 	var email string
// 	db.Model(&models.User{}).Select("email").Where("email = ?", os.Getenv("ADMIN_EMAIL")).Scan(&email)

// 	if email != "" {
// 		log.Println("El admin ya existe")
// 		mainDB = db
// 		return db, nil
// 	}
// 	newId := uuid.NewString()

// 	pass, err := utils.HashPassword(os.Getenv("ADMIN_PASSWORD"))

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := db.Create(&models.User{
// 		ID:        newId,
// 		FirstName: os.Getenv("FIRSTNAME_ADMIN"),
// 		LastName:  os.Getenv("LASTNAME_ADMIN"),
// 		Username:  os.Getenv("ADMIN_USERNAME"),
// 		Email:     os.Getenv("ADMIN_EMAIL"),
// 		Password:  pass,
// 		IsAdmin:   true,
// 	}).Error; err != nil {
// 		return nil, err
// 	}

// 	mainDB = db

// 	return db, nil
// }

// func GetMainDB() *gorm.DB {
// 	return mainDB
// }

// func InitDBCache(maxEntries int) {
// 	var err error
// 	tenantDBs, err = lru.New(maxEntries)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func GetTenantDB(uri string) (*gorm.DB, error) {
// 	filePath := filePathFromURI(uri)
// 	if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 		return nil, fmt.Errorf("la base de datos del tenant no existe: %s", uri)
// 	}

// 	mu.RLock()
//     if val, ok := tenantDBs.Get(uri); ok {
//         entry := val.(*tenantDBEntry)
//         entry.lastUsed = time.Now()
//         mu.RUnlock()
//         return entry.db, nil
//     }
//     mu.RUnlock()

// 	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlDB.SetMaxOpenConns(20)
// 	sqlDB.SetMaxIdleConns(5)
// 	sqlDB.SetConnMaxLifetime(1 * time.Hour)
// 	sqlDB.SetConnMaxIdleTime(15 * time.Minute)

// 	entry := &tenantDBEntry{
// 		db:       db,
// 		lastUsed: time.Now(),
// 	}

// 	mu.Lock()
// 	tenantDBs.Add(uri, entry)
// 	mu.Unlock()

// 	return db, nil
// }

// type tenantDBEntry struct {
//     db        *gorm.DB
//     lastUsed  time.Time
//     // mu        sync.Mutex
//     // inUse     int           // Contador de usos
//     // closed    bool
// }

// func StartDBJanitor() {
// 	ticker := time.NewTicker(1 * time.Hour)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		mu.Lock()
// 		for _, key := range tenantDBs.Keys() {
// 			if val, ok := tenantDBs.Get(key); ok {
// 				entry := val.(*tenantDBEntry)
// 				if time.Since(entry.lastUsed) > dbExpiration {
// 					if db, err := entry.db.DB(); err == nil {
// 						db.Close()
// 					}
// 					tenantDBs.Remove(key)
// 				}
// 			}
// 		}
// 		mu.Unlock()
// 	}
// }

// func CloseDB(db *gorm.DB) error {
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatal("No se pudo obtener la conexión de bajo nivel:", err)
// 	}

// 	if sqlDB != nil {
// 		if err := sqlDB.Close(); err != nil {
// 			log.Fatal("Error al cerrar la conexión:", err)
// 		}
// 	}
// 	return nil
// }

// func CloseAllTenantDBs() error {
// 	for _, key := range tenantDBs.Keys() {
// 		if val, ok := tenantDBs.Get(key); ok {
// 			entry := val.(*tenantDBEntry)
// 			if db, err := entry.db.DB(); err == nil {
// 				db.Close()
// 			}
// 			tenantDBs.Remove(key)
// 		}
// 	}
// 	return nil
// }

// func filePathFromURI(uri string) string {
// 	uri = strings.TrimPrefix(uri, "file:")
// 	if idx := strings.Index(uri, "?"); idx != -1 {
// 		uri = uri[:idx]
// 	}
// 	return uri
// }

// MYSQL

var (
	mainDB       *gorm.DB
	tenantDBs    *lru.Cache
	mu           sync.RWMutex
	dbExpiration = 30 * time.Minute
)

type tenantDBEntry struct {
    db        *gorm.DB
    lastUsed  time.Time
    // mu        sync.Mutex
    // inUse     int           // Contador de usos
    // closed    bool
}

func ConnectDB(dsn string) (*gorm.DB, error) {
	err := EnsureDatabaseExists(dsn)
	if err != nil {
		log.Fatalf("No se pudo crear la base: %v", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(3 * time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	if err := db.AutoMigrate(
		&models.Tenant{},
		&models.User{},
		&models.UserTenant{},
	); err != nil {
		log.Fatalf("Error en migración: %v", err)
	}

	var email string
	db.Model(&models.User{}).Select("email").Where("email = ?", os.Getenv("ADMIN_EMAIL")).Scan(&email)

	if email != "" {
		log.Println("El admin ya existe")
		mainDB = db
		return db, nil
	}
	newId := uuid.NewString()

	pass, err := utils.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		return nil, err
	}

	db.Create(&models.User{
		ID:        newId,
		FirstName: os.Getenv("FIRSTNAME_ADMIN"),
		LastName:  os.Getenv("LASTNAME_ADMIN"),
		Username:  os.Getenv("ADMIN_USERNAME"),
		Email:     os.Getenv("ADMIN_EMAIL"),
		Password:  pass,
		IsAdmin:   true,
	})

	mainDB = db
	return db, nil
}

func GetMainDB() *gorm.DB {
	return mainDB
}

func InitDBCache(maxEntries int) {
	var err error
	tenantDBs, err = lru.New(maxEntries)
	if err != nil {
		log.Println(err)
	}
}

func GetTenantDB(dsn string) (*gorm.DB, error) {
	mu.RLock()
	if val, ok := tenantDBs.Get(dsn); ok {
		entry := val.(*tenantDBEntry)
		entry.lastUsed = time.Now()
		mu.RUnlock()
		return entry.db, nil
	}
	mu.RUnlock()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	sqlDB.SetConnMaxIdleTime(15 * time.Minute)

	entry := &tenantDBEntry{
		db:       db,
		lastUsed: time.Now(),
	}

	mu.Lock()
	tenantDBs.Add(dsn, entry)
	mu.Unlock()

	return db, nil
}

func StartDBJanitor() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		mu.Lock()
		for _, key := range tenantDBs.Keys() {
			if val, ok := tenantDBs.Get(key); ok {
				entry := val.(*tenantDBEntry)
				if time.Since(entry.lastUsed) > dbExpiration {
					if db, err := entry.db.DB(); err == nil {
						db.Close()
					}
					tenantDBs.Remove(key)
				}
			}
		}
		mu.Unlock()
	}
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("No se pudo obtener la conexión de bajo nivel:", err)
		return err
	}
	return sqlDB.Close()
}

func CloseAllTenantDBs() error {
	for _, key := range tenantDBs.Keys() {
		if val, ok := tenantDBs.Get(key); ok {
			entry := val.(*tenantDBEntry)
			if db, err := entry.db.DB(); err == nil {
				db.Close()
			}
			tenantDBs.Remove(key)
		}
	}
	return nil
}

func EnsureDatabaseExists(dsn string) error {
	parts := strings.Split(dsn, "/")
	if len(parts) < 2 {
		return fmt.Errorf("DSN inválido: %s", dsn)
	}
	dbNameAndParams := parts[1]
	dbName := strings.Split(dbNameAndParams, "?")[0]

	dsnWithoutDB := strings.Split(dsn, "/")[0] + "/?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsnWithoutDB)
	if err != nil {
		return fmt.Errorf("error al conectar sin base: %w", err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName))
	if err != nil {
		return fmt.Errorf("error al crear base %s: %w", dbName, err)
	}
	return nil
}
