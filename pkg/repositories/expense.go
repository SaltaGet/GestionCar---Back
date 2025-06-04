package repositories

import (
	"time"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *TenantRepository) GetExpenseByID(id string) (*models.Expense, error) {
		var expense models.Expense
		if err := r.DB.Where("id = ?", id).First(&expense).Error; err != nil {
			return nil, err
		}
		return &expense, nil
}

func (r *TenantRepository) GetAllExpenses() (*[]models.Expense, error) {
		var expenses []models.Expense
		if err := r.DB.Limit(100).Order("created_at desc").Find(&expenses).Error; err != nil {
			return nil, err
		}
		return &expenses, nil
}

func (r *TenantRepository) GetExpenseToday() (*[]models.Expense, error) {
    today := time.Now().Format("2006-01-02")
        var expenses []models.Expense
        if err := r.DB.Where("DATE(created_at) = ?", today).Order("created_at desc").Find(&expenses).Error; err != nil {
            return nil, err
        }
        return &expenses, nil
}

func (r *TenantRepository) CreateExpense(expense *models.ExpenseCreate) (string, error) {
	newID := uuid.NewString()
		if err := r.DB.Create(&models.Expense{
			ID:             newID,
			Details:        expense.Details,
			SupplierID:     expense.SupplierID,
			MovementTypeID: expense.MovementTypeID,
			Amount:         expense.Amount,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
}

func (r *TenantRepository) UpdateExpense(expense *models.ExpenseUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("id = ?", expense.ID).
				Updates(&models.Expense{
					Details: expense.Details, 
					SupplierID: expense.SupplierID, 
					MovementTypeID: expense.MovementTypeID, 
					Amount: expense.Amount,
					}).Error; err != nil {
				return err
			}
			return nil
	})
}

func (r *TenantRepository) DeleteExpenseByID(id string) error {
		if err := r.DB.Where("id = ?", id).Delete(&models.Expense{}).Error; err != nil {
			return err
		}
	return nil
}