package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (e *ExpenseService) ExpenseGetByID(id string) (*models.Expense, error) {
	expense, err := e.ExpenseRepository.ExpenseGetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar movimiento", err)
	}

	return expense, nil
}

func (e *ExpenseService) ExpenseGetAll(workplace string) (*[]models.Expense, error) {
	expenses, err := e.ExpenseRepository.ExpenseGetAll()
	
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return expenses, nil
}

func (e *ExpenseService) ExpenseGetToday(workplace string) (*[]models.Expense, error) {
	expenses, err := e.ExpenseRepository.ExpenseGetToday()
	
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return expenses, nil
}

func (e *ExpenseService) ExpenseCreate(expense *models.ExpenseCreate) (string, error) {
	id, err := e.ExpenseRepository.ExpenseCreate(expense)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear movimiento", err)
	}
	return id, nil
}

func (e *ExpenseService) ExpenseUpdate(expense *models.ExpenseUpdate) error {
	err := e.ExpenseRepository.ExpenseUpdate(expense)
	if err != nil {
		return models.ErrorResponse(500, "Error al actualizar movimiento", err)
	}
	return nil
}

func (e *ExpenseService) ExpenseDelete(id string) error {
	err := e.ExpenseRepository.ExpenseDelete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar movimiento", err)
	}
	return nil
}