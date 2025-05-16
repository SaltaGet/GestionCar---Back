package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func MovementRoutes(app *fiber.App){
	att := app.Group("/movement", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_all", controllers.GetAllMovementTypes)
	att.Post("/create", controllers.MovementTypeCreate)
	att.Put("/update", controllers.MovementTypeUpdate)
	att.Delete("/delete/:id", controllers.MovementTypeDelete)
	att.Get("/:id", controllers.GetMovementTypeByID)
}