package repositories

import (
	"errors"
	"time"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *ExpenseRepository) ExpenseGetByID(id string) (*models.Expense, error) {
	var expense models.Expense
	if err := r.DB.Where("id = ?", id).First(&expense).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error interno al buscar movimiento", err)
	}
	return &expense, nil
}

func (r *ExpenseRepository) ExpenseGetAll() (*[]models.Expense, error) {
	var expenses []models.Expense
	if err := r.DB.Limit(100).Order("created_at desc").Find(&expenses).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar movimientos", err)
	}
	return &expenses, nil
}

func (r *ExpenseRepository) ExpenseGetToday() (*[]models.Expense, error) {
	today := time.Now().Format("2006-01-02")
	var expenses []models.Expense
	if err := r.DB.Where("DATE(created_at) = ?", today).Order("created_at desc").Find(&expenses).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar movimientos", err)
	}
	return &expenses, nil
}

func (r *ExpenseRepository) ExpenseCreate(expense *models.ExpenseCreate) (string, error) {
	newID := uuid.NewString()
	if err := r.DB.Create(&models.Expense{
		ID:             newID,
		Details:        expense.Details,
		SupplierID:     expense.SupplierID,
		MovementTypeID: expense.MovementTypeID,
		Amount:         expense.Amount,
	}).Error; err != nil {
		return "", models.ErrorResponse(500, "Error interno al crear movimiento", err)
	}
	return newID, nil
}

func (r *ExpenseRepository) ExpenseUpdate(expense *models.ExpenseUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", expense.ID).
			Updates(&models.Expense{
				Details:        expense.Details,
				SupplierID:     expense.SupplierID,
				MovementTypeID: expense.MovementTypeID,
				Amount:         expense.Amount,
			}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.ErrorResponse(404, "Movimiento no encontrado", err)
			}
			return models.ErrorResponse(500, "Error interno al actualizar movimiento", err)
		}
		return nil
	})
}

func (r *ExpenseRepository) ExpenseDelete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Expense{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error interno al eliminar movimiento", err)
	}
	return nil
}
