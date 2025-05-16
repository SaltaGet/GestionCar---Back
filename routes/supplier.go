package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func SupplierRoutes(app *fiber.App){
	att := app.Group("/supplier", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_all", controllers.SupplierGetAll)
	att.Get("/get_by_name", controllers.SupplierGetByName)
	att.Post("/create", controllers.SupplierCreate)
	att.Put("/update", controllers.SupplierUpdate)
	att.Delete("/delete/:id", controllers.SupplierDeleteByID)
	att.Get("/:id", controllers.SupplierGetByID)
}