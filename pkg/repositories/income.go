package repositories

import (
	"time"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *TenantRepository) GetIncomeByID(id string) (*models.Income, error) {
		var income models.Income
		if err := r.DB.Where("id = ?", id).First(&income).Error; err != nil {
			return nil, err
		}
		return &income, nil
}

func (r *TenantRepository) GetAllIncomes() (*[]models.Income, error) {
		var incomes []models.Income
		if err := r.DB.Limit(100).Order("created_at desc").Find(&incomes).Error; err != nil {
			return nil, err
		}
		return &incomes, nil
}

func (r *TenantRepository) GetIncomeToday() (*[]models.Income, error) {
	today := time.Now().Format("2006-01-02")
		var incomes []models.Income
		if err := r.DB.Where("DATE(created_at) = ?", today).Order("created_at desc").Find(&incomes).Error; err != nil {
			return nil, err
		}
		return &incomes, nil
}

func (r *TenantRepository) CreateIncome(income *models.IncomeCreate) (string, error) {
	newID := uuid.NewString()
	err := r.DB.Transaction(func(tx *gorm.DB) error {
			if err := r.DB.Create(&models.Income{
				ID:             newID,
				Ticket:         income.Ticket,
				Details:        income.Details,
				ClientID:       income.ClientID,
				VehicleID:      income.VehicleID,
				EmployeeID:     income.EmployeeID,
				Amount:         income.Amount,
				MovementTypeID: income.MovementTypeID,
			}).Error; err != nil {
				return err
			}

			for _, item := range income.ServicesID {
				if err := r.DB.Create(&models.IncomeService{
					ID:              newID,
					IncomeLaundryID: newID,
					ServiceID:       item,
				}).Error; err != nil {
					return err
				}
			}
			return nil
	})
	if err != nil {
		return "", err
	}
	return newID, nil
}

func (r *TenantRepository) UpdateIncome(income *models.IncomeUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("id = ?", income.ID).
				Updates(&models.Income{
					Ticket:         income.Ticket,
					Details:        income.Details,
					ClientID:       income.ClientID,
					VehicleID:      income.VehicleID,
					EmployeeID:     income.EmployeeID,
					Amount:         income.Amount,
					MovementTypeID: income.MovementTypeID,
				}).Error; err != nil {
				return err
			}

			var existingProducts []models.IncomeService
			if err := tx.Where("income_laundry_id = ?", income.ID).Find(&existingProducts).Error; err != nil {
				return err
			}
			existingIDs := map[string]bool{}
			for _, p := range existingProducts {
				existingIDs[p.ID] = true
			}

			receivedIDs := map[string]bool{}
			for _, prod := range income.ServicesID {
				receivedIDs[prod] = true
			}

			for _, p := range existingProducts {
				if !receivedIDs[p.ID] {
					if err := tx.Delete(&models.IncomeService{}, "id = ?", p.ID).Error; err != nil {
						return err
					}
				}
			}

			for _, prod := range income.ServicesID {
				if prod == "" || !existingIDs[prod] {
					newProd := models.IncomeService{
						ID:              uuid.NewString(),
						IncomeLaundryID: income.ID,
						ServiceID:       prod,
					}
					if err := tx.Create(&newProd).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Model(&models.IncomeService{}).
						Where("id = ?", prod).
						Updates(map[string]interface{}{
							"service_id": prod,
						}).Error; err != nil {
						return err
					}
				}
			}
			return nil
	})
}

func (r *TenantRepository) DeleteIncomeByID(id string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("income_laundry_id = ?", id).Delete(&models.IncomeService{}).Error; err != nil {
				return err
			}
			if err := tx.Where("id = ?", id).Delete(&models.Income{}).Error; err != nil {
				return err
			}
		return nil
	})
}
