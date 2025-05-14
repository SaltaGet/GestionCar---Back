package controllers

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
// GetEmployeeByID godoc
// @Summary     Get Employee By ID
// @Description Get Employee By ID
// @Tags        employee
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of Employee"
// @Success     200  {object}  models.Response
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /employee/{id} [get]
// @Security    BearerAuth
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

// VehicleGetAll retrieves all vehicles from the repository.
// @Summary Get all vehicles
// @Description Fetches all vehicles stored in the system.
// @Tags vehicle
// @Produce json
// @Success 200 {object} []models.Vehicle "List of vehicles retrieved successfully"
// @Failure 500 {object} models.Response "Internal server error"
// @Router /vehicles [get]
func VehicleGetAll() (*[]models.Vehicle, error) {
	vehicles, err := repositories.Repo.GetAllVehicles()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los vehiculos", err)
	}
	return &vehicles, nil
}

// VehicleGetByID godoc
// @Summary     Get Vehicle By ID
// @Description Get Vehicle By ID
// @Tags        vehicle
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of Vehicle"
// @Success     200  {object}  models.Vehicle
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /vehicle/{id} [get]
// @Security    BearerAuth
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

// VehicleGetByDomain retrieves all vehicles that contain the given domain.
// @Summary Get Vehicles By Domain
// @Description Fetches all vehicles that contain the given domain.
// @Tags vehicle
// @Produce json
// @Param domain path string true "Domain string"
// @Success 200 {object} []models.Vehicle "List of vehicles retrieved successfully"
// @Failure 400 {object} models.Response "Bad request"
// @Failure 404 {object} models.Response "Not found"
// @Failure 500 {object} models.Response "Internal server error"
// @Router /vehicle/domain/{domain} [get]
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

// VehicleGetByClientID retrieves all vehicles that belong to the given client.
// @Summary Get Vehicles By Client ID
// @Description Fetches all vehicles that belong to the given client.
// @Tags vehicle
// @Produce json
// @Param clientID path string true "Client ID"
// @Success 200 {object} []models.Vehicle "List of vehicles retrieved successfully"
// @Failure 400 {object} models.Response "Bad request"
// @Failure 404 {object} models.Response "Not found"
// @Failure 500 {object} models.Response "Internal server error"
// @Router /vehicle/client/{clientID} [get]
func VehicleGetByClientID(clientID string) (*[]models.Vehicle, error) {
	vehicles, err := repositories.Repo.GetVehicleByClientID(clientID)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los vehiculos", err)
	}
	return &vehicles, nil
}

// VehicleUpdate godoc
// @Summary     Update Vehicle
// @Description Update Vehicle with the given ID.
// @Tags        vehicle
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of Vehicle"
// @Param       vehicleUpdate body models.VehicleUpdate true "VehicleUpdate"
// @Success     200  {object}  models.Response
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /vehicle/{id} [put]
// @Security    BearerAuth
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

// VehicleDelete godoc
// @Summary     Delete Vehicle
// @Description Delete Vehicle with the given ID.
// @Tags        vehicle
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of Vehicle"
// @Success     200
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /vehicle/{id} [delete]
// @Security    BearerAuth
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
