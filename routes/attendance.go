package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func AttendanceRoutes(app *fiber.App){
	att := app.Group("/attendance", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/:id", controllers.GetAttendanceByID)
	att.Get("/get_all", controllers.GetAllAttendances)
	att.Post("/create", controllers.CreateAttendance)
	att.Put("/update/:id", controllers.UpdateAttendance)
	att.Delete("/delete/:id", controllers.DeleteAttendance)
}