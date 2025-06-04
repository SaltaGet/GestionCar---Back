package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type IncomeRepository interface {
	IncomeGetByID(id string) (income *models.Income, err error)
	// IncomeGetByName(name string) (incomes *[]models.Expense, err error)
	IncomeGetAll() (incomes *[]models.Income, err error)
	IncomeGetToday() (incomes *[]models.Income, err error)
	IncomeCreate(incomeCreate *models.IncomeCreate) (id string, err error)
	IncomeUpdate(incomeUpdate *models.IncomeUpdate) (err error)
	IncomeDelete(id string) error
}

type IncomeService interface {
	IncomeGetByID(id string) (income *models.Income, err error)
	// IncomeGetByName(name string) (incomes *[]models.Income, err error)
	IncomeGetAll() (incomes *[]models.Income, err error)
	IncomeGetToday() (incomes *[]models.Income, err error)
	IncomeCreate(incomeCreate *models.IncomeCreate) (id string, err error)
	IncomeUpdate(incomeUpdate *models.IncomeUpdate) (err error)
	IncomeDelete(id string) error
}
// type IncomeService interface {
// 	GetIncomeByID(id string) (client *models.Income, err error)
// 	GetIncomeByName(name string) (clients *[]models.Expense, err error)
// 	GetAllIncomes() (clients *[]models.Expense, err error)
// 	CreateIncome(clientCreate *models.IncomeCreate) (id string, err error)
// 	UpdateIncome(clienUpdate *models.IncomeUpdate) (err error)
// 	DeleteIncome(id string) error
// }

// type IncomeRepository interface {
// 	GetIncomeByID(id string) (income *models.Income, err error)
// 	GetIncomeByName(name string) (incomes *[]models.Income, err error)
// 	GetAllIncomes() (incomes *[]models.Income, err error)
// 	GetIncomeToday() (incomes *[]models.Income, err error)
// 	CreateIncome(incomeCreate *models.IncomeCreate) (id string, err error)
// 	UpdateIncome(incomeUpdate *models.IncomeUpdate) (err error)
// 	DeleteIncome(id string) error
// }