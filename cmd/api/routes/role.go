package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App, controllers *controllers.RoleController){
	role := app.Group("/role", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	role.Get("/get_all", controllers.RoleGetAll)
	role.Post("/create", controllers.RoleCreate)
}