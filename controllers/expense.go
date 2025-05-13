package controllers

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func GetExpenseByID(id string, workplace string) (*models.ExpenseLaundry, *models.ExpenseWorkshop, error) {
	laundries, workshops, err := repositories.Repo.GetExpenseByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al buscar movimiento", err)
	}

	return laundries, workshops, nil
}

func GetAllExpenses(workplace string) (*[]models.ExpenseLaundry, *[]models.ExpenseWorkshop, error) {
	laundries, workshops, err := repositories.Repo.GetAllExpenses(workplace)
	
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return laundries, workshops, nil
}

func GetExpenseToday(workplace string) (*[]models.ExpenseLaundry, *[]models.ExpenseWorkshop, error) {
	laundries, workshops, err := repositories.Repo.GetExpenseToday(workplace)
	
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return laundries, workshops, nil
}

func CreateExpense(expense *models.ExpenseCreate, workplace string) (string, error) {
	laundryID, err := repositories.Repo.CreateExpense(expense, "laundry")
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear movimiento", err)
	}
	return laundryID, nil
}

func UpdateExpense(expense *models.ExpenseUpdate, workplace string) error {
	err := repositories.Repo.UpdateExpense(expense, workplace)
	if err != nil {
		return models.ErrorResponse(500, "Error al actualizar movimiento", err)
	}
	return nil
}

func DeleteExpense(id string, workplace string) error {
	err := repositories.Repo.DeleteExpenseByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar movimiento", err)
	}
	return nil
}