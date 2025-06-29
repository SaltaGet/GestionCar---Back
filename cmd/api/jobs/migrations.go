package jobs

import (
	"os"
	"path/filepath"

	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
)

func Migrations(deps *dependencies.Application, path string) error {
	connections, err := deps.TenantController.TenantService.TenantGetConections()
	if err != nil {
		return err
	}

	tenant := filepath.Join(path, "tenant")
	main := filepath.Join(path, "main")

	err = database.ApplyMigrations(os.Getenv("URI_DB"), main)
	if err != nil {
		return err
	}

	for _, connection := range *connections {
		err = database.ApplyMigrations(connection, tenant)
		if err != nil {
			return err
		}

	}
	return nil
}
