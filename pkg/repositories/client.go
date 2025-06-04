package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func (r *TenantRepository) ClientGetByID(id string) (*models.Client, error) {
	var client models.Client
	if err := r.DB.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *TenantRepository) ClientGetByName(name string) (*[]models.Client, error) {
	var client []models.Client
	if err := r.DB.Where("last_name LIKE ? OR first_name LIKE ?", "%"+name+"%", "%"+name+"%").Find(&client).Error; err != nil {
    return nil, err
	}
	return &client, nil
}

func (r *TenantRepository) ClientGetAll() ([]models.Client, error) {
	var clients []models.Client
	if err := r.DB.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *TenantRepository) ClientCreate(client *models.Client) (string, error) {
	if err := r.DB.Create(client).Error; err != nil {
		return "", err
	}
	return client.ID, nil
}

func (r *TenantRepository) ClientUpdate(client *models.Client) error {
	var existing models.Client
	if err := r.DB.First(&existing, "id = ?", client.ID).Error; err != nil {
			return err 
	}

	if err := r.DB.Save(client).Error; err != nil {
			return err
	}
	return nil
}

func (r *TenantRepository) ClientDelete(id string) error {
	if err := r.DB.Where("client_id = ?", id).Delete(&models.Vehicle{}).Error; err != nil {
			return err
	}
	if err := r.DB.Where("id = ?", id).Delete(&models.Client{}).Error; err != nil {
			return err
	}
	return nil
}
