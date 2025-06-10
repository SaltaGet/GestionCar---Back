package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type IncomeRepository interface {
	IncomeGetByID(id string) (income *models.Income, err error)
	IncomeGetAll() (incomes *[]models.Income, err error)
	IncomeGetToday() (incomes *[]models.Income, err error)
	IncomeCreate(incomeCreate *models.IncomeCreate) (id string, err error)
	IncomeUpdate(incomeUpdate *models.IncomeUpdate) (err error)
	IncomeDelete(id string) error
}

type IncomeService interface {
	IncomeGetByID(id string) (income *models.Income, err error)
	IncomeGetAll() (incomes *[]models.Income, err error)
	IncomeGetToday() (incomes *[]models.Income, err error)
	IncomeCreate(incomeCreate *models.IncomeCreate) (id string, err error)
	IncomeUpdate(incomeUpdate *models.IncomeUpdate) (err error)
	IncomeDelete(id string) error
}
