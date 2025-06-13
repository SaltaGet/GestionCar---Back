package routes

import (
	// "github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	// "github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PermissionRoutes(app *fiber.App){
	// permission := app.Group("/permission", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	// permission.Get("/get_all", controllers.PermissionGetAll)
	// permission.Get("/get_to_me", controllers.PermissionGetToMe)
}
// func PermissionRoutes(app *fiber.App, controllers *controllers.PermissionController){
// 	permission := app.Group("/permission", middleware.AuthMiddleware(), middleware.TenantMiddleware())
// 	permission.Get("/get_all", controllers.PermissionGetAll)
// 	permission.Get("/get_to_me", controllers.PermissionGetToMe)
// }