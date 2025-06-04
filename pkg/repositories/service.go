package repositories

import (
	"errors"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *TenantRepository) ServiceGetByID(id string) (*models.Service, error) {
	var service models.Service
	if err := r.DB.Where("id = ?", id).First(&service).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *TenantRepository) ServiceExistByName(name string) (bool, error) {
	var service models.Service
	if err := r.DB.Where("name = ?", name).First(&service).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *TenantRepository) ServiceGetByName(name string) (*[]models.Service, error) {
	var services []models.Service
	if err := r.DB.Limit(5).Where("name LIKE ?", "%"+name+"%").Find(&services).Error; err != nil {
		return nil, err
	}

	return &services, nil
}

func (r *TenantRepository) ServiceGetAll() (*[]models.Service, error) {
	var services []models.Service
	if err := r.DB.Find(&services).Error; err != nil {
		return nil, err
	}
	return &services, nil
}

func (r *TenantRepository) ServiceCreate(service *models.ServiceCreate) (string, error) {
	newID := uuid.NewString()
	if err := r.DB.Create(&models.Service{
		ID:   newID,
		Name: service.Name,
	}).Error; err != nil {
		return "", err
	}
	return newID, nil
}

func (r *TenantRepository) ServiceUpdate(service *models.ServiceUpdate) error {
	if err := r.DB.Where("id = ?", service.ID).First(&models.Service{}).Error; err != nil {
		return err
	}
	s := models.Service{
		ID:   service.ID,
		Name: service.Name,
	}
	if err := r.DB.Save(&s).Error; err != nil {
		return err
	}
	return nil
}

func (r *TenantRepository) ServiceDelete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Service{}).Error; err != nil {
		return err
	}
	return nil
}
