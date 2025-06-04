package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (i *IncomeService) IncomeGetByID(id string) (*models.Income, error) {
	income, err := i.IncomeRepository.IncomeGetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar movimiento", err)
	}

	return income, nil
}

func (i *IncomeService) IncomeGetAll() (*[]models.Income, error) {
	incomes, err := i.IncomeRepository.IncomeGetAll()
	
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return incomes, nil
}

func (i *IncomeService) IncomeGetToday() (*[]models.Income, error) {
	incomes, err := i.IncomeRepository.IncomeGetToday()
	
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return incomes, nil
}

func (i *IncomeService) IncomeCreate(incomeCreate *models.IncomeCreate) (string, error) {
	id, err := i.IncomeRepository.IncomeCreate(incomeCreate)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear movimiento", err)
	}
	return id, nil
}

func (i *IncomeService) IncomeUpdate(incomeUpdate *models.IncomeUpdate) error {
	err := i.IncomeRepository.IncomeUpdate(incomeUpdate)
	if err != nil {
		return models.ErrorResponse(500, "Error al actualizar movimiento", err)
	}
	return nil
}

func (i *IncomeService) IncomeDelete(id string) error {
	err := i.IncomeRepository.IncomeDelete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar movimiento", err)
	}
	return nil
}