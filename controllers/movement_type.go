package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// GetMovementTypeByID retrieves the movement type by its ID for a specific workplace.
// @Summary     Get Movement Type By ID
// @Description Get Movement Type By ID
// @Tags        movement_type
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the movement type"
// @Success     200  {object}  models.Response{body=models.MovementTypeLaundry} "Movement type details from laundry"
// @Success     200  {object}  models.Response{body=models.MovementTypeWorkshop} "Movement type details from workshop"
// @Failure     400  {object}  models.Response "ID is required or Workplace is required"
// @Failure     404  {object}  models.Response "Movement type not found"
// @Failure     500  {object}  models.Response "Internal server error"
// @Router      /movement_type/{id} [get]
// @Security    BearerAuth
func GetMovementTypeByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.GetMovementTypeByID(id, workplace.Identifier)
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
			Message: "Movimiento obtenido con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Movimiento obtenido con éxito",
	})
}

// GetAllMovementTypes godoc
// @Summary     Get all movement types
// @Description Get all movement types from either laundry or workshop based on the provided isIncome query parameter.
// @Tags        movement_type
// @Accept      json
// @Produce     json
// @Param       isIncome query     bool     true  "Is income movement type"
// @Success     200      {object}  models.Response{body=[]models.MovementTypeLaundry} "List of laundry movement types"
// @Success     200      {object}  models.Response{body=[]models.MovementTypeWorkshop} "List of workshop movement types"
// @Failure     400      {object}  models.Response "Workplace is required or invalid request"
// @Failure     500      {object}  models.Response "Internal server error"
// @Router      /movement_type/get_all [get]
// @Security    BearerAuth
func GetAllMovementTypes(c *fiber.Ctx) error {
	var isIncome bool
	if err := c.QueryParser(&isIncome); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
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

	laundry, workshop, err := services.GetAllMovementTypes(isIncome, workplace.Identifier)
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
			Message: "Movimientos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Movimientos obtenidos con éxito",
	})
}

// MovementTypeCreate handles the creation of a new movement type.
// @Summary     Create Movement Type
// @Description This endpoint creates a new movement type based on the provided JSON payload.
// @Tags        movement_type
// @Accept      json
// @Produce     json
// @Param       movementType body models.MovementTypeCreate true "Movement Type Details"
// @Success     200 {object} models.Response{body=string} "Movement created successfully"
// @Failure     400 {object} models.Response "Invalid request or workplace required"
// @Failure     500 {object} models.Response "Internal server error"
// @Router      /movement_type/create [post]
// @Security    BearerAuth
func MovementTypeCreate(c *fiber.Ctx) error {
	var movementCreate models.MovementTypeCreate
	if err := c.BodyParser(&movementCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := movementCreate.Validate(); err != nil {
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

	id, err := services.MovementTypeCreate(&movementCreate, workplace.Identifier)
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
		Message: "Movimiento creado con éxito",
	})
}

// MovementTypeUpdate handles the update of a movement type.
// @Summary     Update Movement Type
// @Description This endpoint updates a movement type based on the provided JSON payload.
// @Tags        movement_type
// @Accept      json
// @Produce     json
// @Param       movementType body models.MovementTypeUpdate true "Movement Type Details"
// @Success     200 {object} models.Response "Movement updated successfully"
// @Failure     400 {object} models.Response "Invalid request or workplace required"
// @Failure     500 {object} models.Response "Internal server error"
// @Router      /movement_type/update [put]
// @Security    BearerAuth
func MovementTypeUpdate(c *fiber.Ctx) error {
	var movementUpdate models.MovementTypeUpdate
	if err := c.BodyParser(&movementUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := movementUpdate.Validate(); err != nil {
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

	err := services.MovementTypeUpdate(&movementUpdate, workplace.Identifier)
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
		Message: "Movimiento editado con éxito",
	})
}

// MovementTypeDelete deletes a movement type by its ID.
// @Summary     Delete Movement Type
// @Description Deletes a movement type based on its ID.
// @Tags        movement_type
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the movement type"
// @Success     200  {object}  models.Response "Movement type deleted successfully"
// @Failure     400  {object}  models.Response "ID is required or workplace required"
// @Failure     500  {object}  models.Response "Internal server error"
// @Router      /movement_type/{id} [delete]
// @Security    BearerAuth
func MovementTypeDelete(c *fiber.Ctx) error {
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

	err := services.MovementTypeDelete(id, workplace.Identifier)
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
		Message: "Movimiento eliminado con éxito",
	})
}
