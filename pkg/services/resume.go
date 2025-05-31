package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (r *ResumeService) ExpenseResumeCreate(resume *models.ResumeExpenseCreate) (string, error) {
	id, err := r.ResumeExpenseRepository.CreateExpenseResume(resume)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear resumen", err)
	}

	return id, nil
}

func (r *ResumeService) GetExpenseResumeByDateBetween(fromDate string, toDate string) (*[]models.ResumeExpense, error) {
	resumes, err := r.ResumeExpenseRepository.GetExpenseResumeByDateBetween(fromDate, toDate)
	if err != nil {
		return nil, err
	}

	return resumes, nil
}

func (r *ResumeService) GetExpenseResumeByID(id string) (*models.ResumeExpense, error) {
	resumes, err := r.ResumeExpenseRepository.GetExpenseResumeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Resumen no encontrado", err)
		}
		return nil, err
	}

	return resumes, nil
}

func (r *ResumeService) UpdateExpenseResume(resume *models.ResumeExpenseUpdate) error {
	err := r.ResumeExpenseRepository.UpdateExpenseResume(resume)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Resumen no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar resumen", err)
	}

	return nil
}


// INCOME

func (r *ResumeService) CreateIncomeResume(resume *models.ResumeIncomeCreate) (string, error) {
	id, err := r.ResumeIncomeRepository.IncomeResumeCreate(resume)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear resumen", err)
	}

	return id, nil
}

func (r *ResumeService) GetIncomeResumeByDateBetween(fromDate string, toDate string) (*[]models.ResumeIncome, error) {
	incomes, err := r.ResumeIncomeRepository.GetIncomeResumeByDateBetween(fromDate, toDate)
	if err != nil {
		return nil, err
	}

	return incomes, nil
}

func (r *ResumeService) GetIncomeResumeByID(id string) (*models.ResumeIncome, error) {
	income, err := r.ResumeIncomeRepository.GetIncomeResumeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Resumen no encontrado", err)
		}
		return nil, err
	}

	return income, nil
}

func (r *ResumeService) UpdateIncomeResume(resume *models.ResumeIncomeUpdate) error {
	err := r.ResumeIncomeRepository.UpdateIncomeResume(resume)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Resumen no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar resumen", err)
	}

	return nil
}
