package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func VehicleCreate(vehicleCreate *models.VehicleCreate) (string , error) {
	exist, err := repositories.Repo.GetVehicleByDomain(vehicleCreate.Domain)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al buscar el vehiculo", err)
	}

	if exist != nil {
		return "", models.ErrorResponse(400, "El vehiculo ya existe", nil)
	}

	vehicle, err := repositories.Repo.CreateVehicle(&models.Vehicle{
		ID: uuid.NewString(),
		Domain:   vehicleCreate.Domain,
		Brand:    vehicleCreate.Brand,
		Model:    vehicleCreate.Model,
		Color:    vehicleCreate.Color,
		Year:     vehicleCreate.Year,
		ClientID: vehicleCreate.ClientID,
	})

	if err != nil {
		models.ErrorResponse(500, "Error al crear el vehiculo", err)
	}

	return vehicle, nil
}

func VehicleGetAll() (*[]models.Vehicle, error) {
	vehicles, err := repositories.Repo.GetAllVehicles()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los vehiculos", err)
	}
	return &vehicles, nil
}

func VehicleGetByID(id string) (*models.Vehicle, error) {
	vehicle, err := repositories.Repo.GetVehicleByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar usuario", err)
	}
	return vehicle, nil
}

func VehicleGetByDomain(domain string) (*[]models.Vehicle, error) {
	vehicle, err := repositories.Repo.GetVehicleByDomain(domain)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar usuario", err)
	}
	return vehicle, nil
}

func VehicleGetByClientID(clientID string) (*[]models.Vehicle, error) {
	vehicles, err := repositories.Repo.GetVehicleByClientID(clientID)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los vehiculos", err)
	}
	return &vehicles, nil
}

func VehicleUpdate(id string, vehicleUpdate *models.VehicleUpdate) (string, error) {
	err := repositories.Repo.UpdateVehicle(&models.Vehicle{
		ID: id,
		Domain:   vehicleUpdate.Domain,		
		Brand:    vehicleUpdate.Brand,
		Model:    vehicleUpdate.Model,
		Color:    vehicleUpdate.Color,
		Year:     vehicleUpdate.Year,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return "", models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return id, nil
}

func VehicleDelete(id string) (error) {
	err := repositories.Repo.DeleteVehicle(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar cliente", err)
	}
	return nil
}
