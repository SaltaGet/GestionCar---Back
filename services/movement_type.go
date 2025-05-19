package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func MovementTypeCreate(movementType *models.MovementTypeCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreateMovementType(movementType, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func MovementTypeUpdate(movementType *models.MovementTypeUpdate, workplace string) error {
	err := repositories.Repo.UpdateMovementType(movementType, workplace)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}

	return nil
}

func MovementTypeDelete(id string, workplace string) error {
	err := repositories.Repo.DeleteMovementType(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func GetMovementTypeByID(id string, workplace string) (*models.MovementTypeLaundry, *models.MovementTypeWorkshop, error) {
	movementTypeLaundry, movementTypeWorkshop, err := repositories.Repo.GetMovementTypeByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return movementTypeLaundry, movementTypeWorkshop, nil
}

func GetAllMovementTypes(isIncome bool ,workplace string) (*[]models.MovementTypeLaundry, *[]models.MovementTypeWorkshop, error) {
	movementTypesLaundry, movementTypesWorkshop, err := repositories.Repo.GetAllMovementTypes(isIncome, workplace)
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return movementTypesLaundry, movementTypesWorkshop, nil
}