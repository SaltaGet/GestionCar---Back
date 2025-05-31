package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (p *PurchaseOrderService) PurchaseOrderGetByID(id string) (*models.PurchaseOrder, error) {
	purchaseOrder, err := p.PurchaseOrderRepository.PurchaseOrderGetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return purchaseOrder, nil
}

func (p *PurchaseOrderService) PurchaseOrderGetAll() (*[]models.PurchaseOrder, error) {
	purchaseOrder, err := p.PurchaseOrderRepository.PurchaseOrderGetAll()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return purchaseOrder, nil
}

func (p *PurchaseOrderService) PurchaseOrderCreate(purchaseOrder *models.PurchaseOrderCreate) (string, error) {
	id, err := p.PurchaseOrderRepository.PurchaseOrderCreate(purchaseOrder)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func (p *PurchaseOrderService) PurchaseOrderUpdate(purchaseOrder *models.PurchaseOrderUpdate) error {
	err := p.PurchaseOrderRepository.PurchaseOrderUpdate(purchaseOrder)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func (p *PurchaseOrderService) PurchaseOrderDelete(id string) error {
	err := p.PurchaseOrderRepository.PurchaseOrderDelete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}