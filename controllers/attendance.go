package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// GetAttendanceByID godoc
//	@Summary		Get Attendance By ID
//	@Description	Get Attendance by ID
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string	true	"Workplace Token"
//	@Param			id					path		string	true	"ID of Attendance"
//	@Success		200					{object}	models.Response{body=models.AttendanceLaundry}
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		404					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/{id} [get]
func GetAttendanceByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.GetAttendanceByID(id, workplace.Identifier)
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
			Message: "Asistencia obtenida con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Asistencia obtenida con éxito",
	})
}

// GetAllAttendances godoc
//	@Summary		Get all attendances
//	@Description	Get all attendances by workplace required auth token
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string	true	"Workplace Token"
//	@Success		200					{object}	models.Response
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		422					{object}	models.Response
//	@Failure		404					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/get_all [get]
func GetAllAttendances(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundries, workshops, err := services.GetAllAttendances(workplace.Identifier)
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

	if laundries != nil {
		return c.Status(200).JSON(models.Response{
			Status:  true,
			Body:    laundries,
			Message: "Asistencia obtenida con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshops,
		Message: "Asistencia obtenida con éxito",
	})
}

// GetAllAttendancesByDate godoc
//	@Summary		Get all attendances within a date range
//	@Description	Retrieve all attendances within a specified date range for a given workplace
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string				true	"Workplace Token"
//	@Param			dateFrom			body		models.DateBetween	true	"Date Between"
//	@Success		200					{object}	models.Response{body=[]models.AttendanceLaundry}
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		403					{object}	models.Response
//	@Failure		422					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/get_by_date [post]
func GetAllAttendancesByDate(c *fiber.Ctx) error {
	var dateBeetwen models.DateBetween
	if err := c.BodyParser(&dateBeetwen); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := dateBeetwen.Validate(); err != nil {
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

	laundries, workshops, err := services.GetAllAttendancesByDate(dateBeetwen.DateFrom, dateBeetwen.DateTo,workplace.Identifier)
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

	if laundries != nil {
		return c.Status(200).JSON(models.Response{
			Status:  true,
			Body:    laundries,
			Message: "Asistencias obtenidas con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshops,
		Message: "Asistencias obtenidas con éxito",
	})
}

// GetAttendanceByEmployeeID godoc
//	@Summary		Get Attendance By Employee ID
//	@Description	Get Attendance by Employee ID
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string	true	"Workplace Token"
//	@Param			employee_id			path		string	true	"ID of Employee"
//	@Success		200					{object}	models.Response
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		403					{object}	models.Response
//	@Failure		404					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/get_by_employee/{employee_id} [get]
func GetAttendanceByEmployeeID(c *fiber.Ctx) error {
	employee_id := c.Params("employee_id")
	if employee_id == "" {
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

	laundry, workshop, err := services.GetAttendanceByEmployeeID(employee_id, workplace.Identifier)
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
			Message: "Asistencias obtenidas con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Asistencias obtenidas con éxito",
	})
}

// CreateAttendance godoc
//	@Summary		Create Attendance
//	@Description	Create Attendance by given workplace
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string					true	"Workplace Token"
//	@Param			attendanceCreate	body		models.AttendanceCreate	true	"Employee body"
//	@Success		200					{object}	models.Response
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		403					{object}	models.Response
//	@Failure		404					{object}	models.Response
//	@Failure		422					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/create [post]
func CreateAttendance(c *fiber.Ctx) error {
	var attendanceCreate models.AttendanceCreate
	if err := c.BodyParser(&attendanceCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := attendanceCreate.Validate(); err != nil {
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

	id, err := services.CreateAttendance(&attendanceCreate, workplace.Identifier)
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
		Message: "Asistencia creada con éxito",
	})
}

// UpdateAttendance godoc
//	@Summary		Update Attendance
//	@Description	Update Attendance by ID
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string					true	"Workplace Token"
//	@Param			attendanceUpdate	body		models.AttendanceUpdate	true	"Employee body"
//	@Success		200					{object}	models.Response
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		403					{object}	models.Response
//	@Failure		404					{object}	models.Response
//	@Failure		422					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/update [put]
func UpdateAttendance(c *fiber.Ctx) error {
	var attendanceUpdate models.AttendanceUpdate
	if err := c.BodyParser(&attendanceUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := attendanceUpdate.Validate(); err != nil {
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

	err := services.UpdateAttendance(&attendanceUpdate, workplace.Identifier)
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
		Message: "Asistencia editada con éxito",
	})
}

// DeleteAttendance godoc
//	@Summary		Delete Attendance
//	@Description	Delete Attendance by ID
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string	true	"Workplace Token"
//	@Param			id					path		string	true	"ID of Attendance"
//	@Success		200					{object}	models.Response
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		403					{object}	models.Response
//	@Failure		404					{object}	models.Response
//	@Failure		422					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/delete/{id} [delete]
func DeleteAttendance(c *fiber.Ctx) error {
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

	err := services.DeleteAttendance(id, workplace.Identifier)
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
		Message: "Asistencia eliminada con éxito",
	})
}
