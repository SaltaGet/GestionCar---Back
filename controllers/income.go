package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// GetIncomeByID godoc
//	@Summary		Get Income By ID
//	@Description	Fetches income details from either laundry or workshop based on the provided ID and workplace context.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string										true	"Workplace Token"
//	@Param			id					path		string										true	"ID of the income"
//	@Success		200					{object}	models.Response{body=models.IncomeLaundry}	"Income details fetched successfully"
//	@Failure		400					{object}	models.Response								"Bad Request"
//	@Failure		401					{object}	models.Response								"Auth is required"
//	@Failure		403					{object}	models.Response								"Not Authorized"
//	@Failure		404					{object}	models.Response								"Expense not found"
//	@Failure		500					{object}	models.Response								"Internal server error"
//	@Router			/income/{id} [get]
func GetIncomeByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.GetIncomeByID(id, workplace.Identifier)
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
			Message: "Ingreso obtenido con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Ingreso obtenido con éxito",
	})
}

// GetAllIncomes godoc
//	@Summary		Get all incomes
//	@Description	Fetches all incomes from the specified workplace, either in laundry or workshop.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string											true	"Workplace Token"
//	@Success		200					{object}	models.Response{body=[]models.IncomeLaundry}	"List of incomes"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		404					{object}	models.Response									"Expense not found"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/income/get_all [get]
func GetAllIncomes(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.GetAllIncomes(workplace.Identifier)
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
			Message: "Ingresos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Ingresos obtenidos con éxito",
	})
}

// GetIncomeToday godoc
//	@Summary		Get Income Today
//	@Description	Fetches all incomes from the specified workplace, either in laundry or workshop, on the current day.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string											true	"Workplace Token"
//	@Success		200					{object}	models.Response{body=[]models.IncomeLaundry}	"List of laundry incomes"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		404					{object}	models.Response									"Expense not found"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/income/get_today [get]
func GetIncomeToday(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.GetIncomeToday(workplace.Identifier)
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
			Message: "Ingresos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Ingresos obtenidos con éxito",
	})
}

// CreateIncome godoc
//	@Summary		Create Income
//	@Description	Parses the request body to create a new income entry for either laundry or workshop.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string							true	"Workplace Token"
//	@Param			incomeCreate		body		models.IncomeCreate				true	"Income information"
//	@Success		200					{object}	models.Response{body=string}	"Income created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		404					{object}	models.Response					"Expense not found"
//	@Failure		422					{object}	models.Response					"Model Invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/income/create [post]
func CreateIncome(c *fiber.Ctx) error {
	var incomeCreate models.IncomeCreate
	if err := c.BodyParser(&incomeCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := incomeCreate.Validate(); err != nil {
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

	id, err := services.CreateIncome(&incomeCreate, workplace.Identifier)
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
		Message: "Ingreso creado con éxito",
	})
}

// UpdateIncome godoc
//	@Summary		Update Income
//	@Description	Updates the details of an income based on the provided data.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string				true	"Workplace Token"
//	@Param			incomeUpdate		body		models.IncomeUpdate	true	"Income data to update"
//	@Success		200					{object}	models.Response		"Income updated successfully"
//	@Failure		400					{object}	models.Response		"Bad Request"
//	@Failure		401					{object}	models.Response		"Auth is required"
//	@Failure		403					{object}	models.Response		"Not Authorized"
//	@Failure		404					{object}	models.Response		"Expense not found"
//	@Failure		422					{object}	models.Response		"Model Invalid"
//	@Failure		500					{object}	models.Response		"Internal server error"
//	@Router			/income/update [put]
func UpdateIncome(c *fiber.Ctx) error {
	var incomeUpdate models.IncomeUpdate
	if err := c.BodyParser(&incomeUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := incomeUpdate.Validate(); err != nil {
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

	err := services.UpdateIncome(&incomeUpdate, workplace.Identifier)
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
		Message: "Ingreso editado con éxito",
	})
}

// DeleteIncome godoc
//	@Summary		Delete Income
//	@Description	Deletes an income entry based on the provided ID and workplace context.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string			true	"Workplace Token"
//	@Param			id					path		string			true	"ID of the income"
//	@Success		200					{object}	models.Response	"Income deleted successfully"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Expense not found"
//	@Failure		500					{object}	models.Response	"Error interno"
//	@Router			/income/delete/{id} [delete]
func DeleteIncome(c *fiber.Ctx) error {
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

	err := services.DeleteIncome(id, workplace.Identifier)
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
		Message: "Ingreso eliminado con éxito",
	})
}
