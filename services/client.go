package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ClientCreate(clientCreate *models.ClientCreate) (string, error) {
	client, err := repositories.Repo.CreateClient(&models.Client{
		ID: uuid.NewString(),
		FirstName: clientCreate.FirstName,
		LastName:  clientCreate.LastName,		
		CUIL:      clientCreate.CUIL,
		DNI:       clientCreate.DNI,
		Email:     clientCreate.Email,
	})

	if err != nil {
		return "", err
	}

	return client, nil
}

func ClientGetAll() (*[]models.Client, error) {
	clients, err := repositories.Repo.GetAllClients()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los clientes", err)
	}
	return &clients, nil
}

func ClientGetByID(id string) (*models.Client, error) {
	client, err := repositories.Repo.GetClientByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Cliente no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return client, nil
}

func ClientGetByName(name string) (*[]models.Client, error) {
	client, err := repositories.Repo.GetClientByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Cliente no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return client, nil
}

func ClientUpdate(clientUpdate *models.ClientUpdate) (string, error) {
	err := repositories.Repo.UpdateClient(&models.Client{
		ID: clientUpdate.ID,
		FirstName: clientUpdate.FirstName,
		LastName:  clientUpdate.LastName,		
		CUIL:      clientUpdate.CUIL,
		DNI:       clientUpdate.DNI,
		Email:     clientUpdate.Email,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrorResponse(404, "Cliente no encontrado", err)
		}
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return clientUpdate.ID, nil
}

func ClientDelete(id string) (string, error) {
	err := repositories.Repo.DeleteClient(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrorResponse(404, "Cliente no encontrado", err)
		}
		return "", models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return id, nil
}