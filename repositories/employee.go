package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
)

func (r *Repository) GetEmployeeByID(id string, workplace string) (*models.EmployeeLaundry, *models.EmployeeWorkshop, error) {
	if workplace == "laundry" {
		var employee models.EmployeeLaundry
		if err := r.DB.Where("id = ?", id).First(&employee).Error; err != nil {
			return nil, nil, err
		}
		return &employee, nil, nil
	} else if workplace == "workshop" {
		var employee models.EmployeeWorkshop
		if err := r.DB.Where("id = ?", id).First(&employee).Error; err != nil {
			return nil, nil, err
		}
		return nil, &employee, nil
	}
	return nil, nil, fmt.Errorf("tipo de espacio no soportado")
}

func (r *Repository) GetAllEmployees(workplace string) ([]models.EmployeeLaundry, []models.EmployeeWorkshop, error) {
	if workplace == "laundry" {
		var employees []models.EmployeeLaundry
		if err := r.DB.Find(&employees).Error; err != nil {
			return nil, nil, err
		}
		return employees, nil, nil
	} else	if workplace == "workshop" {
		var employees []models.EmployeeWorkshop
		if err := r.DB.Find(&employees).Error; err != nil {
			return nil, nil, err
		}
		return nil, employees, nil
	}
	return nil, nil, fmt.Errorf("tipo de espacio no soportado")
}

func (r *Repository) CreateEmployee(employee *models.EmployeeCreate, workplace string) (string, error) {
	newID := uuid.NewString()
	switch workplace{
	case "laundry":
		if err := r.DB.Create(&models.EmployeeLaundry{
			ID: newID,
			Name: employee.Name,
			Phone: employee.Phone,
			Email: employee.Email,
			Address: employee.Address,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	case "workshop":
		if err := r.DB.Create(&models.EmployeeWorkshop{
			ID: newID,
			Name: employee.Name,
			Phone: employee.Phone,
			Email: employee.Email,
			Address: employee.Address,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	default:
		return "", fmt.Errorf("tipo de empleado no soportado")
	}
}

func (r *Repository) UpdateEmployee(employeeUpdate *models.EmployeeUpdate, workplace string) error {
	switch workplace{
	case "laundry":
		var employee models.EmployeeLaundry
		if err := r.DB.Where("id = ?", employee.ID).First(&employee).Error; err != nil {
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
	case "workshop":
		var employee models.EmployeeWorkshop
		if err := r.DB.Where("id = ?", employee.ID).First(&employee).Error; err != nil {
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
	default:
		return fmt.Errorf("tipo de empleado no soportado")
	}
}

func (r *Repository) DeleteEmployee(id string, workplace string) error {
	if workplace == "laundry" {
		var employee models.EmployeeLaundry
		if err := r.DB.Where("id = ?", id).Delete(&employee).Error; err != nil {
			return err
		}
	} else if workplace == "workshop" {
		var employee models.EmployeeWorkshop
		if err := r.DB.Where("id = ?", id).Delete(&employee).Error; err != nil {
			return err
		}
	} else {
		return fmt.Errorf("tipo de espacio no soportado")
	}
	return nil
}

func (r *Repository) GetEmployeeByName(name string, workplace string) (*[]models.EmployeeLaundry, *[]models.EmployeeWorkshop, error) {
	if workplace == "laundry" {
		var employees []models.EmployeeLaundry
		if err := r.DB.Where("last_name LIKE ? OR first_name LIKE ?", "%"+name+"%", "%"+name+"%").Find(&employees).Error; err != nil {
			return nil, nil, err
		}
		return &employees, nil, nil
	} else	if workplace == "workshop" {
		var employees []models.EmployeeWorkshop
		if err := r.DB.Where("last_name LIKE ? OR first_name LIKE ?", "%"+name+"%", "%"+name+"%").Find(&employees).Error; err != nil {
			return nil, nil, err
		}
		return nil, &employees, nil
	}
	return nil, nil, fmt.Errorf("tipo de espacio no soportado")
}