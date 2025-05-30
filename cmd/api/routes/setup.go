package routes

import (
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, appDependencies *dependencies.Application) {
	AttendanceRoutes(app)
	AuthRoutes(app)
	ClientRoutes(app)
	EmployeeRoutes(app)
	ExpenseRoutes(app)
	IncomeRoutes(app)
	MovementRoutes(app)
	ProductRoutes(app)
	PurchaseOrderRoutes(app)
	PurchaseProductRoutes(app)
	RoleRoutes(app)
	ServiceRoutes(app)
	SupplierRoutes(app)
	UserRoutes(app)
	VehicleRoutes(app)
	WorkplaceRoutes(app)
}