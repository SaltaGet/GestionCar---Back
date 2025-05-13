package repositories

import (
	"fmt"
	"time"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetExpenseByID(id string, workplace string) (*models.ExpenseLaundry, *models.ExpenseWorkshop, error) {
	switch workplace {
	case "laundry":
		var expense models.ExpenseLaundry
		if err := r.DB.Where("id = ?", id).First(&expense).Error; err != nil {
			return nil, nil, err
		}
		return &expense, nil, nil
	case "workshop":
		var expense models.ExpenseWorkshop
		if err := r.DB.Where("id = ?", id).First(&expense).Error; err != nil {
			return nil, nil, err
		}
		return nil, &expense, nil
	default:
		return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) GetAllExpenses(workplace string) (*[]models.ExpenseLaundry, *[]models.ExpenseWorkshop, error) {
	if workplace == "laundry" {
		var expenses []models.ExpenseLaundry
		if err := r.DB.Limit(100).Order("created_at desc").Find(&expenses).Error; err != nil {
			return nil, nil, err
		}
		return &expenses, nil, nil
	} else if workplace == "workshop" {
		var expenses []models.ExpenseWorkshop
		if err := r.DB.Limit(100).Order("created_at desc").Find(&expenses).Error; err != nil {
			return nil, nil, err
		}
		return nil, &expenses, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetExpenseToday(workplace string) (*[]models.ExpenseLaundry, *[]models.ExpenseWorkshop, error) {
    today := time.Now().Format("2006-01-02")
    switch workplace {
    case "laundry":
        var expenses []models.ExpenseLaundry
        if err := r.DB.Where("DATE(created_at) = ?", today).Order("created_at desc").Find(&expenses).Error; err != nil {
            return nil, nil, err
        }
        return &expenses, nil, nil
    case "workshop":
        var expenses []models.ExpenseWorkshop
        if err := r.DB.Where("DATE(created_at) = ?", today).Order("created_at desc").Find(&expenses).Error; err != nil {
            return nil, nil, err
        }
        return nil, &expenses, nil
    default:
        return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
    }
}

func (r *Repository) CreateExpense(expense *models.ExpenseCreate, workplace string) (string, error) {
	newID := uuid.NewString()
	switch workplace {
	case "laundry":
		if err := r.DB.Create(&models.ExpenseLaundry{
			ID:             newID,
			Details:        expense.Details,
			SupplierID:     expense.SupplierID,
			MovementTypeID: expense.MovementTypeID,
			Amount:         expense.Amount,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	case "workshop":
		if err := r.DB.Create(&models.ExpenseWorkshop{
			ID:             newID,
			Details:        expense.Details,
			SupplierID:     expense.SupplierID,
			MovementTypeID: expense.MovementTypeID,
			Amount:         expense.Amount,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	default:
		return "", fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdateExpense(expense *models.ExpenseUpdate, workplace string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		switch workplace {
		case "laundry":
			if err := tx.Where("id = ?", expense.ID).
				Updates(&models.ExpenseLaundry{
					Details: expense.Details, 
					SupplierID: expense.SupplierID, 
					MovementTypeID: expense.MovementTypeID, 
					Amount: expense.Amount,
					}).Error; err != nil {
				return err
			}
			return nil
		case "workshop":
			if err := r.DB.Model(&models.ExpenseWorkshop{}).
				Where("id = ?", expense.ID).
				Updates(map[string]interface{}{"details": expense.Details, "supplier_id": expense.SupplierID, "movement_type_id": expense.MovementTypeID, "amount": expense.Amount}).Error; err != nil {
				return err
			}
			return nil
		default:
			return fmt.Errorf("tipo de movimiento no soportado")
		}
	})
}

func (r *Repository) DeleteExpenseByID(id string, workplace string) error {
	if workplace == "laundry" {
		if err := r.DB.Where("id = ?", id).Delete(&models.ExpenseLaundry{}).Error; err != nil {
			return err
		}
	} else if workplace == "workshop" {
		if err := r.DB.Where("id = ?", id).Delete(&models.ExpenseWorkshop{}).Error; err != nil {
			return err
		}
	} else {
		return fmt.Errorf("tipo de movimiento no soportado")
	}
	return nil
}