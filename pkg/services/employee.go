package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (e *EmployeeService) GetEmployeeByID(id string) (*models.Employee, error) {
	employee, err := e.EmployeeRepository.GetEmployeeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return employee, nil
}

func (e *EmployeeService) GetEmployeeByName(name string) (*[]models.Employee, error) {
	employees, err := e.EmployeeRepository.GetEmployeeByName(name)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener clientes", err)
	}
	return employees, nil
}

func (e *EmployeeService) GetAllEmployees(workplace string) (*[]models.Employee, error) {
	employees, err := e.EmployeeRepository.GetAllEmployees()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return employees, nil
}

func (e *EmployeeService) CreateEmployee(employee *models.EmployeeCreate) (string, error) {
	id, err := e.EmployeeRepository.CreateEmployee(employee)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func (e *EmployeeService) UpdateEmployee(employee *models.EmployeeUpdate) error {
	err := e.EmployeeRepository.UpdateEmployee(employee)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func (e *EmployeeService) DeleteEmployee(id string) error {
	err := e.EmployeeRepository.DeleteEmployee(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}