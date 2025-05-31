package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (m *MovementTypeService) MovementTypeCreate(movementType *models.MovementTypeCreate) (string, error) {
	id, err := m.MovementTypeRepository.CreateMovementType(movementType)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}

	return id, nil
}

func (m *MovementTypeService) MovementTypeUpdate(movementType *models.MovementTypeUpdate) error {
	err := m.MovementTypeRepository.UpdateMovementType(movementType)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}

	return nil
}

func (m *MovementTypeService) MovementTypeDelete(id string) error {
	err := m.MovementTypeRepository.DeleteMovementType(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func (m *MovementTypeService) GetMovementTypeByID(id string) (*models.MovementType, error) {
	movementType, err := m.MovementTypeRepository.GetMovementTypeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return movementType, nil
}

func (m *MovementTypeService) GetAllMovementTypes(isIncome bool) (*[]models.MovementType, error) {
	movementTypes, err := m.MovementTypeRepository.GetAllMovementTypes(isIncome)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return movementTypes, nil
}