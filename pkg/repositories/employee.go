package repositories

import (

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
)

func (r *TenantRepository) GetEmployeeByID(id string) (*models.Employee, error) {
	var employee models.Employee
	if err := r.DB.Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *TenantRepository) GetAllEmployees() (*[]models.Employee, error) {
	var employees []models.Employee
	if err := r.DB.Find(&employees).Error; err != nil {
		return nil, err
	}
	return &employees, nil
}

func (r *TenantRepository) CreateEmployee(employee *models.EmployeeCreate) (string, error) {
	newID := uuid.NewString()
	if err := r.DB.Create(&models.Employee{
		ID:      newID,
		Name:    employee.Name,
		Phone:   employee.Phone,
		Email:   employee.Email,
		Address: employee.Address,
	}).Error; err != nil {
		return "", err
	}
	return newID, nil
}

func (r *TenantRepository) UpdateEmployee(employeeUpdate *models.EmployeeUpdate) error {
	var employee models.Employee
	if err := r.DB.Where("id = ?", employeeUpdate.ID).First(&employee).Error; err != nil {
		return err
	}
	employee.Name = employeeUpdate.Name
	employee.Phone = employeeUpdate.Phone
	employee.Email = employeeUpdate.Email
	employee.Address = employeeUpdate.Address
	if err := r.DB.Save(&employee).Error; err != nil {
		return err
	}
	return nil
}

func (r *TenantRepository) DeleteEmployee(id string) error {
	var employee models.Employee
	if err := r.DB.Where("id = ?", id).Delete(&employee).Error; err != nil {
		return err
	}
	return nil
}

func (r *TenantRepository) GetEmployeeByName(name string) (*[]models.Employee, error) {
	var employees []models.Employee
	if err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&employees).Error; err != nil {
		return nil, err
	}
	return &employees, nil
}
