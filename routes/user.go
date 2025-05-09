package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	auth := app.Group("/user")
	auth.Post(
		"/create", 
		middleware.AuthMiddleware(), 
		middleware.WorkplaceMiddleware(), 
		middleware.RoleAuthMiddleware([]string{"super_admin","admin"}), 
		controllers.CreateUser,
	)
}