package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ClientRoutes(app *fiber.App){
	att := app.Group("/client", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	att.Get("/get_all", controllers.ClientGetAll)
	att.Get("/get_by_name", controllers.ClientGetByName)
	att.Post("/create", controllers.CreateClient)
	att.Put("/update", controllers.ClientUpdate)
	att.Delete("/delete/:id", controllers.ClientDelete)
	att.Get("/:id", controllers.ClientGetByID)
}