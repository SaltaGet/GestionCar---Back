package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

// EXPENSE

func (r *TenantRepository) CreateExpenseResume(resume *models.ResumeExpenseCreate) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "", err
	}
	return "", nil
}

func (r *TenantRepository) GetExpenseResumeByDateBetween(fromDate string, toDate string) (*[]models.ResumeExpense, error) {
	return nil, nil
}

func (r *TenantRepository) GetExpenseResumeByID(id string) (*models.ResumeExpense, error) {
	return nil, nil
}

func (r *TenantRepository) UpdateExpenseResume(resume *models.ResumeExpense) error {
	return nil
}


// INCOME

func (r *TenantRepository) CreateIncomeResume(resume *models.ResumeIncome) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "",err
	}
	return "", nil
}

func (r *TenantRepository) GetIncomeResumeByDateBetween(fromDate string, toDate string) (*[]models.ResumeIncome, error) {
	return nil, nil
}

func (r *TenantRepository) GetIncomeResumeByID(id string) (*models.ResumeIncome, error) {
	return nil, nil
}

func (r *TenantRepository) UpdateIncomeResume(resume *models.ResumeIncome) error {
	return nil
}
