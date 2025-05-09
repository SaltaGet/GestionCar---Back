package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
)

func (r *Repository) GetServiceByID(id string, workplace string) (*models.ServiceLaundry, *models.ServiceWorkshop, error) {
	if workplace == "laundry" {
		var service models.ServiceLaundry
		if err := r.DB.Where("id = ?", id).First(&service).Error; err != nil {
			return nil, nil, err
		}
		return &service, nil, nil
	} else if workplace == "workshop" {
		var service models.ServiceWorkshop
		if err := r.DB.Where("id = ?", id).First(&service).Error; err != nil {
			return nil, nil, err
		}
		return nil, &service, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetServiceByName(name string, workplace string) (*models.ServiceLaundry, *models.ServiceWorkshop, error) {
	if workplace == "laundry" {
		var service models.ServiceLaundry
		if err := r.DB.Where("name = ?", name).First(&service).Error; err != nil {
			return nil, nil, err
		}
		return &service, nil, nil
	} else if workplace == "workshop" {
		var service models.ServiceWorkshop
		if err := r.DB.Where("name = ?", name).First(&service).Error; err != nil {
			return nil, nil, err
		}
		return nil, &service, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetAllServices(workplace string) (*[]models.ServiceLaundry, *[]models.ServiceWorkshop, error) {
	if workplace == "laundry" {
		var services []models.ServiceLaundry
		if err := r.DB.Find(&services).Error; err != nil {
			return nil, nil, err
		}
		return &services, nil, nil
	} else if workplace == "workshop" {
		var services []models.ServiceWorkshop
		if err := r.DB.Find(&services).Error; err != nil {
			return nil, nil, err
		}
		return nil, &services, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) CreateService(service *models.ServiceCreate, workplace string) (string, error) {
	newID := uuid.NewString()
	switch workplace {
	case "laundry":
		if err := r.DB.Create(models.ServiceLaundry{
			ID: newID,
			Name: service.Name,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	case "workshop":
		if err := r.DB.Create(models.ServiceWorkshop{
			ID: newID,
			Name: service.Name,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	default:
		return "", fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdateService(service *models.ServiceUpdate, workplace string) error {
	switch workplace {
	case "laundry":
		if err := r.DB.Where("id = ?", service.ID).First(&models.ServiceLaundry{}).Error; err != nil {
			return err
		}
		s := models.ServiceLaundry{
			ID: service.ID,
			Name: service.Name,
		}
		if err := r.DB.Save(&s).Error; err != nil {
			return err
		}
	case "workshop":
		if err := r.DB.Where("id = ?", service.ID).First(&models.ServiceWorkshop{}).Error; err != nil {
			return err
		}
		s := models.ServiceWorkshop{
			ID: service.ID,
			Name: service.Name,
		}
		if err := r.DB.Save(&s).Error; err != nil {
			return err
		}
	default:
		return fmt.Errorf("tipo de movimiento no soportado")
	}
	return nil
}

func (r *Repository) DeleteServiceByID(id string, workplace string) error {
	if workplace == "laundry" {
		if err := r.DB.Where("id = ?", id).Delete(&models.ServiceLaundry{}).Error; err != nil {
			return err
		}
	} else if workplace == "workshop" {
		if err := r.DB.Where("id = ?", id).Delete(&models.ServiceWorkshop{}).Error; err != nil {
			return err
		}
	} else {
		return fmt.Errorf("tipo de movimiento no soportado")
	}
	return nil
}