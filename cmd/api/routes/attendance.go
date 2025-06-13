package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func AttendanceRoutes(app *fiber.App) {
	att := app.Group("/attendance", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	att.Get("/get_all",
		func(c *fiber.Ctx) error {
			ctrl, ok := c.Locals("AttendanceController").(*controllers.AttendanceController)
			if !ok || ctrl == nil {
				return c.Status(500).SendString("AttendanceController no inicializado correctamente")
			}
			return ctrl.GetAllAttendances(c)
		},
	)
	// att.Post("/get_by_date", controllers.GetAllAttendancesByDate)
	// att.Post("/create", controllers.CreateAttendance)
	// att.Put("/update", controllers.UpdateAttendance)
	// att.Get("/get_by_employee/:employee_id", controllers.GetAttendanceByEmployeeID)
	// att.Delete("/delete/:id", controllers.DeleteAttendance)
	// att.Get("/:id", controllers.GetAttendanceByID)
}

// func AttendanceRoutes(app *fiber.App, controllers *controllers.AttendanceController){
// 	att := app.Group("/attendance", middleware.AuthMiddleware(), middleware.TenantMiddleware())
// 	att.Get("/get_all", controllers.GetAllAttendances)
// 	att.Post("/get_by_date", controllers.GetAllAttendancesByDate)
// 	att.Post("/create", controllers.CreateAttendance)
// 	att.Put("/update", controllers.UpdateAttendance)
// 	att.Get("/get_by_employee/:employee_id", controllers.GetAttendanceByEmployeeID)
// 	att.Delete("/delete/:id", controllers.DeleteAttendance)
// 	att.Get("/:id", controllers.GetAttendanceByID)
// }
