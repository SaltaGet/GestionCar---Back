package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/repositories"
	"gorm.io/gorm"
)

func ExpenseResumeCreate(resume *models.ExpenseResumeCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreateExpenseResume(resume, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear resumen", err)
	}

	return id, nil
}

func GetExpenseResumeByDateBetween(fromDate string, toDate string, workplace string) (*[]models.ExpenseResumeLaundry, *[]models.ExpenseResumeWorkshop, error) {
	laundry, workshop, err := repositories.Repo.GetExpenseResumeByDateBetween(fromDate, toDate, workplace)
	if err != nil {
		return nil, nil, err
	}

	return laundry, workshop, nil
}

func GetExpenseResumeByID(id string, workplace string) (*models.ExpenseResumeLaundry, *models.ExpenseResumeWorkshop, error) {
	laundry, workshop, err := repositories.Repo.GetExpenseResumeByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Resumen no encontrado", err)
		}
		return nil, nil, err
	}

	return laundry, workshop, nil
}

func UpdateExpenseResume(resume *models.ExpenseResumeUpdate, workplace string) error {
	err := repositories.Repo.UpdateExpenseResume(resume, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Resumen no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar resumen", err)
	}

	return nil
}


// INCOME

func CreateIncomeResume(resume *models.IncomeResumeCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreateIncomeResume(resume, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear resumen", err)
	}

	return id, nil
}

func GetIncomeResumeByDateBetween(fromDate string, toDate string, workplace string) (*[]models.IncomeResumeLaundry, *[]models.IncomeResumeWorkshop, error) {
	laundry, workshop, err := repositories.Repo.GetIncomeResumeByDateBetween(fromDate, toDate, workplace)
	if err != nil {
		return nil, nil, err
	}

	return laundry, workshop, nil
}

func GetIncomeResumeByID(id string, workplace string) (*models.IncomeResumeLaundry, *models.IncomeResumeWorkshop, error) {
	laundry, workshop, err := repositories.Repo.GetIncomeResumeByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Resumen no encontrado", err)
		}
		return nil, nil, err
	}

	return laundry, workshop, nil
}

func UpdateIncomeResume(resume *models.IncomeResumeUpdate, workplace string) error {
	err := repositories.Repo.UpdateIncomeResume(resume, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Resumen no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar resumen", err)
	}

	return nil
}
