package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (c *ClientService) ClientCreate(clientCreate *models.ClientCreate) (string, error) {
	client, err := c.ClientRepository.CreateClient(clientCreate)

	if err != nil {
		return "", err
	}

	return client, nil
}

func (c *ClientService) ClientGetAll() (*[]models.Client, error) {
	clients, err := c.ClientRepository.GetAllClients()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los clientes", err)
	}
	return clients, nil
}

func (c *ClientService) ClientGetByID(id string) (*models.Client, error) {
	client, err := c.ClientRepository.GetClientByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Cliente no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return client, nil
}

func (c *ClientService) ClientGetByName(name string) (*[]models.Client, error) {
	client, err := c.ClientRepository.GetClientByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Cliente no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return client, nil
}

func (c *ClientService) ClientUpdate(clientUpdate *models.ClientUpdate) (string, error) {
	err := c.ClientRepository.UpdateClient(clientUpdate)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrorResponse(404, "Cliente no encontrado", err)
		}
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return clientUpdate.ID, nil
}

func (c *ClientService) ClientDelete(id string) (string, error) {
	err := c.ClientRepository.DeleteClient(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrorResponse(404, "Cliente no encontrado", err)
		}
		return "", models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return id, nil
}