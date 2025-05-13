package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func GetIncomeByID(id string, workplace string) (*models.IncomeLaundry, *models.IncomeWorkshop, error) {
	laundries, workshops, err := repositories.Repo.GetIncomeByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al buscar movimiento", err)
	}

	return laundries, workshops, nil
}

func GetAllIncomes(workplace string) (*[]models.IncomeLaundry, *[]models.IncomeWorkshop, error) {
	laundries, workshops, err := repositories.Repo.GetAllIncomes(workplace)
	
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return laundries, workshops, nil
}

func GetIncomeToday(workplace string) (*[]models.IncomeLaundry, *[]models.IncomeWorkshop, error) {
	laundries, workshops, err := repositories.Repo.GetIncomeToday(workplace)
	
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al buscar movimientos", err)
	}

	return laundries, workshops, nil
}

func CreateIncome(expense *models.IncomeCreate, workplace string) (string, error) {
	laundryID, err := repositories.Repo.CreateIncome(expense, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear movimiento", err)
	}
	return laundryID, nil
}

func UpdateIncome(expense *models.IncomeUpdate, workplace string) error {
	err := repositories.Repo.UpdateIncome(expense, workplace)
	if err != nil {
		return models.ErrorResponse(500, "Error al actualizar movimiento", err)
	}
	return nil
}

func DeleteIncome(id string, workplace string) error {
	err := repositories.Repo.DeleteExpenseByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar movimiento", err)
	}
	return nil
}