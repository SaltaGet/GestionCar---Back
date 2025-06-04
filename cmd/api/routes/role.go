package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App, controllers *controllers.RoleController){
	att := app.Group("/role", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	att.Get("/get_all", controllers.GetRolesWorkplace)
}