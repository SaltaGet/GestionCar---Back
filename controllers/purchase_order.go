package controllers

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func PurchaseOrderGetByID(id string, workplace string) (*models.PurchaseOrderLaundry, *models.PurchaseOrderWorkshop, error) {
	purchaseOrderLaundry, purchaseOrderWorkshop, err := repositories.Repo.GetPurchaseOrderByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return purchaseOrderLaundry, purchaseOrderWorkshop, nil
}

func PurchaseOrderGetAll(workplace string) (*[]models.PurchaseOrderLaundry, *[]models.PurchaseOrderWorkshop, error) {
	purchaseOrderLaundry, purchaseOrderWorkshop, err := repositories.Repo.GetAllPurchaseOrders(workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return purchaseOrderLaundry, purchaseOrderWorkshop, nil
}

func PurchaseOrderCreate(purchaseOrder *models.PurchaseOrderCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreatePurchaseOrder(purchaseOrder, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func PurchaseOrderUpdate(purchaseOrder *models.PurchaseOrderUpdate, workplace string) error {
	err := repositories.Repo.UpdatePurchaseOrder(purchaseOrder, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func PurchaseOrderDelete(id string, workplace string) error {
	err := repositories.Repo.DeletePurchaseOrderByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}