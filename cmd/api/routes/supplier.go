package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SupplierRoutes(app *fiber.App, controllers *controllers.SupplierController){
	att := app.Group("/supplier", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	att.Get("/get_all", controllers.SupplierGetAll)
	att.Get("/get_by_name", controllers.SupplierGetByName)
	att.Post("/create", controllers.SupplierCreate)
	att.Put("/update", controllers.SupplierUpdate)
	att.Delete("/delete/:id", controllers.SupplierDeleteByID)
	att.Get("/:id", controllers.SupplierGetByID)
}