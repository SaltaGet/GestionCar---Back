package repositories

import (
	"github.com/DanielChachagua/GestionCar/models"
)

func (r *Repository) GetClientByID(id string) (*models.Client, error) {
	var client models.Client
	if err := r.DB.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *Repository) GetClientByName(name string) (*[]models.Client, error) {
	var client []models.Client
	if err := r.DB.Where("last_name LIKE ? OR first_name LIKE ?", "%"+name+"%", "%"+name+"%").Find(&client).Error; err != nil {
    return nil, err
	}
	return &client, nil
}

func (r *Repository) GetAllClients() ([]models.Client, error) {
	var clients []models.Client
	if err := r.DB.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *Repository) CreateClient(client *models.Client) (string, error) {
	if err := r.DB.Create(client).Error; err != nil {
		return "", err
	}
	return client.ID, nil
}

func (r *Repository) UpdateClient(client *models.Client) error {
	var existing models.Client
	if err := r.DB.First(&existing, "id = ?", client.ID).Error; err != nil {
			return err 
	}

	if err := r.DB.Save(client).Error; err != nil {
			return err
	}
	return nil
}

func (r *Repository) DeleteClient(id string) error {
	if err := r.DB.Where("client_id = ?", id).Delete(&models.Vehicle{}).Error; err != nil {
			return err
	}
	if err := r.DB.Where("id = ?", id).Delete(&models.Client{}).Error; err != nil {
			return err
	}
	return nil
}
