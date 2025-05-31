package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type ResumeExpenseService interface {
	GetExpenseResumeByID(id string) (resume *models.ResumeExpense, err error)
	ExpenseResumeCreate(resume *models.ResumeExpenseCreate) (id string, err error)
	GetExpenseResumeByDateBetween(fromDate string, toDate string) (resumes *[]models.ResumeExpense, err error)
	UpdateExpenseResume(resume *models.ResumeExpenseUpdate) (err error)
}

type ResumeExpenseRepository interface {
	GetExpenseResumeByID(id string) (resume *models.ResumeExpense, err error)
	CreateExpenseResume(resume *models.ResumeExpenseCreate) (id string, err error)
	GetExpenseResumeByDateBetween(fromDate string, toDate string) (resumes *[]models.ResumeExpense, err error)
	UpdateExpenseResume(resume *models.ResumeExpenseUpdate) (err error)
}


