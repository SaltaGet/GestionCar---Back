package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (i *IncomeService) GetIncomeByID(id string) (*models.Income, error) {
	income, err := i.IncomeRepository.GetIncomeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar movimiento", err)
	}

	return income, nil
}

func (i *IncomeService) GetAllIncomes() (*[]models.Income, error) {
	incomes, err := i.IncomeRepository.GetAllIncomes()
	
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return incomes, nil
}

func (i *IncomeService) GetIncomeToday() (*[]models.Income, error) {
	incomes, err := i.IncomeRepository.GetIncomeToday()
	
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return incomes, nil
}

func (i *IncomeService) CreateIncome(incomeCreate *models.IncomeCreate) (string, error) {
	id, err := i.IncomeRepository.CreateIncome(incomeCreate)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear movimiento", err)
	}
	return id, nil
}

func (i *IncomeService) UpdateIncome(incomeUpdate *models.IncomeUpdate) error {
	err := i.IncomeRepository.UpdateIncome(incomeUpdate)
	if err != nil {
		return models.ErrorResponse(500, "Error al actualizar movimiento", err)
	}
	return nil
}

func (i *IncomeService) DeleteIncome(id string) error {
	err := i.IncomeRepository.DeleteIncome(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar movimiento", err)
	}
	return nil
}