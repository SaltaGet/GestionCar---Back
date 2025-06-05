package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// GetIncomeByID godoc
//	@Summary		Get Income By ID
//	@Description	Fetches income details from based on the provided ID and tenant context.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string										true	"ID of the income"
//	@Success		200					{object}	models.Response{body=models.Income}	"Income details fetched successfully"
//	@Failure		400					{object}	models.Response								"Bad Request"
//	@Failure		401					{object}	models.Response								"Auth is required"
//	@Failure		403					{object}	models.Response								"Not Authorized"
//	@Failure		404					{object}	models.Response								"Expense not found"
//	@Failure		500					{object}	models.Response								"Internal server error"
//	@Router			/income/{id} [get]
func (i *IncomeController) GetIncomeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	income, err := i.IncomeService.IncomeGetByID(id)
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
		Body:    income,
		Message: "Ingreso obtenido con éxito",
	})
}

// GetAllIncomes godoc
//	@Summary		Get all incomes
//	@Description	Fetches all incomes from the specified tenant.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200					{object}	models.Response{body=[]models.Income}	"List of incomes"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		404					{object}	models.Response									"Expense not found"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/income/get_all [get]
func (i *IncomeController) GetAllIncomes(c *fiber.Ctx) error {
	incomes, err := i.IncomeService.IncomeGetAll()
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
		Body:    incomes,
		Message: "Ingresos obtenidos con éxito",
	})
}

// GetIncomeToday godoc
//	@Summary		Get Income Today
//	@Description	Fetches all incomes from the specified tenant, on the current day.
//	@Tags			Income
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200					{object}	models.Response{body=[]models.Income}	"List of all incomes"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		404					{object}	models.Response									"Expense not found"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/income/get_today [get]
func (i *IncomeController) GetIncomeToday(c *fiber.Ctx) error {
	incomes, err := i.IncomeService.IncomeGetToday()
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
		Body:    incomes,
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
//	@Param			incomeCreate		body		models.IncomeCreate				true	"Income information"
//	@Success		200					{object}	models.Response{body=string}	"Income created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		404					{object}	models.Response					"Expense not found"
//	@Failure		422					{object}	models.Response					"Model Invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/income/create [post]
func (i *IncomeController) CreateIncome(c *fiber.Ctx) error {
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

	id, err := i.IncomeService.IncomeCreate(&incomeCreate)
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
//	@Param			incomeUpdate		body		models.IncomeUpdate	true	"Income data to update"
//	@Success		200					{object}	models.Response		"Income updated successfully"
//	@Failure		400					{object}	models.Response		"Bad Request"
//	@Failure		401					{object}	models.Response		"Auth is required"
//	@Failure		403					{object}	models.Response		"Not Authorized"
//	@Failure		404					{object}	models.Response		"Expense not found"
//	@Failure		422					{object}	models.Response		"Model Invalid"
//	@Failure		500					{object}	models.Response		"Internal server error"
//	@Router			/income/update [put]
func (i *IncomeController) UpdateIncome(c *fiber.Ctx) error {
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

	err := i.IncomeService.IncomeUpdate(&incomeUpdate)
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
//	@Param			id					path		string			true	"ID of the income"
//	@Success		200					{object}	models.Response	"Income deleted successfully"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Expense not found"
//	@Failure		500					{object}	models.Response	"Error interno"
//	@Router			/income/delete/{id} [delete]
func (i *IncomeController) DeleteIncome(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := i.IncomeService.IncomeDelete(id)
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
