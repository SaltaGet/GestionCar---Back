package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

// EXPENSE
func (r *ResumeRepository) ResumeExpenseGetByDateBetween(fromDate string, toDate string) (*[]models.ResumeExpense, error) {
	return nil, nil
}

func (r *ResumeRepository) ResumeExpenseGetByID(id string) (*models.ResumeExpense, error) {
	return nil, nil
}

func (r *ResumeRepository) ResumeExpenseCreate(resume *models.ResumeExpenseCreate) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "", err
	}
	return "", nil
}

func (r *ResumeRepository) ResumeExpenseUpdate(resume *models.ResumeExpenseUpdate) error {
	return nil
}


// INCOME



func (r *ResumeRepository) ResumeIncomeGetByDateBetween(fromDate string, toDate string) (*[]models.ResumeIncome, error) {
	return nil, nil
}

func (r *ResumeRepository) ResumeIncomeGetByID(id string) (*models.ResumeIncome, error) {
	return nil, nil
}

func (r *ResumeRepository) ResumeIncomeCreate(resume *models.ResumeIncomeCreate) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "",err
	}
	return "", nil
}

func (r *ResumeRepository) ResumeIncomeUpdate(resume *models.ResumeIncomeUpdate) error {
	return nil
}
