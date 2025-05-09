package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App){
	auth := app.Group("/auth")
	auth.Post("/login", controllers.AuthLogin)
	auth.Get("/workplace_login/:workplace_id", middleware.AuthMiddleware(), controllers.AuthWorkplace)
}