package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (s *SupplierService) SupplierCreate(supplier *models.SupplierCreate) (string, error) {
	id, err := s.SupplierRepository.SupplierCreate(supplier)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear proveedor", err)
	}
	return id, nil
}

func (s *SupplierService) SupplierGetAll() (*[]models.Supplier, error) {
	suppliers, err := s.SupplierRepository.SupplierGetAll()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los proveedores", err)
	}
	return suppliers,nil
}

func (s *SupplierService) SupplierGetByID(id string) (*models.Supplier, error) {
	supplier, err := s.SupplierRepository.SupplierGetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Proveedor no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar proveedor", err)
	}
	return supplier, nil
}

func (s *SupplierService) SupplierGetByName(name string) (*[]models.Supplier, error) {
	supplier, err := s.SupplierRepository.SupplierGetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Proveedor no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar proveedor", err)
	}
	return supplier, nil
}

func (s *SupplierService) SupplierDelete(id string) error {
	err := s.SupplierRepository.SupplierDelete(id)
	if err != nil {
		return models.ErrorResponse(500, "Error al eliminar proveedor", err)
	}
	return nil
}

func (s *SupplierService) SupplierUpdate(supplierUpdate *models.SupplierUpdate) error {
	err := s.SupplierRepository.SupplierUpdate(supplierUpdate)
	if err != nil {
		return models.ErrorResponse(500, "Error al actualizar proveedor", err)
	}
	return nil
}