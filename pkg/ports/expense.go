package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type ExpenseService interface {
	GetExpenseByID(id string) (client *models.Expense, err error)
	GetExpenseByName(name string) (clients *[]models.Expense, err error)
	GetAllExpenses() (clients *[]models.Expense, err error)
	CreateExpense(clientCreate *models.ExpenseCreate) (id string, err error)
	UpdateExpense(clienUpdate *models.ExpenseUpdate) (err error)
	DeleteExpense(id string) error
}

type ExpenseRepository interface {
	GetExpenseByID(id string) (client *models.Expense, err error)
	GetExpenseByName(name string) (clients *[]models.Expense, err error)
	GetAllExpenses() (expenses *[]models.Expense, err error)
	GetExpenseToday() (expenses *[]models.Expense, err error)
	CreateExpense(clientCreate *models.ExpenseCreate) (id string, err error)
	UpdateExpense(clienUpdate *models.ExpenseUpdate) (err error)
	DeleteExpense(id string) error
}