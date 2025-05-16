package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)


// GetEmployeeByID godoc
//	@Summary		Get Employee By ID
//	@Description	Get Employee By ID
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string											true	"Workplace Token"
//	@Param			id					path		string											true	"ID of Employee"
//	@Success		200					{object}	models.Response{body=models.EmployeeLaundry}	"Employee obtained successfully"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		404					{object}	models.Response									"Employee not found"
//	@Failure		500					{object}	models.Response
//	@Router			/employee/{id} [get]
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

// GetAllEmployees godoc
//	@Summary		Get all employees
//	@Description	Fetches all employees from the specified workplace, either in laundry or workshop.
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string											true	"Workplace Token"
//	@Success		200					{object}	models.Response{body=[]models.EmployeeLaundry}	"List of employees"
//	@Failure		400					{object}	models.Response									"Bad request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/employee/get_all [get]
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

// GetEmployeeByName godoc
//	@Summary		Get Employee By Name
//	@Description	Fetches employees from either laundry or workshop based on the provided name and workplace.
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string											true	"Workplace Token"
//	@Param			name				query		string											true	"Name of the Employee"
//	@Success		200					{object}	models.Response{body=[]models.EmployeeLaundry}	"List of laundry employees"
//	@Failure		400					{object}	models.Response									"Bad request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/employee/get_by_name [get]
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

// CreateEmployee godoc
//	@Summary		Create Employee
//	@Description	Creates an employee for either laundry or workshop based on the provided information.
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string							true	"Workplace Token"
//	@Param			employeeCreate		body		models.EmployeeCreate			true	"Employee information"
//	@Success		200					{object}	models.Response{body=string}	"Employee created"
//	@Failure		400					{object}	models.Response					"Bad request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		422					{object}	models.Response					"Model Invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/employee/create [post]
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

// UpdateEmployee godoc
//	@Summary		Update Employee
//	@Description	Updates the details of an employee based on the provided data.
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string					true	"Workplace Token"
//	@Param			employeeUpdate		body		models.EmployeeUpdate	true	"Employee data to update"
//	@Success		200					{object}	models.Response			"Empleado editado con éxito"
//	@Failure		400					{object}	models.Response			"Invalid request or Workplace is required"
//	@Failure		401					{object}	models.Response			"Auth is required"
//	@Failure		403					{object}	models.Response			"Not Authorized"
//	@Failure		404					{object}	models.Response			"Not Found"
//	@Failure		422					{object}	models.Response			"Model Invalid"
//	@Failure		500					{object}	models.Response			"Error interno"
//	@Router			/employee/update [put]
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

// DeleteEmployee godoc
//	@Summary		Delete Employee
//	@Description	Removes an employee from the database based on the provided ID and workplace context.
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string			true	"Workplace Token"
//	@Param			id					path		string			true	"ID of the employee"
//	@Success		200					{object}	models.Response	"Empleado eliminado con éxito"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Not Found"
//	@Failure		500					{object}	models.Response	"Error interno"
//	@Router			/employee/delete/{id} [delete]
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

