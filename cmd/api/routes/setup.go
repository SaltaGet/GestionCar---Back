package routes

import (
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, appDependencies *dependencies.Application, tenantDependencies *dependencies.TenantApplication) {
	// AttendanceRoutes(app, tenantDependencies.AttendanceController)
	AuthRoutes(app, appDependencies.AuthController)
	// ClientRoutes(app, tenantDependencies.ClientController)
	// EmployeeRoutes(app, tenantDependencies.EmployeeController)
	// ExpenseRoutes(app, tenantDependencies.ExpenseController)
	// IncomeRoutes(app, tenantDependencies.IncomeController)
	MemberRoutes(app, tenantDependencies.MemberController)
	// MovementRoutes(app, tenantDependencies.MovementController)
	// ProductRoutes(app, tenantDependencies.ProductController)
	// PurchaseOrderRoutes(app, tenantDependencies.PurchaseOrderController)
	// PurchaseProductRoutes(app, tenantDependencies.PurchaseProductController)
	// RoleRoutes(app, tenantDependencies.RoleController)
	// ServiceRoutes(app, tenantDependencies.ServiceController)
	// SupplierRoutes(app, tenantDependencies.SupplierController)
	UserRoutes(app, appDependencies.UserController)
	// VehicleRoutes(app, tenantDependencies.VehicleController)
	TenantRoutes(app, appDependencies.TenantController)
}