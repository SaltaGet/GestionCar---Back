package routes

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ResumeRoutes(app *fiber.App, controllers *controllers.ResumeController) {
	resume := app.Group("/resume", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	resume.Put("/update/expense", controllers.ExpenseResumeUpdate)
	resume.Put("/update/income", controllers.IncomeResumeUpdate)
	resume.Get("/get_by_date/expense", controllers.ExpenseResumeGetByDateBetween)
	resume.Get("/get_by_date/income", controllers.IncomeResumeGetByDateBetween)
	resume.Get("/get/expense/:id", controllers.ExpenseResumeGetByID)
	resume.Get("/get/income/:id", controllers.IncomeResumeGetByID)
}