package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// ServiceGetByID get a income by id
//
// @Summary Get a income by id
// @Description Get a income by id
// @Tags services
// @Accept json
// @Produce json
// @Param id path string true "ID of the income to get"
// @Success 200 {object} models.Response{body=models.Income}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /services/{id} [get]
func ServiceGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.ServiceGetByID(id, workplace.Identifier)
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

	if laundry != nil {
		return c.Status(200).JSON(models.Response{
			Status:  true,
			Body:    laundry,
			Message: "Servicio obtenido con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Servicio obtenido con éxito",
	})
}

// ServiceGetAll get all services from workplace
//
// @Summary Get all services from workplace
// @Description Get all services from workplace
// @Tags services
// @Accept json
// @Produce json
// @Success 200 {object} models.Response{body=[]models.ServiceLaundry}
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /services [get]
func ServiceGetAll(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.ServiceGetAll(workplace.Identifier)
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

	if laundry != nil {
		return c.Status(200).JSON(models.Response{
			Status:  true,
			Body:    laundry,
			Message: "Servicios obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Servicios obtenidos con éxito",
	})
}

// ServiceCreate godoc
// @Summary     Create Service
// @Description Creates a service and returns its ID.
// @Tags        services
// @Accept      json
// @Produce     json
// @Param       serviceCreate body     models.ServiceCreate true  "Service creation data"
// @Success     200             {object} models.Response{body=string} "Service created successfully"
// @Failure     400             {object} models.Response       "Invalid request"
// @Failure     400             {object} models.Response       "Workplace is required"
// @Failure     500             {object} models.Response       "Internal server error"
// @Router      /services      [post]
// @Security    BearerAuth
func ServiceCreate(c *fiber.Ctx) error {
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

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	id, err := services.ServiceCreate(&serviceCreate, workplace.Identifier)
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

// ServiceUpdate updates the information of a service.
//
// @Summary     Update Service
// @Description Updates the details of a service based on the provided data.
// @Tags        services
// @Accept      json
// @Produce     json
// @Param       serviceUpdate  body      models.ServiceUpdate  true  "Service data to update"
// @Success     200  {object}  models.Response  "Servicio editado con éxito"
// @Failure     400  {object}  models.Response  "Invalid request or Workplace is required"
// @Failure     500  {object}  models.Response  "Error interno"
// @Router      /services [put]
// @Security    BearerAuth
func ServiceUpdate(c *fiber.Ctx) error {
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

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	err := services.ServiceUpdate(&serviceUpdate, workplace.Identifier)
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

// ServiceDeleteByID deletes a service by its ID for a specific workplace.
// @Summary     Delete Service
// @Description Deletes a service based on the provided ID and workplace context.
// @Tags        services
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the service"
// @Success     200  {object}  models.Response  "Servicio eliminado con éxito"
// @Failure     400  {object}  models.Response  "ID is required or Workplace is required"
// @Failure     500  {object}  models.Response  "Error interno"
// @Router      /services/{id} [delete]
// @Security    BearerAuth
func ServiceDeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	err := services.ServiceDeleteByID(id, workplace.Identifier)
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
