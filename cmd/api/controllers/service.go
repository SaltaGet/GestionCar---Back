package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// ServiceGetByID godoc
//	@Summary		Get a service by id
//	@Description	Get a service by id
//	@Tags			Service
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string	true	"ID of the income to get"
//	@Success		200					{object}	models.Response{body=models.Service}
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Service not found"
//	@Failure		500					{object}	models.Response
//	@Router			/service/{id} [get]
func (s *ServiceController) ServiceGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	service, err := s.ServiceService.ServiceGetByID(id)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    service,
		Message: "Servicio obtenido con éxito",
	})
}

// ServiceGetAll godoc
//	@Summary		Get all services from workplace
//	@Description	Get all services from workplace
//	@Tags			Service
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200					{object}	models.Response{body=[]models.Service}
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		500					{object}	models.Response
//	@Router			/service/get_all [get]
func (s *ServiceController) ServiceGetAll(c *fiber.Ctx) error {
	services, err := s.ServiceService.ServiceGetAll()
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    services,
		Message: "Servicios obtenidos con éxito",
	})
}

// ServiceCreate godoc
//	@Summary		Create Service
//	@Description	Creates a service and returns its ID.
//	@Tags			Service
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			serviceCreate		body		models.ServiceCreate			true	"Service creation data"
//	@Success		200					{object}	models.Response{body=string}	"Service created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		422					{object}	models.Response					"Model is invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/service/create      [post]
func (s *ServiceController) ServiceCreate(c *fiber.Ctx) error {
	var serviceCreate models.ServiceCreate
	if err := c.BodyParser(&serviceCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := serviceCreate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := s.ServiceService.ServiceCreate(&serviceCreate)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    id,
		Message: "Servicio creado con éxito",
	})
}

// ServiceUpdate godoc
//	@Summary		Update Service
//	@Description	Updates the details of a service based on the provided data.
//	@Tags			Service
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			serviceUpdate		body		models.ServiceUpdate	true	"Service data to update"
//	@Success		200					{object}	models.Response			"Servicio editado con éxito"
//	@Failure		400					{object}	models.Response			"Bad Request"
//	@Failure		401					{object}	models.Response			"Auth is required"
//	@Failure		403					{object}	models.Response			"Not Authorized"
//	@Failure		404					{object}	models.Response			"Expense not found"
//	@Failure		422					{object}	models.Response			"Model is invalid"
//	@Failure		500					{object}	models.Response			"Error interno"
//	@Router			/service/update [put]
//	@Security		BearerAuth
func (s *ServiceController) ServiceUpdate(c *fiber.Ctx) error {
	var serviceUpdate models.ServiceUpdate
	if err := c.BodyParser(&serviceUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := serviceUpdate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := s.ServiceService.ServiceUpdate(&serviceUpdate)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    nil,
		Message: "Servicio editado con éxito",
	})
}

// ServiceDeleteByID godoc
//	@Summary		Delete Service
//	@Description	Deletes a service based on the provided ID and workplace context.
//	@Tags			Service
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string			true	"ID of the service"
//	@Success		200					{object}	models.Response	"Servicio eliminado con éxito"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Service not found"
//	@Failure		500					{object}	models.Response	"Error interno"
//	@Router			/service/delete/{id} [delete]
func (s *ServiceController) ServiceDeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := s.ServiceService.ServiceDelete(id)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    nil,
		Message: "Servicio eliminado con éxito",
	})
}
