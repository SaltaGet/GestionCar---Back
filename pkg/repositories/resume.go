package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

// EXPENSE

func (r *Repository) CreateExpenseResume(resume *models.ExpenseResumeCreate, workplace string) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "", err
	}
	return "", nil
}

func (r *Repository) GetExpenseResumeByDateBetween(fromDate string, toDate string, workplace string) (*[]models.ExpenseResumeLaundry, *[]models.ExpenseResumeWorkshop, error) {
	return nil, nil, nil
}

func (r *Repository) GetExpenseResumeByID(id string, workplace string) (*models.ExpenseResumeLaundry, *models.ExpenseResumeWorkshop, error) {
	return nil, nil, nil
}

func (r *Repository) UpdateExpenseResume(resume *models.ExpenseResumeUpdate, workplace string) error {
	return nil
}


// INCOME

func (r *Repository) CreateIncomeResume(resume *models.IncomeResumeCreate, workplace string) (string, error) {
	err := r.DB.Create(&resume).Error
	if err != nil {
		return "",err
	}
	return "", nil
}

func (r *Repository) GetIncomeResumeByDateBetween(fromDate string, toDate string, workplace string) (*[]models.IncomeResumeLaundry, *[]models.IncomeResumeWorkshop, error) {
	return nil, nil, nil
}

func (r *Repository) GetIncomeResumeByID(id string, workplace string) (*models.IncomeResumeLaundry, *models.IncomeResumeWorkshop, error) {
	return nil, nil, nil
}

func (r *Repository) UpdateIncomeResume(resume *models.IncomeResumeUpdate, workplace string) error {
	return nil
}
