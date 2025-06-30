package controllers

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/logging"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// GetAttendanceByID godoc
//	@Summary		Get Attendance By ID
//	@Description	Get Attendance by ID
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"ID of Attendance"
//	@Success		200	{object}	models.Response{body=models.AttendanceDTO}
//	@Failure		400	{object}	models.Response
//	@Failure		401	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/attendance/{id} [get]
func (a *AttendanceController) GetAttendanceByID(c *fiber.Ctx) error {
	logging.INFO("Obtener asistencia por ID")
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	attendance, err := a.AttendanceService.AttendanceGetByID(id)
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

	logging.INFO("Asistencia obtenida con éxito")
	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    attendance,
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
//	@Success		200	{object}	models.Response{body=[]models.AttendanceDTO}
//	@Failure		400	{object}	models.Response
//	@Failure		401	{object}	models.Response
//	@Failure		422	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/attendance/get_all [get]
func (a *AttendanceController) GetAllAttendances(c *fiber.Ctx) error {
	logging.INFO("Obteniendo asistencias")
	attendances, err := a.AttendanceService.AttendanceGetAll()
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Asistencias obtenidas con éxito")
	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    attendances,
		Message: "Asistencias obtenida con éxito",
	})
}

// GetAllAttendancesByDate godoc
//	@Summary		Get all attendances within a date range
//	@Description	Retrieve all attendances within a specified date range for a given workplace
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			dateFrom	body		models.DateBetween	true	"Date Between"
//	@Success		200			{object}	models.Response{body=[]models.AttendanceDTO}
//	@Failure		400			{object}	models.Response
//	@Failure		401			{object}	models.Response
//	@Failure		403			{object}	models.Response
//	@Failure		422			{object}	models.Response
//	@Failure		500			{object}	models.Response
//	@Router			/attendance/get_by_date [post]
func (a *AttendanceController) GetAllAttendancesByDate(c *fiber.Ctx) error {
	logging.INFO("Obteniendo asistencias por rango de fechas")
	var dateBeetwen models.DateBetween
	if err := c.BodyParser(&dateBeetwen); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request" + err.Error(),
		})
	}
	if err := dateBeetwen.Validate(); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	attendances, err := a.AttendanceService.AttendanceGetByDate(dateBeetwen.DateFrom, dateBeetwen.DateTo)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Asistencias obtenidas con éxito")
	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    attendances,
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
//	@Param			employee_id	path		string	true	"ID of Employee"
//	@Success		200			{object}	models.Response{body=[]models.AttendanceDTO}
//	@Failure		400			{object}	models.Response
//	@Failure		401			{object}	models.Response
//	@Failure		403			{object}	models.Response
//	@Failure		404			{object}	models.Response
//	@Failure		500			{object}	models.Response
//	@Router			/attendance/get_by_employee/{employee_id} [get]
func (a *AttendanceController) GetAttendanceByEmployeeID(c *fiber.Ctx) error {
	logging.INFO("Obtener asistencias por ID de empleado")
	employee_id := c.Params("employee_id")
	if employee_id == "" {
		logging.ERROR("Error: ID is required")
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	attendances, err := a.AttendanceService.AttendanceGetByEmployeeID(employee_id)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Asistencias obtenidas con éxito")
	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    attendances,
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
//	@Param			attendanceCreate	body		models.AttendanceCreate	true	"Employee body"
//	@Success		200					{object}	models.Response
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		403					{object}	models.Response
//	@Failure		404					{object}	models.Response
//	@Failure		422					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/create [post]
func (a *AttendanceController) CreateAttendance(c *fiber.Ctx) error {
	logging.INFO("Creando asistencia")
	var attendanceCreate models.AttendanceCreate
	if err := c.BodyParser(&attendanceCreate); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request" + err.Error(),
		})
	}
	if err := attendanceCreate.Validate(); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := a.AttendanceService.AttendanceCreate(&attendanceCreate)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Asistencia creada con éxito")
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
//	@Param			attendanceUpdate	body		models.AttendanceUpdate	true	"Employee body"
//	@Success		200					{object}	models.Response
//	@Failure		400					{object}	models.Response
//	@Failure		401					{object}	models.Response
//	@Failure		403					{object}	models.Response
//	@Failure		404					{object}	models.Response
//	@Failure		422					{object}	models.Response
//	@Failure		500					{object}	models.Response
//	@Router			/attendance/update [put]
func (a *AttendanceController) UpdateAttendance(c *fiber.Ctx) error {
	logging.INFO("Actualizando asistencia")
	var attendanceUpdate models.AttendanceUpdate
	if err := c.BodyParser(&attendanceUpdate); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request" + err.Error(),
		})
	}
	if err := attendanceUpdate.Validate(); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := a.AttendanceService.AttendanceUpdate(&attendanceUpdate)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Asistencia editada con éxito")
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
//	@Param			id	path		string	true	"ID of Attendance"
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Failure		401	{object}	models.Response
//	@Failure		403	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Failure		422	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/attendance/delete/{id} [delete]
func (a *AttendanceController) DeleteAttendance(c *fiber.Ctx) error {
	logging.INFO("Eliminando asistencia")
	id := c.Params("id")
	if id == "" {
		logging.ERROR("Error: ID is required")
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := a.AttendanceService.AttendanceDelete(id)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Asistencia eliminada con éxito")
	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    nil,
		Message: "Asistencia eliminada con éxito",
	})
}


// UpdatePayAttendance godoc
//	@Summary		Update Pay Attendance
//	@Description	Update Pay Attendance by listID
//	@Tags			Attendance
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			listIds	body		models.UpdatePay	true	"listID of Attendance"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Response
//	@Failure		401		{object}	models.Response
//	@Failure		403		{object}	models.Response
//	@Failure		404		{object}	models.Response
//	@Failure		422		{object}	models.Response
//	@Failure		500		{object}	models.Response
//	@Router			/attendance/update_pay [put]
func (a *AttendanceController) UpdatePay(c *fiber.Ctx) error {
	logging.INFO("Actualizando estado de pagos de asistencia")
	var updatePay models.UpdatePay
	if err := c.BodyParser(&updatePay); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request" + err.Error(),
		})
	}
	if err := updatePay.Validate(); err != nil {
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := a.AttendanceService.AttendanceUpdatePay(updatePay.ListIDs)
	if err != nil {
		if errResp, ok := err.(*models.ErrorStruc); ok {
			logging.ERROR("Error: %s", errResp.Err.Error())
			return c.Status(errResp.StatusCode).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: errResp.Message,
			})
		}
		logging.ERROR("Error: %s", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Error interno",
		})
	}

	logging.INFO("Pago de las Asistencias actualizada con éxito")
	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    nil,
		Message: "Pagos actualizados con éxito",
	})
}
