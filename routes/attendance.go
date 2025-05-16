package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func AttendanceRoutes(app *fiber.App){
	att := app.Group("/attendance", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_all", controllers.GetAllAttendances)
	att.Post("/get_by_date", controllers.GetAllAttendancesByDate)
	att.Post("/create", controllers.CreateAttendance)
	att.Put("/update", controllers.UpdateAttendance)
	att.Get("/get_by_employee/:employee_id", controllers.GetAttendanceByEmployeeID)
	att.Delete("/delete/:id", controllers.DeleteAttendance)
	att.Get("/:id", controllers.GetAttendanceByID)
}