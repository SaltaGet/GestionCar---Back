package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	// "github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, controllers *controllers.AuthController){
	auth := app.Group("/auth")
	auth.Post("/login", controllers.AuthLogin)
	auth.Get("/tenant_login/:tenant_id", middleware.AuthMiddleware(), controllers.AuthTenant)
}