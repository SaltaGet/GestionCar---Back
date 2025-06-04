package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)


// GetExpenseByID godoc
//	@Summary		Get Expense By ID
//	@Description	Get Expense By ID
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string										true	"ID of Expense"
//	@Success		200					{object}	models.Response{body=models.Expense}	"Expense obtained successfully"
//	@Failure		400					{object}	models.Response								"Bad Request"
//	@Failure		401					{object}	models.Response								"Auth is required"
//	@Failure		403					{object}	models.Response								"Not Authorized"
//	@Failure		404					{object}	models.Response								"Expense not found"
//	@Failure		500					{object}	models.Response
//	@Router			/expense/{id} [get]
func (e *ExpenseController) GetExpenseByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	expense, err := e.ExpenseService.ExpenseGetByID(id)
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
		Body:    expense,
		Message: "Egreso obtenido con éxito",
	})
}

// GetAllExpenses godoc
//	@Summary		Get all expenses
//	@Description	Fetches all expenses from the specified tenant.
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200					{object}	models.Response{body=[]models.Expense}	"List of expenses"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/expense/get_all [get]
func (e *ExpenseController) GetAllExpenses(c *fiber.Ctx) error {
	expenses, err := e.ExpenseService.ExpenseGetAll()
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
		Body:    expenses,
		Message: "Egresos obtenidos con éxito",
	})
}

// GetExpenseToday godoc
//	@Summary		Get expense today
//	@Description	Fetches all expenses from the specified workplace, either in laundry or workshop, on the current day.
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200					{object}	models.Response{body=[]models.ExpenseLaundry}	"List of laundry expenses"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/expense/get_today [get]
func (e *ExpenseController) GetExpenseToday(c *fiber.Ctx) error {
	expenses, err := e.ExpenseService.ExpenseGetToday()
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
		Body:    expenses,
		Message: "Egresos obtenidos con éxito",
	})
}

// CreateExpense godoc
//	@Summary		Create Expense
//	@Description	Parses the request body to create a new expense entry.
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			expenseCreate		body		models.ExpenseCreate			true	"Expense information"
//	@Success		200					{object}	models.Response{body=string}	"Expense created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		422					{object}	models.Response					"Model Invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/expense/create [post]
func (e *ExpenseController) CreateExpense(c *fiber.Ctx) error {
	var expenseCreate models.ExpenseCreate
	if err := c.BodyParser(&expenseCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := expenseCreate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := e.ExpenseService.ExpenseCreate(&expenseCreate)
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
		Message: "Egreso creado con éxito",
	})
}

// UpdateExpense godoc
//	@Summary		Update Expense
//	@Description	Updates the details of an expense based on the provided data.
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			expenseUpdate		body		models.ExpenseUpdate	true	"Expense data to update"
//	@Success		200					{object}	models.Response			"Expense updated successfully"
//	@Failure		400					{object}	models.Response			"Bad Request"
//	@Failure		401					{object}	models.Response			"Auth is required"
//	@Failure		403					{object}	models.Response			"Not Authorized"
//	@Failure		422					{object}	models.Response			"Model Invalid"
//	@Failure		500					{object}	models.Response			"Internal server error"
//	@Router			/expense/update [put]
func (e *ExpenseController) UpdateExpense(c *fiber.Ctx) error {
	var expenseUpdate models.ExpenseUpdate
	if err := c.BodyParser(&expenseUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := expenseUpdate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := e.ExpenseService.ExpenseUpdate(&expenseUpdate)
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
		Message: "Egreso editado con éxito",
	})
}

// DeleteExpense deletes an expense by its ID from the specified workplace.
//	@Summary		Delete Expense
//	@Description	Deletes an expense based on the provided ID and workplace context.
//	@Tags			Expense
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string			true	"ID of the expense"
//	@Success		200					{object}	models.Response	"Expense deleted successfully"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		500					{object}	models.Response	"Internal server error"
//	@Router			/expense/delete/{id} [delete]
func (e *ExpenseController) DeleteExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := e.ExpenseService.ExpenseDelete(id)
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
		Message: "Egreso eliminado con éxito",
	})
}

