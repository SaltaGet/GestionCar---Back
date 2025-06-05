package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func MovementRoutes(app *fiber.App, controllers *controllers.MovementTypeController){
	att := app.Group("/movement", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	att.Get("/get_all", controllers.GetAllMovementTypes)
	att.Post("/create", controllers.MovementTypeCreate)
	att.Put("/update", controllers.MovementTypeUpdate)
	att.Delete("/delete/:id", controllers.MovementTypeDelete)
	att.Get("/:id", controllers.GetMovementTypeByID)
}