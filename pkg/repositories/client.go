package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (r *TenantRepository) ClientGetAll() (*[]models.Client, error) {
	var clients []models.Client
	if err := r.DB.Find(&clients).Error; err != nil {
		return nil, err
	}
	return &clients, nil
}

func (r *TenantRepository) ClientCreate(client *models.ClientCreate) (string, error) {
	newClient := models.Client{
		ID: uuid.NewString(),
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Cuil:      client.Cuil,
		Dni:       client.Dni,
		Email:     client.Email,
	}
	if err := r.DB.Create(&newClient).Error; err != nil {
		return "", err
	}
	return newClient.ID, nil
}

func (r *TenantRepository) ClientUpdate(client *models.ClientUpdate) error {
	if err := r.DB.Where("id = ?", client.ID).Updates(&models.Client{
		FirstName: client.FirstName,
		LastName:  client.LastName,
		Cuil:      client.Cuil,
		Dni:       client.Dni,
		Email:     client.Email,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Cliente no encontrado", err)
		}	
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
