package routes

import (
	// "github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	// "github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ServiceRoutes(app *fiber.App){
	// att := app.Group("/service", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	// att.Get("/get_all", controllers.ServiceGetAll)
	// att.Post("/create", controllers.ServiceCreate)
	// att.Put("/update", controllers.ServiceUpdate)
	// att.Delete("/delete/:id", controllers.ServiceDeleteByID)
	// att.Get("/:id", controllers.ServiceGetByID)
}
// func ServiceRoutes(app *fiber.App, controllers *controllers.ServiceController){
// 	att := app.Group("/service", middleware.AuthMiddleware(), middleware.TenantMiddleware())
// 	att.Get("/get_all", controllers.ServiceGetAll)
// 	att.Post("/create", controllers.ServiceCreate)
// 	att.Put("/update", controllers.ServiceUpdate)
// 	att.Delete("/delete/:id", controllers.ServiceDeleteByID)
// 	att.Get("/:id", controllers.ServiceGetByID)
// }