package controllers

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func GetEmployeeByID(id string, workplace string) (*models.EmployeeLaundry, *models.EmployeeWorkshop, error) {
	employeeLaundry, employeeWorkshop, err := repositories.Repo.GetEmployeeByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return employeeLaundry, employeeWorkshop, nil
}

func GetEmployeeByName(name string, workplace string) (*[]models.EmployeeLaundry, *[]models.EmployeeWorkshop, error) {
	employeesLaundry, employeesWorkshop, err := repositories.Repo.GetEmployeeByName(name, workplace)
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return employeesLaundry, employeesWorkshop, nil
}

func GetAllEmployees(workplace string) ([]models.EmployeeLaundry, []models.EmployeeWorkshop, error) {
	employeesLaundry, employeesWorkshop, err := repositories.Repo.GetAllEmployees(workplace)
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return employeesLaundry, employeesWorkshop, nil
}

func CreateEmployee(employee *models.EmployeeCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreateEmployee(employee, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func UpdateEmployee(employee *models.EmployeeUpdate, workplace string) error {
	err := repositories.Repo.UpdateEmployee(employee, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func DeleteEmployee(id string, workplace string) error {
	err := repositories.Repo.DeleteEmployee(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}