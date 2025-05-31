package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (e *ExpenseService) GetExpenseByID(id string) (*models.Expense, error) {
	expense, err := e.ExpenseRepository.GetExpenseByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar movimiento", err)
	}

	return expense, nil
}

func (e *ExpenseService) GetAllExpenses(workplace string) (*[]models.Expense, error) {
	expenses, err := e.ExpenseRepository.GetAllExpenses()
	
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return expenses, nil
}

func (e *ExpenseService) GetExpenseToday(workplace string) (*[]models.Expense, error) {
	expenses, err := e.ExpenseRepository.GetExpenseToday()
	
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return expenses, nil
}

func (e *ExpenseService) CreateExpense(expense *models.ExpenseCreate) (string, error) {
	id, err := e.ExpenseRepository.CreateExpense(expense)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear movimiento", err)
	}
	return id, nil
}

func (e *ExpenseService) UpdateExpense(expense *models.ExpenseUpdate) error {
	err := e.ExpenseRepository.UpdateExpense(expense)
	if err != nil {
		return models.ErrorResponse(500, "Error al actualizar movimiento", err)
	}
	return nil
}

func (e *ExpenseService) DeleteExpense(id string) error {
	err := e.ExpenseRepository.DeleteExpense(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar movimiento", err)
	}
	return nil
}