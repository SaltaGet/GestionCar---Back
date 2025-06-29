package database

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"   // MySQL driver
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3" // SQLite driver
	_ "github.com/golang-migrate/migrate/v4/source/file"     // File source for migrations

	// Importar el driver MySQL para su función de parseo de DSN
	// mysqlDriver "github.com/go-sql-driver/mysql"
)

// ApplyMigrations aplica las migraciones de la base de datos usando golang-migrate.
// dbURI: La cadena de conexión a la base de datos.
// migrationPath: La ruta absoluta al directorio que contiene los archivos de migración.
func ApplyMigrations(dbURI string, migrationPath string) error {
	// Verificar si el directorio de migraciones existe y está vacío
	entries, err := os.ReadDir(migrationPath)
	if err != nil {
		// Si el directorio no existe o hay otro error al leerlo, retornarlo
		return fmt.Errorf("error al leer el directorio de migraciones %s: %w", migrationPath, err)
	}

	// Filtrar . y .. si os.ReadDir los devuelve (normalmente no lo hace para entradas reales)
	// O simplemente verificar si hay archivos/directorios reales
	hasActualFiles := false
	for _, entry := range entries {
		if entry.Name() != "." && entry.Name() != ".." {
			hasActualFiles = true
			break
		}
	}

	if !hasActualFiles {
		fmt.Printf("Directorio de migraciones %s está vacío. No hay migraciones para aplicar.\n", migrationPath)
		return nil // No hay error, no hay migraciones para aplicar
	}

	var databaseURI string

	env := os.Getenv("ENV")

	if env == "prod" {
		databaseURI = dbURI // La URI de MySQL es directa
		// Asegurarse de que la base de datos exista antes de migrar
		if err := ensureDatabaseExists(dbURI); err != nil {
			return fmt.Errorf("error al asegurar que la base de datos exista: %w", err)
		}
	} else {
		// Para SQLite, migrate necesita la ruta del archivo con el prefijo "sqlite3://"
		databaseURI = "sqlite3://" + filePathFromURI(dbURI)
	}

	m, err := migrate.New(
		"file://"+migrationPath, // Ruta a tus archivos de migración
		databaseURI,
	)
	if err != nil {
		return fmt.Errorf("error al crear instancia de migrate: %w", err)
	}

	// Aplicar todas las migraciones pendientes
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error al aplicar migraciones: %w", err)
	}

	fmt.Println("Migraciones aplicadas exitosamente.")
	return nil
}

// filePathFromURI es una función auxiliar para extraer la ruta del archivo de una URI de SQLite
// func filePathFromURI(uri string) string {
// 	// Asume que la URI de SQLite es algo como "file:./data.db?_foreign_keys=on"
// 	// o simplemente "./data.db"
// 	if strings.HasPrefix(uri, "file:") {
// 		parts := strings.SplitN(uri, "?", 2)
// 		return strings.TrimPrefix(parts[0], "file:")
// 	}
// 	return uri
// }

// // ensureDatabaseExists es una función auxiliar para MySQL que usa el parser oficial del driver.
// func ensureDatabaseExists(uri string) error {
// 	// Usar el parser oficial del driver MySQL
// 	cfg, err := mysqlDriver.ParseDSN(uri)
// 	if err != nil {
// 		return fmt.Errorf("error al parsear DSN de MySQL: %w", err)
// 	}

// 	dbName := cfg.DBName // Guardar el nombre de la base de datos
// 	cfg.DBName = ""      // Eliminar el nombre de la base de datos para conectar al servidor

// 	// Conectar al servidor MySQL sin especificar una base de datos
// 	db, err := sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		return fmt.Errorf("error al conectar a MySQL para verificar DB: %w", err)
// 	}
// 	defer db.Close()

// 	// Crear la base de datos si no existe
// 	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName))
// 	if err != nil {
// 		return fmt.Errorf("error al crear base de datos %s: %w", dbName, err)
// 	}
// 	return nil
// }