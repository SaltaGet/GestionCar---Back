package routes

import (
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, appDependencies *dependencies.Application, tenantDependencies *dependencies.TenantApplication) {
	// AttendanceRoutes(app)
	AuthRoutes(app, appDependencies.AuthController)
	// ClientRoutes(app)
	// EmployeeRoutes(app)
	// ExpenseRoutes(app)
	// IncomeRoutes(app)
	MemberRoutes(app, tenantDependencies.MemberController)
	// MovementRoutes(app)
	// ProductRoutes(app)
	// PurchaseOrderRoutes(app)
	// PurchaseProductRoutes(app)
	// RoleRoutes(app)
	// ServiceRoutes(app)
	// SupplierRoutes(app)
	UserRoutes(app, appDependencies.UserController)
	// VehicleRoutes(app)
	TenantRoutes(app, appDependencies.TenantController)
}