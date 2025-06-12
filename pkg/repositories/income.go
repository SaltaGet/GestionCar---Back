package repositories

import (
	"errors"
	"time"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *TenantRepository) IncomeGetByID(id string) (*models.Income, error) {
	var income models.Income
	if err := r.DB.Where("id = ?", id).First(&income).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Movimiento no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error interno al buscar movimiento", err)
	}
	return &income, nil
}

func (r *TenantRepository) IncomeGetAll() (*[]models.Income, error) {
	var incomes []models.Income
	if err := r.DB.Limit(100).Order("created_at desc").Find(&incomes).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar movimientos", err)
	}
	return &incomes, nil
}

func (r *TenantRepository) IncomeGetToday() (*[]models.Income, error) {
	today := time.Now().Format("2006-01-02")
	var incomes []models.Income
	if err := r.DB.Where("DATE(created_at) = ?", today).Order("created_at desc").Find(&incomes).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar movimientos", err)
	}
	return &incomes, nil
}

func (r *TenantRepository) IncomeCreate(income *models.IncomeCreate) (string, error) {
	newID := uuid.NewString()

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		var services []models.Service

		if err := tx.Where("id IN ?", income.ServicesID).Find(&services).Error; err != nil {
			return models.ErrorResponse(500, "Error interno al buscar servicios", err) 
		}

		if err := tx.Create(&models.Income{
			ID:             newID,
			Ticket:         income.Ticket,
			Details:        income.Details,
			ClientID:       income.ClientID,
			VehicleID:      income.VehicleID,
			EmployeeID:     income.EmployeeID,
			Amount:         income.Amount,
			MovementTypeID: income.MovementTypeID,
			Services:       services,
		}).Error; err != nil {
			return models.ErrorResponse(500, "Error interno al crear movimiento", err)
		}

		return nil 
	})

	if err != nil {
		return "", models.ErrorResponse(500, "Error interno al crear movimiento", err)
	}

	return newID, nil
}


func (r *TenantRepository) IncomeUpdate(incomeUpdate *models.IncomeUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var income models.Income

		if err := tx.Where("id = ?", incomeUpdate.ID).First(&income).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.ErrorResponse(404, "Movimiento no encontrado", err)
			}
			return models.ErrorResponse(500, "Error interno al actualizar movimiento", err)
		}

		income.Ticket = incomeUpdate.Ticket
		income.Details = incomeUpdate.Details
		income.ClientID = incomeUpdate.ClientID
		income.VehicleID = incomeUpdate.VehicleID
		income.EmployeeID = incomeUpdate.EmployeeID
		income.Amount = incomeUpdate.Amount
		income.MovementTypeID = incomeUpdate.MovementTypeID
		income.UpdatedAt = time.Now().UTC()

		var services []models.Service
		if err := tx.Where("id IN ?", incomeUpdate.ServicesID).Find(&services).Error; err != nil {
			return models.ErrorResponse(500, "Error interno al buscar servicios", err)
		}

		if err := tx.Model(&income).Association("Services").Replace(services); err != nil {
			return models.ErrorResponse(500, "Error interno al actualizar servicios", err)
		}

		if err := tx.Save(&income).Error; err != nil {
			return models.ErrorResponse(500, "Error interno al actualizar movimiento", err)
		}

		return nil 
	})
}

func (r *TenantRepository) IncomeDelete(id string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&models.Income{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.ErrorResponse(404, "Movimiento no encontrado", err)
			}
			return models.ErrorResponse(500, "Error interno al eliminar movimiento", err)
		}
		return nil
	})
}
