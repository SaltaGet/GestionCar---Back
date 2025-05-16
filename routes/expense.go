package routes

import (
	"github.com/DanielChachagua/GestionCar/controllers"
	"github.com/DanielChachagua/GestionCar/middleware"
	"github.com/gofiber/fiber/v2"
)

func ExpenseRoutes(app *fiber.App){
	att := app.Group("/expense", middleware.AuthMiddleware(), middleware.WorkplaceMiddleware())
	att.Get("/get_all", controllers.GetAllExpenses)
	att.Get("/get_today", controllers.GetExpenseToday)
	att.Post("/create", controllers.CreateExpense)
	att.Put("/update", controllers.UpdateExpense)
	att.Delete("/delete/:id", controllers.DeleteExpense)
	att.Get("/:id", controllers.GetExpenseByID)
}