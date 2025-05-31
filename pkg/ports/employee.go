package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type EmployeeService interface {
	GetEmployeeByID(id string) (client *models.Employee, err error)
	GetEmployeeByName(name string) (clients *[]models.Employee, err error)
	GetAllEmployees() (clients *[]models.Employee, err error)
	CreateEmployee(clientCreate *models.EmployeeCreate) (id string, err error)
	UpdateEmployee(clienUpdate *models.EmployeeUpdate) (err error)
	DeleteEmployee(id string) error
}

type EmployeeRepository interface {
	GetEmployeeByID(id string) (client *models.Employee, err error)
	GetEmployeeByName(name string) (clients *[]models.Employee, err error)
	GetAllEmployees() (clients *[]models.Employee, err error)
	CreateEmployee(clientCreate *models.EmployeeCreate) (id string, err error)
	UpdateEmployee(clienUpdate *models.EmployeeUpdate) (err error)
	DeleteEmployee(id string) error
}
