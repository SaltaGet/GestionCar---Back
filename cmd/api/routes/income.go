package routes

// import (
// 	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
// 	"github.com/DanielChachagua/GestionCar/cmd/api/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func IncomeRoutes(app *fiber.App){
// 	att := app.Group("/income", middleware.AuthMiddleware(), middleware.TenantMiddleware())
// 	att.Get("/get_all", controllers.GetAllIncomes)
// 	att.Get("/get_today", controllers.GetIncomeToday)
// 	att.Post("/create", controllers.CreateIncome)
// 	att.Put("/update", controllers.UpdateIncome)
// 	att.Delete("/delete/:id", controllers.DeleteIncome)
// 	att.Get("/:id", controllers.GetIncomeByID)
// }