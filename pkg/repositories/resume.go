package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

// EXPENSE

func (r *Repository) CreateExpenseResume(resume *models.ResumeExpenseCreate) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "", err
	}
	return "", nil
}

func (r *Repository) GetExpenseResumeByDateBetween(fromDate string, toDate string) (*[]models.ResumeExpense, error) {
	return nil, nil
}

func (r *Repository) GetExpenseResumeByID(id string) (*models.ResumeExpense, error) {
	return nil, nil
}

func (r *Repository) UpdateExpenseResume(resume *models.ResumeExpense) error {
	return nil
}


// INCOME

func (r *Repository) CreateIncomeResume(resume *models.ResumeIncome) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "",err
	}
	return "", nil
}

func (r *Repository) GetIncomeResumeByDateBetween(fromDate string, toDate string) (*[]models.ResumeIncome, error) {
	return nil, nil
}

func (r *Repository) GetIncomeResumeByID(id string) (*models.ResumeIncome, error) {
	return nil, nil
}

func (r *Repository) UpdateIncomeResume(resume *models.ResumeIncome) error {
	return nil
}
