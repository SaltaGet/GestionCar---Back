package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func MemberRoutes(app *fiber.App, controllers *controllers.MemberController){
	auth := app.Group("/member", middleware.AuthMiddleware(), middleware.TenantMiddleware(), middleware.AdminTenantMiddleware())
	auth.Get("/get_all", controllers.MemberGetAll)
}