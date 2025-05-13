package repositories

import (
	"fmt"
	"time"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetIncomeByID(id string, workplace string) (*models.IncomeLaundry, *models.IncomeWorkshop, error) {
	switch workplace {
	case "laundry":
		var income models.IncomeLaundry
		if err := r.DB.Where("id = ?", id).First(&income).Error; err != nil {
			return nil, nil, err
		}
		return &income, nil, nil
	case "workshop":
		var income models.IncomeWorkshop
		if err := r.DB.Where("id = ?", id).First(&income).Error; err != nil {
			return nil, nil, err
		}
		return nil, &income, nil
	default:
		return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) GetAllIncomes(workplace string) (*[]models.IncomeLaundry, *[]models.IncomeWorkshop, error) {
	if workplace == "laundry" {
		var incomes []models.IncomeLaundry
		if err := r.DB.Limit(100).Order("created_at desc").Find(&incomes).Error; err != nil {
			return nil, nil, err
		}
		return &incomes, nil, nil
	} else if workplace == "workshop" {
		var incomes []models.IncomeWorkshop
		if err := r.DB.Limit(100).Order("created_at desc").Find(&incomes).Error; err != nil {
			return nil, nil, err
		}
		return nil, &incomes, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetIncomeToday(workplace string) (*[]models.IncomeLaundry, *[]models.IncomeWorkshop, error) {
	today := time.Now().Format("2006-01-02")
	switch workplace {
	case "laundry":
		var incomes []models.IncomeLaundry
		if err := r.DB.Where("DATE(created_at) = ?", today).Order("created_at desc").Find(&incomes).Error; err != nil {
			return nil, nil, err
		}
		return &incomes, nil, nil
	case "workshop":
		var incomes []models.IncomeWorkshop
		if err := r.DB.Where("DATE(created_at) = ?", today).Order("created_at desc").Find(&incomes).Error; err != nil {
			return nil, nil, err
		}
		return nil, &incomes, nil
	default:
		return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) CreateIncome(income *models.IncomeCreate, workplace string) (string, error) {
	newID := uuid.NewString()
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		switch workplace {
		case "laundry":
			if err := r.DB.Create(&models.IncomeLaundry{
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
				if err := r.DB.Create(&models.IncomeServiceLaundry{
					ID:              newID,
					IncomeLaundryID: newID,
					ServiceID:       item,
				}).Error; err != nil {
					return err
				}
			}
			return nil
		case "workshop":
			if err := r.DB.Create(&models.IncomeWorkshop{
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
				if err := r.DB.Create(&models.IncomeServiceWorkshop{
					ID:               newID,
					IncomeWorkshopID: newID,
					ServiceID:        item,
				}).Error; err != nil {
					return err
				}
			}
			return nil
		default:
			return fmt.Errorf("tipo de movimiento no soportado")
		}
	})
	if err != nil {
		return "", err
	}
	return newID, nil
}

func (r *Repository) UpdateIncome(income *models.IncomeUpdate, workplace string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		switch workplace {
		case "laundry":
			if err := tx.Where("id = ?", income.ID).
				Updates(&models.IncomeLaundry{
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

			var existingProducts []models.IncomeServiceLaundry
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
					if err := tx.Delete(&models.IncomeServiceLaundry{}, "id = ?", p.ID).Error; err != nil {
						return err
					}
				}
			}

			for _, prod := range income.ServicesID {
				if prod == "" || !existingIDs[prod] {
					newProd := models.IncomeServiceLaundry{
						ID:              uuid.NewString(),
						IncomeLaundryID: income.ID,
						ServiceID:       prod,
					}
					if err := tx.Create(&newProd).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Model(&models.IncomeServiceLaundry{}).
						Where("id = ?", prod).
						Updates(map[string]interface{}{
							"service_id": prod,
						}).Error; err != nil {
						return err
					}
				}
			}
			return nil
		case "workshop":
			if err := tx.Where("id = ?", income.ID).
				Updates(&models.IncomeWorkshop{
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

			var existingProducts []models.IncomeServiceWorkshop
			if err := tx.Where("income_workshop_id = ?", income.ID).Find(&existingProducts).Error; err != nil {
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
					if err := tx.Delete(&models.IncomeServiceWorkshop{}, "id = ?", p.ID).Error; err != nil {
						return err
					}
				}
			}

			for _, prod := range income.ServicesID {
				if prod == "" || !existingIDs[prod] {
					newProd := models.IncomeServiceWorkshop{
						ID:              uuid.NewString(),
						IncomeWorkshopID: income.ID,
						ServiceID:       prod,
					}
					if err := tx.Create(&newProd).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Model(&models.IncomeServiceWorkshop{}).
						Where("id = ?", prod).
						Updates(map[string]interface{}{
							"service_id": prod,
						}).Error; err != nil {
						return err
					}
				}
			}
			return nil
		default:
			return fmt.Errorf("tipo de movimiento no soportado")
		}
	})
}

func (r *Repository) DeleteIncomeByID(id string, workplace string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if workplace == "laundry" {
			if err := tx.Where("income_laundry_id = ?", id).Delete(&models.IncomeServiceLaundry{}).Error; err != nil {
				return err
			}
			if err := tx.Where("id = ?", id).Delete(&models.IncomeLaundry{}).Error; err != nil {
				return err
			}
		} else if workplace == "workshop" {
			if err := tx.Where("income_workshop_id = ?", id).Delete(&models.IncomeServiceWorkshop{}).Error; err != nil {
				return err
			}
			if err := tx.Where("id = ?", id).Delete(&models.IncomeWorkshop{}).Error; err != nil {
				return err
			}
		} else {
			return fmt.Errorf("tipo de movimiento no soportado")
		}
		return nil
	})
}
