package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)


// GetExpenseByID godoc
// @Summary     Get Expense By ID
// @Description Get Expense By ID
// @Tags        expense
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of Expense"
// @Success     200  {object}  models.Response
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /expense/{id} [get]
// @Security    BearerAuth
func GetExpenseByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.GetExpenseByID(id, workplace.Identifier)
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
			Message: "Egreso obtenido con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Egreso obtenido con éxito",
	})
}

// GetAllExpenses retrieves all expenses for a specific workplace.
// @Summary		Get all expenses
// @Description	Fetches all expenses from the specified workplace, either in laundry or workshop.
// @Tags			Expense
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200			{object}	models.Response{body=[]models.ExpenseLaundry} "List of laundry expenses"
// @Success		200			{object}	models.Response{body=[]models.ExpenseWorkshop} "List of workshop expenses"
// @Failure		400			{object}	models.Response "Workplace is required"
// @Failure		500			{object}	models.Response "Internal server error"
// @Router			/expense/get_all [get]
func GetAllExpenses(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.GetAllExpenses(workplace.Identifier)
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
			Message: "Egresos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Egresos obtenidos con éxito",
	})
}

// GetExpenseToday retrieves all expenses for a specific workplace on the current day.
// @Summary		Get expense today
// @Description	Fetches all expenses from the specified workplace, either in laundry or workshop, on the current day.
// @Tags			Expense
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200			{object}	models.Response{body=[]models.ExpenseLaundry} "List of laundry expenses"
// @Success		200			{object}	models.Response{body=[]models.ExpenseWorkshop} "List of workshop expenses"
// @Failure		400			{object}	models.Response "Workplace is required"
// @Failure		500			{object}	models.Response "Internal server error"
// @Router			/expense/get_today [get]
func GetExpenseToday(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.GetExpenseToday(workplace.Identifier)
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
			Message: "Egresos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Egresos obtenidos con éxito",
	})
}

// CreateExpense handles the creation of a new expense for a specific workplace.
// @Summary     Create Expense
// @Description Parses the request body to create a new expense entry for either laundry or workshop.
// @Tags        Expense
// @Accept      json
// @Produce     json
// @Param       expenseCreate body models.ExpenseCreate true "Expense information"
// @Success     200 {object} models.Response{body=string} "Expense created successfully"
// @Failure     400 {object} models.Response "Invalid request or workplace required"
// @Failure     500 {object} models.Response "Internal server error"
// @Router      /expense [post]
// @Security    BearerAuth
func CreateExpense(c *fiber.Ctx) error {
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

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	id, err := services.CreateExpense(&expenseCreate, workplace.Identifier)
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

// UpdateExpense updates the details of an expense in the specified workplace.
// @Summary     Update Expense
// @Description Updates the details of an expense based on the provided data.
// @Tags        Expense
// @Accept      json
// @Produce     json
// @Param       expenseUpdate  body      models.ExpenseUpdate  true  "Expense data to update"
// @Success     200            {object}  models.Response       "Expense updated successfully"
// @Failure     400            {object}  models.Response       "Invalid request or Workplace is required"
// @Failure     500            {object}  models.Response       "Internal server error"
// @Router      /expense [put]
// @Security    BearerAuth
func UpdateExpense(c *fiber.Ctx) error {
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

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	err := services.UpdateExpense(&expenseUpdate, workplace.Identifier)
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
// @Summary     Delete Expense
// @Description Deletes an expense based on the provided ID and workplace context.
// @Tags        Expense
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the expense"
// @Success     200  {object}  models.Response  "Expense deleted successfully"
// @Failure     400  {object}  models.Response  "ID is required or Workplace is required"
// @Failure     500  {object}  models.Response  "Internal server error"
// @Router      /expense/{id} [delete]
// @Security    BearerAuth
func DeleteExpense(c *fiber.Ctx) error {
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

	err := services.DeleteExpense(id, workplace.Identifier)
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

