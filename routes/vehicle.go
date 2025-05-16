package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func VehicleRoutes(app *fiber.App){
	att := app.Group("/vehicle", middleware.AuthMiddleware())
	att.Get("/get_all", controllers.VehicleGetAll)
	att.Get("/get_by_domain", controllers.VehicleGetByDomain)
	att.Post("/create", controllers.VehicleCreate)
	att.Put("/update", controllers.VehicleUpdate)
	att.Get("/get_by_client/:client_id", controllers.VehicleGetByClientID)
	att.Delete("/delete/:id", controllers.VehicleDelete)
	att.Get("/:id", controllers.VehicleGetByID)
}