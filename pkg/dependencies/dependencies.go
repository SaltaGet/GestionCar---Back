package dependencies

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/pkg/repositories"
	"github.com/DanielChachagua/GestionCar/pkg/services"
	"gorm.io/gorm"
)

var App *Application

type Application struct {
	AuthController *controllers.AuthController
	UserController *controllers.UserController
	TenantController *controllers.TenantController
}

func NewApplication(mainDB *gorm.DB) *Application {
	mainRepo := &repositories.MainRepository{DB: mainDB}

	authServ := &services.AuthService{AuthRepository: mainRepo, UserRepository: mainRepo}
	userServ := &services.UserService{UserRepository: mainRepo}
	tenantServ := &services.TenantService{TenantRepository: mainRepo}

	return &Application{
		AuthController: &controllers.AuthController{AuthService: authServ},
		UserController: &controllers.UserController{UserService: userServ},
		TenantController: &controllers.TenantController{TenantService: tenantServ},
	}
}

var TenantApp *TenantApplication

type TenantApplication struct {
	AttendanceController *controllers.AttendanceController
	ClientController *controllers.ClientController
	EmployeeController *controllers.EmployeeController
	ExpenseController *controllers.ExpenseController	
	IncomeController *controllers.IncomeController
	MovementTypeController *controllers.MovementTypeController
	ProductController *controllers.ProductController
	PurchaseOrderController *controllers.PurchaseOrderController
	PurchaseProductController *controllers.PurchaseProductController
	ServiceController *controllers.ServiceController
	SupplierController *controllers.SupplierController
	VehicleController *controllers.VehicleController
	ResumeController *controllers.ResumeController
	MemberController *controllers.MemberController
	RoleController *controllers.RoleController
	PermissionController *controllers.PermissionController
}

func TenantDBRepository(db *gorm.DB) *TenantApplication {
	tenantRepo := &repositories.TenantRepository{DB: db,}
	
	memberService := &services.MemberService{MemberRepository: tenantRepo}
	roleService := &services.RoleService{RoleRepository: tenantRepo}
	permissionService := &services.PermissionService{PermissionRepository: tenantRepo}


	return &TenantApplication{
		MemberController: &controllers.MemberController{MemberService: memberService},
		RoleController: &controllers.RoleController{RoleService: roleService},
		PermissionController: &controllers.PermissionController{PermissionService: permissionService},
	}
}


func (app *TenantApplication) SetDBTenantRepository(db *gorm.DB) {
	tenantRepo := &repositories.TenantRepository{DB: db,}
	
	attendanceService := &services.AttendanceService{AttendanceRepository: tenantRepo}
	clientService := &services.ClientService{ClientRepository: tenantRepo}
	employeeService := &services.EmployeeService{EmployeeRepository: tenantRepo}
	expenseService := &services.ExpenseService{ExpenseRepository: tenantRepo}
	incomeService := &services.IncomeService{IncomeRepository: tenantRepo}
	memberService := &services.MemberService{MemberRepository: tenantRepo}
	movementService := &services.MovementTypeService{MovementTypeRepository: tenantRepo}
	permissionService := &services.PermissionService{PermissionRepository: tenantRepo}
	productService := &services.ProductService{ProductRepository: tenantRepo}
	purchaseOrderService := &services.PurchaseOrderService{PurchaseOrderRepository: tenantRepo}
	purchaseProductService := &services.PurchaseProductService{PurchaseProductRepository: tenantRepo}
	resumeService := &services.ResumeService{ResumeExpenseRepository: tenantRepo, ResumeIncomeRepository: tenantRepo}
	roleService := &services.RoleService{RoleRepository: tenantRepo}
	serviceService := &services.ServiceService{ServiceRepository: tenantRepo}
	supplierService := &services.SupplierService{SupplierRepository: tenantRepo}
	vehicleService := &services.VehicleService{VehicleRepository: tenantRepo}

	app.AttendanceController.AttendanceService = attendanceService
	app.ClientController.ClientService = clientService
	app.EmployeeController.EmployeeService = employeeService
	app.ExpenseController.ExpenseService = expenseService
	app.IncomeController.IncomeService = incomeService
	app.MemberController.MemberService = memberService
	app.MovementTypeController.MovementService = movementService
	app.PermissionController.PermissionService = permissionService
	app.ProductController.ProductService = productService
	app.PurchaseOrderController.PurchaseOrderService = purchaseOrderService
	app.PurchaseProductController.PurchaseProductService = purchaseProductService
	app.ResumeController.ExpenseResumeService = resumeService
	app.RoleController.RoleService = roleService
	app.ServiceController.ServiceService = serviceService
	app.SupplierController.SupplierService = supplierService
	app.VehicleController.VehicleService = vehicleService
}