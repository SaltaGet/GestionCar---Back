package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App){
	att := app.Group("/product", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_all", controllers.ProductGetAll)
	att.Get("/get_by_name", controllers.ProductGetByName)
	att.Get("/get_by_identifier", controllers.ProductGetByIdentifier)
	att.Post("/create", controllers.ProductCreate)
	att.Put("/update", controllers.ProductUpdate)
	att.Put("/update_stock/:id", controllers.ProductUpdateStock)
	att.Delete("/delete/:id", controllers.ProductDelete)
	att.Get("/:id", controllers.ProductGetByID)
}