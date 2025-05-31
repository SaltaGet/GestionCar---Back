package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

// v VehicleService

func (v *VehicleService) VehicleCreate(vehicleCreate *models.VehicleCreate) (string , error) {
	exist, err := v.VehicleRepository.VehicleExistByDomain(vehicleCreate.Domain)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al buscar el vehiculo", err)
	}

	if exist {
		return "", models.ErrorResponse(400, "El dominio ya existe", nil)
	}

	vehicle, err := v.VehicleRepository.VehicleCreate(vehicleCreate)

	if err != nil {
		models.ErrorResponse(500, "Error al crear el vehiculo", err)
	}

	return vehicle, nil
}

func (v *VehicleService) VehicleGetAll() (*[]models.Vehicle, error) {
	vehicles, err := v.VehicleRepository.VehicleGetAll()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los vehiculos", err)
	}
	return vehicles, nil
}

func (v *VehicleService) VehicleGetByID(id string) (*models.Vehicle, error) {
	vehicle, err := v.VehicleRepository.VehicleGetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Vehiculo no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar usuario", err)
	}
	return vehicle, nil
}

func (v *VehicleService) VehicleGetByDomain(domain string) (*[]models.Vehicle, error) {
	vehicle, err := v.VehicleRepository.VehicleGetByDomain(domain)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Vehiculo no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar usuario", err)
	}
	return vehicle, nil
}

func (v *VehicleService) VehicleGetByClientID(clientID string) (*[]models.Vehicle, error) {
	vehicles, err := v.VehicleRepository.VehicleGetByClientID(clientID)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los vehiculos", err)
	}
	return vehicles, nil
}

func (v *VehicleService) VehicleUpdate(vehicleUpdate *models.VehicleUpdate) error {
	err := v.VehicleRepository.VehicleUpdate(vehicleUpdate)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Vehiculo no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return nil
}

func (v *VehicleService) VehicleDelete(id string) (error) {
	err := v.VehicleRepository.VehicleDelete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Vehiculo no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return nil
}
