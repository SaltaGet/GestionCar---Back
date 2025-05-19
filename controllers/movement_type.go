package controllers

import (
	"strconv"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// GetMovementTypeByID godoc
//	@Summary		Get Movement Type By ID
//	@Description	Get Movement Type By ID
//	@Tags			Movement
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string												true	"Workplace Token"
//	@Param			id					path		string												true	"ID of the movement type"
//	@Success		200					{object}	models.Response{body=models.MovementTypeLaundry}	"Movement type details"
//	@Failure		400					{object}	models.Response										"Bad Request"
//	@Failure		401					{object}	models.Response										"Auth is required"
//	@Failure		403					{object}	models.Response										"Not Authorized"
//	@Failure		404					{object}	models.Response										"Expense not found"
//	@Failure		500					{object}	models.Response										"Internal server error"
//	@Router			/movement/{id} [get]
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
//	@Summary		Get all movement types
//	@Description	Get all movement types from either laundry or workshop based on the provided isIncome query parameter.
//	@Tags			Movement
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string												true	"Workplace Token"
//	@Param			isIncome			query		bool												true	"Is income movement type"
//	@Success		200					{object}	models.Response{body=[]models.MovementTypeLaundry}	"List of movement types"
//	@Failure		400					{object}	models.Response										"Bad Request"
//	@Failure		401					{object}	models.Response										"Auth is required"
//	@Failure		403					{object}	models.Response										"Not Authorized"
//	@Failure		404					{object}	models.Response										"Expense not found"
//	@Failure		500					{object}	models.Response										"Internal server error"
//	@Router			/movement/get_all [get]
func GetAllMovementTypes(c *fiber.Ctx) error {
	isIncomeStr := c.Query("isIncome")
isIncome := false
if isIncomeStr != "" {
    var err error
    isIncome, err = strconv.ParseBool(isIncomeStr)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(models.Response{
            Status:  false,
            Body:    nil,
            Message: "Invalid value for isIncome",
        })
    }
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

// MovementTypeCreate godoc
//	@Summary		Create Movement Type
//	@Description	This endpoint creates a new movement type based on the provided JSON payload.
//	@Tags			Movement
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string							true	"Workplace Token"
//	@Param			movementType		body		models.MovementTypeCreate		true	"Movement Type Details"
//	@Success		200					{object}	models.Response{body=string}	"Movement created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		404					{object}	models.Response					"Expense not found"
//	@Failure		422					{object}	models.Response					"Model invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/movement/create [post]
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

// MovementTypeUpdate godoc
//	@Summary		Update Movement Type
//	@Description	This endpoint updates a movement type based on the provided JSON payload.
//	@Tags			Movement
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string						true	"Workplace Token"
//	@Param			movementType		body		models.MovementTypeUpdate	true	"Movement Type Details"
//	@Success		200					{object}	models.Response				"Movement updated successfully"
//	@Failure		400					{object}	models.Response				"Bad Request"
//	@Failure		401					{object}	models.Response				"Auth is required"
//	@Failure		403					{object}	models.Response				"Not Authorized"
//	@Failure		404					{object}	models.Response				"Expense not found"
//	@Failure		422					{object}	models.Response				"Model invalid"
//	@Failure		500					{object}	models.Response				"Internal server error"
//	@Router			/movement/update [put]
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

// MovementTypeDelete godoc
//	@Summary		Delete Movement Type
//	@Description	Deletes a movement type based on its ID.
//	@Tags			Movement
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string			true	"Workplace Token"
//	@Param			id					path		string			true	"ID of the movement type"
//	@Success		200					{object}	models.Response	"Movement type deleted successfully"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Expense not found"
//	@Failure		500					{object}	models.Response	"Internal server error"
//	@Router			/movement/delete/{id} [delete]
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
