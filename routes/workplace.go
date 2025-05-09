package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func WorkplaceRoutes(app *fiber.App){
	auth := app.Group("/workplace")
	auth.Get("/get_all", middleware.AuthMiddleware(), controllers.GetWorkplaces)
}