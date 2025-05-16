package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App){
	att := app.Group("/role", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_all", controllers.GetRolesWorkplace)
}