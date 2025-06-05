package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (s *ServiceService) ServiceGetByID(id string) (*models.Service, error) {
	service, err := s.ServiceRepository.ServiceGetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Servicio no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar servicio", err)
	}
	return service, nil
}

func (s *ServiceService) ServiceGetByName(name string) (*[]models.Service, error) {
	services, err := s.ServiceRepository.ServiceGetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Servicio no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar servicio", err)
	}
	return services, nil
}

func (s *ServiceService) ServiceGetAll() (*[]models.Service, error) {
	services, err := s.ServiceRepository.ServiceGetAll()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener servicios", err)
	}
	return services, nil
}

func (s *ServiceService) ServiceCreate(service *models.ServiceCreate) (string, error) {
	exist, err := s.ServiceRepository.ServiceExistByName(service.Name)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al buscar servicio", err)
	}

	if exist {
		return "", models.ErrorResponse(400, "El servicio ya existe", nil)
	}

	id, err := s.ServiceRepository.ServiceCreate(service)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear servicio", err)
	}
	return id, nil
}

func (s *ServiceService) ServiceUpdate(service *models.ServiceUpdate) error {
	err := s.ServiceRepository.ServiceUpdate(service)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Servicio no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al buscar servicio", err)
	}
	return nil
}

func (s *ServiceService) ServiceDelete(id string) error {
	err := s.ServiceRepository.ServiceDelete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Servicio no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al buscar servicio", err)
	}
	return nil
}