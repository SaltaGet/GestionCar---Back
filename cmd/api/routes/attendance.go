package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func AttendanceRoutes(app *fiber.App) {
	att := app.Group("/attendance", middleware.AuthMiddleware(), middleware.TenantMiddleware())

	att.Get("/get_all", GetController("AttendanceController", func(c *fiber.Ctx, ctrl *controllers.AttendanceController) error {
	return ctrl.GetAllAttendances(c)
}))

	att.Post("/get_by_date", GetController("AttendanceController", func(c *fiber.Ctx, ctrl *controllers.AttendanceController) error {
		return ctrl.GetAllAttendancesByDate(c)
	}))

	att.Post("/create", GetController("AttendanceController", func(c *fiber.Ctx, ctrl *controllers.AttendanceController) error {
		return ctrl.CreateAttendance(c)
	}))

	att.Put("/update", GetController("AttendanceController", func(c *fiber.Ctx, ctrl *controllers.AttendanceController) error {
		return ctrl.UpdateAttendance(c)
	}))

	att.Get("/get_by_employee/:employee_id", GetController("AttendanceController", func(c *fiber.Ctx, ctrl *controllers.AttendanceController) error {
		return ctrl.GetAttendanceByEmployeeID(c)
	}))

	att.Delete("/delete/:id", GetController("AttendanceController", func(c *fiber.Ctx, ctrl *controllers.AttendanceController) error {
		return ctrl.DeleteAttendance(c)
	}))

	att.Get("/:id", GetController("AttendanceController", func(c *fiber.Ctx, ctrl *controllers.AttendanceController) error {
		return ctrl.GetAttendanceByID(c)
	}))
}

