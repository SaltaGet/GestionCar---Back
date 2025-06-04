package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)


// GetEmployeeByID godoc
// 	@Summary		Get Employee By ID
// 	@Description	Get Employee By ID
// 	@Tags			Employee
// 	@Accept			json
// 	@Produce		json
// 	@Security		BearerAuth
// 	@Param			id					path		string											true	"ID of Employee"
// 	@Success		200					{object}	models.Response{body=models.Employee}	"Employee obtained successfully"
// 	@Failure		400					{object}	models.Response									"Bad Request"
// 	@Failure		401					{object}	models.Response									"Auth is required"
// 	@Failure		403					{object}	models.Response									"Not Authorized"
// 	@Failure		404					{object}	models.Response									"Employee not found"
// 	@Failure		500					{object}	models.Response
// 	@Router			/employee/{id} [get]
func (e *EmployeeController) GetEmployeeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	employee, err := e.EmployeeService.GetEmployeeByID(id)
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
		Body:    employee,
		Message: "Empleado obtenido con éxito",
	})
}

// GetAllEmployees godoc
// 	@Summary		Get all employees
// 	@Description	Fetches all employees from the specified tenant.
// 	@Tags			Employee
// 	@Accept			json
// 	@Produce		json
// 	@Security		BearerAuth
// 	@Success		200					{object}	models.Response{body=[]models.Employee}	"List of employees"
// 	@Failure		400					{object}	models.Response									"Bad request"
// 	@Failure		401					{object}	models.Response									"Auth is required"
// 	@Failure		403					{object}	models.Response									"Not Authorized"
// 	@Failure		500					{object}	models.Response									"Internal server error"
// 	@Router			/employee/get_all [get]
func (e *EmployeeController) GetAllEmployees(c *fiber.Ctx) error {
	employees, err := e.EmployeeService.GetAllEmployees()
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
		Body:    employees,
		Message: "Empleados obtenidos con éxito",
	})
}

// GetEmployeeByName godoc
// 	@Summary		Get Employee By Name
// 	@Description	Fetches employees from either laundry or workshop based on the provided name and workplace.
// 	@Tags			Employee
// 	@Accept			json
// 	@Produce		json
// 	@Security		BearerAuth
// 	@Param			name				query		string											true	"Name of the Employee"
// 	@Success		200					{object}	models.Response{body=[]models.EmployeeLaundry}	"List of laundry employees"
// 	@Failure		400					{object}	models.Response									"Bad request"
// 	@Failure		401					{object}	models.Response									"Auth is required"
// 	@Failure		403					{object}	models.Response									"Not Authorized"
// 	@Failure		500					{object}	models.Response									"Internal server error"
// 	@Router			/employee/get_by_name [get]
func (e *EmployeeController) GetEmployeeByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" || len(name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El valor no debe de ser vacio o menor a 3 caracteres",
		})
	}

	employees, err := e.EmployeeService.GetEmployeeByName(name)
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
		Body:    employees,
		Message: "Empleados obtenidos con éxito",
	})
}

// CreateEmployee godoc
// 	@Summary		Create Employee
// 	@Description	Creates an employee for either laundry or workshop based on the provided information.
// 	@Tags			Employee
// 	@Accept			json
// 	@Produce		json
// 	@Security		BearerAuth
// 	@Param			employeeCreate		body		models.EmployeeCreate			true	"Employee information"
// 	@Success		200					{object}	models.Response{body=string}	"Employee created"
// 	@Failure		400					{object}	models.Response					"Bad request"
// 	@Failure		401					{object}	models.Response					"Auth is required"
// 	@Failure		403					{object}	models.Response					"Not Authorized"
// 	@Failure		422					{object}	models.Response					"Model Invalid"
// 	@Failure		500					{object}	models.Response					"Internal server error"
// 	@Router			/employee/create [post]
func (e *EmployeeController) CreateEmployee(c *fiber.Ctx) error {
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

	id, err := e.EmployeeService.CreateEmployee(&employeeCreate)
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

// UpdateEmployee godoc
// 	@Summary		Update Employee
// 	@Description	Updates the details of an employee based on the provided data.
// 	@Tags			Employee
// 	@Accept			json
// 	@Produce		json
// 	@Security		BearerAuth
// 	@Param			employeeUpdate		body		models.EmployeeUpdate	true	"Employee data to update"
// 	@Success		200					{object}	models.Response			"Empleado editado con éxito"
// 	@Failure		400					{object}	models.Response			"Invalid request or Workplace is required"
// 	@Failure		401					{object}	models.Response			"Auth is required"
// 	@Failure		403					{object}	models.Response			"Not Authorized"
// 	@Failure		404					{object}	models.Response			"Not Found"
// 	@Failure		422					{object}	models.Response			"Model Invalid"
// 	@Failure		500					{object}	models.Response			"Error interno"
// 	@Router			/employee/update [put]
func (e *EmployeeController) UpdateEmployee(c *fiber.Ctx) error {
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

	err := e.EmployeeService.UpdateEmployee(&employeeUpdate)
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

// DeleteEmployee godoc
// 	@Summary		Delete Employee
// 	@Description	Removes an employee from the database based on the provided ID and tenant context.
// 	@Tags			Employee
// 	@Accept			json
// 	@Produce		json
// 	@Security		BearerAuth
// 	@Param			id					path		string			true	"ID of the employee"
// 	@Success		200					{object}	models.Response	"Empleado eliminado con éxito"
// 	@Failure		400					{object}	models.Response	"Bad Request"
// 	@Failure		401					{object}	models.Response	"Auth is required"
// 	@Failure		403					{object}	models.Response	"Not Authorized"
// 	@Failure		404					{object}	models.Response	"Not Found"
// 	@Failure		500					{object}	models.Response	"Error interno"
// 	@Router			/employee/delete/{id} [delete]
func (e *EmployeeController) DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := e.EmployeeService.DeleteEmployee(id)
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

