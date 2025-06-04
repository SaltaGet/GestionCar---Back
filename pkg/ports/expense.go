package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type ExpenseRepository interface {
	ExpenseGetByID(id string) (expense *models.Expense, err error)
	ExpenseGetByName(name string) (expenses *[]models.Expense, err error)
	ExpenseGetAll() (expenses *[]models.Expense, err error)
	ExpenseGetToday() (expenses *[]models.Expense, err error)
	ExpenseCreate(expenseCreate *models.ExpenseCreate) (id string, err error)
	ExpenseUpdate(expenseUpdate *models.ExpenseUpdate) (err error)
	ExpenseDelete(id string) error
}

type ExpenseService interface {
	ExpenseGetByID(id string) (expense *models.Expense, err error)
	ExpenseGetByName(name string) (expenses *[]models.Expense, err error)
	ExpenseGetAll() (expenses *[]models.Expense, err error)
	ExpenseGetToday() (expenses *[]models.Expense, err error)
	ExpenseCreate(expenseCreate *models.ExpenseCreate) (id string, err error)
	ExpenseUpdate(expenseUpdate *models.ExpenseUpdate) (err error)
	ExpenseDelete(id string) error
}