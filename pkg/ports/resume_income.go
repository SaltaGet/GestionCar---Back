package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type ResumeIncomeService interface {
	GetIncomeResumeByID(id string) (resume *models.ResumeIncome, err error)
	IncomeResumeCreate(resume *models.ResumeIncomeCreate) (id string, err error)
	GetIncomeResumeByDateBetween(fromDate string, toDate string, workplace string) (resumes *[]models.ResumeIncome, err error)
	UpdateIncomeResume(resume *models.ResumeIncomeUpdate) (err error)
}

type ResumeIncomeRepository interface {
	GetIncomeResumeByID(id string) (resume *models.ResumeIncome, err error)
	IncomeResumeCreate(resume *models.ResumeIncomeCreate) (id string, err error)
	GetIncomeResumeByDateBetween(fromDate string, toDate string) (resumes *[]models.ResumeIncome, err error)
	UpdateIncomeResume(resume *models.ResumeIncomeUpdate) (err error)
}