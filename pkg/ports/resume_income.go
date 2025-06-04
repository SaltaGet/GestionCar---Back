package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type ResumeIncomeService interface {
	ResumeIncomeGetByID(id string) (resume *models.ResumeIncome, err error)
	ResumeIncomeCreate(resume *models.ResumeIncomeCreate) (id string, err error)
	ResumeIncomeGetByDateBetween(fromDate string, toDate string, workplace string) (resumes *[]models.ResumeIncome, err error)
	ResumeIncomeUpdate(resume *models.ResumeIncomeUpdate) (err error)
}

type ResumeIncomeRepository interface {
	ResumeIncomeGetByID(id string) (resume *models.ResumeIncome, err error)
	ResumeIncomeCreate(resume *models.ResumeIncomeCreate) (id string, err error)
	ResumeIncomeGetByDateBetween(fromDate string, toDate string) (resumes *[]models.ResumeIncome, err error)
	ResumeIncomeUpdate(resume *models.ResumeIncomeUpdate) (err error)
}