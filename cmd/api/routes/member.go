package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func MemberRoutes(app *fiber.App, controllers *controllers.MemberController){
	member := app.Group("/member", middleware.AuthMiddleware(), middleware.TenantMiddleware(), middleware.AdminTenantMiddleware())
	member.Get("/get_all", controllers.MemberGetAll)
	member.Post("/create", controllers.MemberCreate)
}