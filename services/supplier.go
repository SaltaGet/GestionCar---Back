package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func SupplierCreate(supplier *models.SupplierCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreateSupplier(supplier, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear proveedor", err)
	}
	return id, nil
}

func SupplierGetAll(workplace string) (*[]models.SupplierLaundry, *[]models.SupplierWorkshop, error) {
	suppliersLaundry, suppliersWorkshop, err := repositories.Repo.GetAllSuppliers(workplace)
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al buscar los proveedores", err)
	}
	return &suppliersLaundry, &suppliersWorkshop,nil
}

func SupplierGetByID(id string, workplace string) (*models.SupplierLaundry, *models.SupplierWorkshop, error) {
	supplierLaundry, supplierWorkshop, err := repositories.Repo.GetSupplierByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Proveedor no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al buscar proveedor", err)
	}
	return supplierLaundry, supplierWorkshop, nil
}

func SupplierGetByName(name string, workplace string) (*models.SupplierLaundry, *models.SupplierWorkshop, error) {
	supplierLaundry, supplierWorkshop, err := repositories.Repo.GetSupplierByName(name, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Proveedor no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al buscar proveedor", err)
	}
	return supplierLaundry, supplierWorkshop, nil
}

func SupplierDeleteByID(id string, workplace string) error {
	err := repositories.Repo.DeleteSupplierByID(id, workplace)
	if err != nil {
		return models.ErrorResponse(500, "Error al eliminar proveedor", err)
	}
	return nil
}

func SupplierUpdate(id string, supplierUpdate *models.SupplierUpdate, workplace string) error {
	err := repositories.Repo.UpdateSupplier(id, supplierUpdate, workplace)
	if err != nil {
		return models.ErrorResponse(500, "Error al actualizar proveedor", err)
	}
	return nil
}