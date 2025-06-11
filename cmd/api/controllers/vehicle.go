package controllers

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/logging"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// VehicleCreate godoc
//	@Summary		Create Vehicle
//	@Description	Create a new Vehicle
//	@Tags			Vehicle
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			vehicleCreate	body		models.VehicleCreate	true	"Vehicle information"
//	@Success		201				{object}	models.Response
//	@Failure		400				{object}	models.Response	"Bad Request"
//	@Failure		401				{object}	models.Response	"Auth is required"
//	@Failure		403				{object}	models.Response	"Not Authorized"
//	@Failure		422				{object}	models.Response	"Model is invalid"
//	@Failure		500				{object}	models.Response
//	@Router			/vehicle/create [post]
func (v *VehicleController) VehicleCreate(c *fiber.Ctx) error{
	logging.INFO("Crear vehiculo")
	var vehicleCreate models.VehicleCreate
	if err := c.BodyParser(&vehicleCreate); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request" + err.Error(),
		})
	}
	if err := vehicleCreate.Validate(); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	vehicle, err := v.VehicleService.VehicleCreate(&vehicleCreate)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}
	
	logging.INFO("Vehiculo creado con exito")
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    vehicle,
		Message: "Vehiculo creado con exito",
	})
	
}

// VehicleGetAll godoc
//	@Summary		Get all vehicles
//	@Description	Fetches all vehicles stored in the system.
//	@Tags			Vehicle
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	[]models.Vehicle	"List of vehicles retrieved successfully"
//	@Failure		400	{object}	models.Response		"Bad Request"
//	@Failure		401	{object}	models.Response		"Auth is required"
//	@Failure		403	{object}	models.Response		"Not Authorized"
//	@Failure		500	{object}	models.Response		"Internal server error"
//	@Router			/vehicle/get_all [get]
func (v *VehicleController) VehicleGetAll(c *fiber.Ctx) error {
	logging.INFO("Obtener todos los vehiculos")
	vehicles, err := v.VehicleService.VehicleGetAll()
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Vehiculos obtenidos con exito")
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    vehicles,
		Message: "Vehiculos obtenidos con exito",
	})
}

// VehicleGetByID godoc
//	@Summary		Get Vehicle By ID
//	@Description	Get Vehicle By ID
//	@Tags			Vehicle
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string			true	"ID of Vehicle"
//	@Success		200	{object}	models.Vehicle	"Vehicle retrieved successfully"
//	@Failure		400	{object}	models.Response	"Bad Request"
//	@Failure		401	{object}	models.Response	"Auth is required"
//	@Failure		403	{object}	models.Response	"Not Authorized"
//	@Failure		404	{object}	models.Response	"Vehicle not found"
//	@Failure		500	{object}	models.Response
//	@Router			/vehicle/{id} [get]
func (v *VehicleController) VehicleGetByID(c *fiber.Ctx) error {
	logging.INFO("Obtener un vehiculo por ID")
	id := c.Params("id")
	if id == "" {
		logging.ERROR("Error: ID is required")
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	vehicle, err := v.VehicleService.VehicleGetByID(id)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Vehiculo obtenido con exito")
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    vehicle,
		Message: "Vehiculo obtenido con exito",
	})
}

// VehicleGetByDomain godoc
//	@Summary		Get Vehicles By Domain
//	@Description	Fetches all vehicles that contain the given domain.
//	@Tags			Vehicle
//	@Produce		json
//	@Security		BearerAuth
//	@Param			domain	query		string				true	"Domain string"
//	@Success		200		{object}	[]models.Vehicle	"List of vehicles retrieved successfully"
//	@Failure		400		{object}	models.Response		"Bad Request"
//	@Failure		401		{object}	models.Response		"Auth is required"
//	@Failure		403		{object}	models.Response		"Not Authorized"
//	@Failure		500		{object}	models.Response		"Internal server error"
//	@Router			/vehicle/get_by_domain [get]
func (v *VehicleController) VehicleGetByDomain(c *fiber.Ctx) error {
	logging.INFO("Obtener vehiculos por dominio")
	domain := c.Query("domain")
	if domain == "" || len(domain) < 3 {
		logging.ERROR("Error: Dominio requerido o debe de tener al menos 3 caracteres")
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Dominio requerido o debe de tener al menos 3 caracteres",
		})
	}

	vehicles, err := v.VehicleService.VehicleGetByDomain(domain)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Vehiculos obtenidos con exito")
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    vehicles,
		Message: "Vehiculos obtenidos con exito",
	})
}

// VehicleGetByClientID godoc
//	@Summary		Get Vehicles By Client ID
//	@Description	Fetches all vehicles that belong to the given client.
//	@Tags			Vehicle
//	@Produce		json
//	@Security		BearerAuth
//	@Param			client_id	path		string				true	"Client ID"
//	@Success		200			{object}	[]models.Vehicle	"List of vehicles retrieved successfully"
//	@Failure		400			{object}	models.Response		"Bad Request"
//	@Failure		401			{object}	models.Response		"Auth is required"
//	@Failure		403			{object}	models.Response		"Not Authorized"
//	@Failure		500			{object}	models.Response		"Internal server error"
//	@Router			/vehicle/get_by_client/{client_id} [get]
func (v *VehicleController) VehicleGetByClientID(c *fiber.Ctx) error {
	logging.INFO("Obtener vehiculos por ID de cliente")
	clientID := c.Params("client_id")
	if clientID == "" {
		logging.ERROR("Error: Client ID is required")
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Client ID is required",
		})
	}

	vehicles, err := v.VehicleService.VehicleGetByClientID(clientID)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Vehiculos obtenidos con exito")
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    vehicles,
		Message: "Vehiculos obtenidos con exito",
	})
}

// VehicleUpdate godoc
//	@Summary		Update Vehicle
//	@Description	Update Vehicle with the given ID.
//	@Tags			Vehicle
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			vehicleUpdate	body		models.VehicleUpdate	true	"VehicleUpdate"
//	@Success		200				{object}	models.Response			"Vehicle updated successfully"
//	@Failure		400				{object}	models.Response			"Bad Request"
//	@Failure		401				{object}	models.Response			"Auth is required"
//	@Failure		403				{object}	models.Response			"Not Authorized"
//	@Failure		404				{object}	models.Response			"Vehicle not found"
//	@Failure		422				{object}	models.Response			"Model is invalid"
//	@Failure		500				{object}	models.Response
//	@Router			/vehicle/update [put]
func (v *VehicleController) VehicleUpdate(c *fiber.Ctx) error {
	logging.INFO("Actualizar vehiculo")
	var vehicleUpdate models.VehicleUpdate
	if err := c.BodyParser(&vehicleUpdate); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request" + err.Error(),
		})
	}

	if err := vehicleUpdate.Validate(); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := v.VehicleService.VehicleUpdate(&vehicleUpdate)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Vehiculo actualizado con exito")
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    nil,
		Message: "Vehiculo actualizado con exito",
	})
}

// VehicleDelete godoc
//	@Summary		Delete Vehicle
//	@Description	Delete Vehicle with the given ID.
//	@Tags			Vehicle
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string			true	"ID of Vehicle"
//	@Failure		200	{object}	models.Response	"Vehicle deleted successfully"
//	@Failure		400	{object}	models.Response	"Bad Request"
//	@Failure		401	{object}	models.Response	"Auth is required"
//	@Failure		403	{object}	models.Response	"Not Authorized"
//	@Failure		404	{object}	models.Response	"Vehicle not found"
//	@Failure		500	{object}	models.Response
//	@Router			/vehicle/delete/{id} [delete]
func (v *VehicleController) VehicleDelete(c *fiber.Ctx) error {
	logging.INFO("Eliminar vehiculo")
	id := c.Params("id")
	if id == "" {
		logging.ERROR("Error: ID is required")
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := v.VehicleService.VehicleDelete(id)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Vehiculo eliminado con exito")
	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    nil,
		Message: "Vehiculo eliminado con exito",
	})
}
