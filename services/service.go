package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func ServiceCreate(service *models.ServiceCreate, workplace string) (string, error) {
	exist, err := repositories.Repo.GetServiceByName(service.Name, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al buscar servicio", err)
	}

	if exist {
		return "", models.ErrorResponse(400, "El servicio ya existe", nil)
	}

	id, err := repositories.Repo.CreateService(service, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear servicio", err)
	}
	return id, nil
}

func ServiceUpdate(service *models.ServiceUpdate, workplace string) error {
	err := repositories.Repo.UpdateService(service, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Servicio no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al buscar servicio", err)
	}
	return nil
}

func ServiceDeleteByID(id string, workplace string) error {
	err := repositories.Repo.DeleteServiceByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Servicio no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al buscar servicio", err)
	}
	return nil
}

func ServiceGetAll(workplace string) (*[]models.ServiceLaundry, *[]models.ServiceWorkshop, error) {
	servicesLaundry, servicesWorkshop, err := repositories.Repo.GetAllServices(workplace)
	if err != nil {
		return nil, nil, models.ErrorResponse(500, "Error al obtener servicios", err)
	}
	return servicesLaundry, servicesWorkshop, nil
}

func ServiceGetByID(id string, workplace string) (*models.ServiceLaundry, *models.ServiceWorkshop, error) {
	serviceLaundry, serviceWorkshop, err := repositories.Repo.GetServiceByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Servicio no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al buscar servicio", err)
	}
	return serviceLaundry, serviceWorkshop, nil
}