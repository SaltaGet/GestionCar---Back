package repositories

import (
	"errors"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetServiceByID(id string) (*models.Service, error) {
		var service models.Service
		if err := r.DB.Where("id = ?", id).First(&service).Error; err != nil {
			return nil, err
		}
		return &service, nil
}

func (r *Repository) GetServiceByName(name string) (bool, error) {
		var service models.Service
		if err := r.DB.Where("name = ?", name).First(&service).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return false, nil
			}
			return false, err
		}
		return true, nil
}

func (r *Repository) GetAllServices() (*[]models.Service, error) {
		var services []models.Service
		if err := r.DB.Find(&services).Error; err != nil {
			return nil, err
		}
		return &services, nil
}

func (r *Repository) CreateService(service *models.ServiceCreate) (string, error) {
	newID := uuid.NewString()
		if err := r.DB.Create(&models.Service{
			ID: newID,
			Name: service.Name,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
}

func (r *Repository) UpdateService(service *models.ServiceUpdate) error {
		if err := r.DB.Where("id = ?", service.ID).First(&models.Service{}).Error; err != nil {
			return err
		}
		s := models.Service{
			ID: service.ID,
			Name: service.Name,
		}
		if err := r.DB.Save(&s).Error; err != nil {
			return err
		}
		return nil
	}

func (r *Repository) DeleteServiceByID(id string, workplace string) error {
		if err := r.DB.Where("id = ?", id).Delete(&models.Service{}).Error; err != nil {
			return err
		}
	return nil
}