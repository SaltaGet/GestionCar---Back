package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
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
func GetEmployeeByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.GetEmployeeByID(id, workplace.Identifier)
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
			Message: "Empleado obtenido con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Empleado obtenido con éxito",
	})
}

// GetAllEmployees retrieves all employees for a specific workplace.
// @Summary		Get all employees
// @Description	Fetches all employees from the specified workplace, either in laundry or workshop.
// @Tags			Employee
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200			{object}	models.Response{body=[]models.EmployeeLaundry} "List of laundry employees"
// @Success		200			{object}	models.Response{body=[]models.EmployeeWorkshop} "List of workshop employees"
// @Failure		400			{object}	models.Response "Workplace is required"
// @Failure		500			{object}	models.Response "Internal server error"
// @Router			/employee/get_all [get]
func GetAllEmployees(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.GetAllEmployees(workplace.Identifier)
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
			Message: "Empleados obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Empleados obtenidos con éxito",
	})
}

// GetEmployeeByName retrieves employees by their name from the specified workplace.
// @Summary     Get Employee By Name
// @Description Fetches employees from either laundry or workshop based on the provided name and workplace.
// @Tags        Employee
// @Accept      json
// @Produce     json
// @Param       name  query     string  true  "Name of the Employee"
// @Success     200   {object}  models.Response{body=[]models.EmployeeLaundry} "List of laundry employees"
// @Success     200   {object}  models.Response{body=[]models.EmployeeWorkshop} "List of workshop employees"
// @Failure     400   {object}  models.Response "Invalid name or workplace"
// @Failure     500   {object}  models.Response "Internal server error"
// @Router      /employee/by-name [get]
// @Security    BearerAuth
func GetEmployeeByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" || len(name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El valor no debe de ser vacio o menor a 3 caracteres",
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

	laundry, workshop, err := services.GetEmployeeByName(name, workplace.Identifier)
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
			Message: "Empleados obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Empleados obtenidos con éxito",
	})
}

// CreateEmployee creates an employee for a specific workplace.
// @Summary		Create Employee
// @Description	Creates an employee for either laundry or workshop based on the provided information.
// @Tags			Employee
// @Accept			json
// @Produce		json
// @Param			employeeCreate	body		models.EmployeeCreate	true	"Employee information"
// @Success		200			{object}	models.Response{body=string} "Employee created"
// @Failure		400			{object}	models.Response "Invalid request or workplace required"
// @Failure		500			{object}	models.Response "Internal server error"
// @Router			/employee [post]
// @Security		BearerAuth
func CreateEmployee(c *fiber.Ctx) error {
	var employeeCreate models.EmployeeCreate
	if err := c.BodyParser(&employeeCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := employeeCreate.Validate(); err != nil {
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

	id, err := services.CreateEmployee(&employeeCreate, workplace.Identifier)
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
		Message: "Empleado creado con éxito",
	})
}

// UpdateEmployee updates an employee's information in the specified workplace.
// @Summary     Update Employee
// @Description Updates the details of an employee based on the provided data.
// @Tags        employee
// @Accept      json
// @Produce     json
// @Param       employeeUpdate  body      models.EmployeeUpdate  true  "Employee data to update"
// @Success     200  {object}  models.Response  "Empleado editado con éxito"
// @Failure     400  {object}  models.Response  "Invalid request or Workplace is required"
// @Failure     500  {object}  models.Response  "Error interno"
// @Router      /employee [put]
// @Security    BearerAuth
func UpdateEmployee(c *fiber.Ctx) error {
	var employeeUpdate models.EmployeeUpdate
	if err := c.BodyParser(&employeeUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := employeeUpdate.Validate(); err != nil {
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

	err := services.UpdateEmployee(&employeeUpdate, workplace.Identifier)
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
		Message: "Empleado editado con éxito",
	})
}

// DeleteEmployee deletes an employee by their ID from the specified workplace.
// @Summary     Delete Employee
// @Description Removes an employee from the database based on the provided ID and workplace context.
// @Tags        Employee
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the employee"
// @Success     200  {object}  models.Response  "Empleado eliminado con éxito"
// @Failure     400  {object}  models.Response  "ID is required or Workplace is required"
// @Failure     500  {object}  models.Response  "Error interno"
// @Router      /employee/{id} [delete]
// @Security    BearerAuth
func DeleteEmployee(c *fiber.Ctx) error {
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

	err := services.DeleteEmployee(id, workplace.Identifier)
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
		Message: "Empleado eliminado con éxito",
	})
}

