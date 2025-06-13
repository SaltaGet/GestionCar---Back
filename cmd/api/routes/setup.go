package routes

import (
	// "github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, appDependencies *dependencies.Application) {
	// app.Use(func(c *fiber.Ctx) error {
  //       c.Locals("AttendanceController", new(*controllers.AttendanceController))
  //       c.Locals("MemberController", new(*controllers.MemberController))
  //       return c.Next()
  //   })


	AttendanceRoutes(app)
	AuthRoutes(app, appDependencies.AuthController)
	ClientRoutes(app)
	EmployeeRoutes(app)
	ExpenseRoutes(app)
	IncomeRoutes(app)
	MemberRoutes(app)
	MovementRoutes(app)
	ProductRoutes(app)
	PurchaseOrderRoutes(app)
	// PurchaseProductRoutes(app, tenantDependencies.PurchaseProductController)
	RoleRoutes(app)
	ServiceRoutes(app)
	SupplierRoutes(app)
	UserRoutes(app, appDependencies.UserController)
	VehicleRoutes(app)
	TenantRoutes(app, appDependencies.TenantController)
	// AttendanceRoutes(app, tenantDependencies.AttendanceController)
	// AuthRoutes(app, appDependencies.AuthController)
	// ClientRoutes(app, tenantDependencies.ClientController)
	// EmployeeRoutes(app, tenantDependencies.EmployeeController)
	// ExpenseRoutes(app, tenantDependencies.ExpenseController)
	// IncomeRoutes(app, tenantDependencies.IncomeController)
	// MemberRoutes(app, tenantDependencies.MemberController)
	// MovementRoutes(app, tenantDependencies.MovementTypeController)
	// ProductRoutes(app, tenantDependencies.ProductController)
	// PurchaseOrderRoutes(app, tenantDependencies.PurchaseOrderController)
	// // PurchaseProductRoutes(app, tenantDependencies.PurchaseProductController)
	// RoleRoutes(app, tenantDependencies.RoleController)
	// ServiceRoutes(app, tenantDependencies.ServiceController)
	// SupplierRoutes(app, tenantDependencies.SupplierController)
	// UserRoutes(app, appDependencies.UserController)
	// VehicleRoutes(app, tenantDependencies.VehicleController)
	// TenantRoutes(app, appDependencies.TenantController)
}

