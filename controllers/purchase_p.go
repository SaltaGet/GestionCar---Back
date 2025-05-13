package controllers

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func PurchaseProductCreate(purchaseOrder *models.PurchaseProductCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreatePurchaseElement(purchaseOrder, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func PurchaseProductUpdate(purchaseOrder *models.PurchaseProductUpdate, workplace string) error {
	err := repositories.Repo.UpdatePurchaseElement(purchaseOrder, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func PurchaseProductDelete(id string, workplace string) error {
	err := repositories.Repo.DeletePurchaseElementByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func PurchaseProductGetByID(id string, workplace string) (*models.PurchaseProductLaundry, *models.PurchasePartWorkshop, error) {
	purchaseOrderLaundry, purchaseOrderWorkshop, err := repositories.Repo.GetPurchaseElementByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return purchaseOrderLaundry, purchaseOrderWorkshop, nil
}

func PurchaseProductGetAllByPurhcaseID(purchaseID string, workplace string) (*[]models.PurchaseProductLaundry, *[]models.PurchasePartWorkshop, error) {
	purchaseOrderLaundry, purchaseOrderWorkshop, err := repositories.Repo.GetPurchaseElementByPurchaseID(purchaseID, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return purchaseOrderLaundry, purchaseOrderWorkshop, nil
}