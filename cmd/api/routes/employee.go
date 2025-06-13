package routes

import (
	// "github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	// "github.com/DanielChachagua/GestionCar/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(app *fiber.App){
	// att := app.Group("/employee", middleware.AuthMiddleware(), middleware.TenantMiddleware())
	// att.Get("/get_all", controllers.GetAllEmployees)
	// att.Get("/get_by_name", controllers.GetEmployeeByName)
	// att.Post("/create", controllers.CreateEmployee)
	// att.Put("/update", controllers.UpdateEmployee)
	// att.Delete("/delete/:id", controllers.DeleteEmployee)
	// att.Get("/:id", controllers.GetEmployeeByID)
}
// func EmployeeRoutes(app *fiber.App, controllers *controllers.EmployeeController){
// 	att := app.Group("/employee", middleware.AuthMiddleware(), middleware.TenantMiddleware())
// 	att.Get("/get_all", controllers.GetAllEmployees)
// 	att.Get("/get_by_name", controllers.GetEmployeeByName)
// 	att.Post("/create", controllers.CreateEmployee)
// 	att.Put("/update", controllers.UpdateEmployee)
// 	att.Delete("/delete/:id", controllers.DeleteEmployee)
// 	att.Get("/:id", controllers.GetEmployeeByID)
// }