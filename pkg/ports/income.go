package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type IncomeService interface {
	GetIncomeByID(id string) (client *models.Income, err error)
	GetIncomeByName(name string) (clients *[]models.Expense, err error)
	GetAllIncomes() (clients *[]models.Expense, err error)
	CreateIncome(clientCreate *models.IncomeCreate) (id string, err error)
	UpdateIncome(clienUpdate *models.IncomeUpdate) (err error)
	DeleteIncome(id string) error
}

type IncomeRepository interface {
	GetIncomeByID(id string) (client *models.Income, err error)
	GetIncomeByName(name string) (clients *[]models.Income, err error)
	GetAllIncomes() (clients *[]models.Client, err error)
	CreateIncome(clientCreate *models.IncomeCreate) (id string, err error)
	UpdateIncome(clienUpdate *models.IncomeUpdate) (err error)
	DeleteIncome(id string) error
}