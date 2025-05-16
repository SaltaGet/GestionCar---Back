package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func PurchaseOrderRoutes(app *fiber.App){
	att := app.Group("/purchase_order", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_all", controllers.PurchaseOrderGetAll)
	att.Post("/create", controllers.PurchaseOrderCreate)
	att.Put("/update", controllers.PurchaseOrderUpdate)
	att.Delete("/delete/:id", controllers.PurchaseOrderDelete)
	att.Get("/:id", controllers.PurchaseOrderGetByID)
}