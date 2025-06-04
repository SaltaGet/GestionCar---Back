package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

// EXPENSE
func (r *TenantRepository) ResumeExpenseGetByDateBetween(fromDate string, toDate string) (*[]models.ResumeExpense, error) {
	return nil, nil
}

func (r *TenantRepository) ResumeExpenseGetByID(id string) (*models.ResumeExpense, error) {
	return nil, nil
}

func (r *TenantRepository) ResumeExpenseCreate(resume *models.ResumeExpenseCreate) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "", err
	}
	return "", nil
}

func (r *TenantRepository) ResumeExpenseUpdate(resume *models.ResumeExpenseUpdate) error {
	return nil
}


// INCOME



func (r *TenantRepository) ResumeIncomeGetByDateBetween(fromDate string, toDate string) (*[]models.ResumeIncome, error) {
	return nil, nil
}

func (r *TenantRepository) ResumeIncomeGetByID(id string) (*models.ResumeIncome, error) {
	return nil, nil
}

func (r *TenantRepository) ResumeIncomeCreate(resume *models.ResumeIncomeCreate) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "",err
	}
	return "", nil
}

func (r *TenantRepository) ResumeIncomeUpdate(resume *models.ResumeIncomeUpdate) error {
	return nil
}
