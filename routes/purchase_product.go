package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func PurchaseProductRoutes(app *fiber.App){
	att := app.Group("/purchase_product", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_purchase/:purchase_id", controllers.PurchaseProductGetAllByPurhcaseID)
	att.Post("/create", controllers.PurchaseProductCreate)
	att.Put("/update", controllers.PurchaseProductUpdate)
	att.Delete("/delete/:id", controllers.PurchaseProductDelete)
	att.Get("/:id", controllers.PurchaseProductGetByID)
}