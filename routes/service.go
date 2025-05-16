package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func ServiceRoutes(app *fiber.App){
	att := app.Group("/service", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_all", controllers.ServiceGetAll)
	att.Post("/create", controllers.ServiceCreate)
	att.Put("/update", controllers.ServiceUpdate)
	att.Delete("/delete/:id", controllers.ServiceDeleteByID)
	att.Get("/:id", controllers.ServiceGetByID)
}