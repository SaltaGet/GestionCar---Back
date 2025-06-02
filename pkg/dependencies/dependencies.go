package dependencies

import (
	"github.com/DanielChachagua/GestionCar/cmd/api/controllers"
	"github.com/DanielChachagua/GestionCar/pkg/repositories"
	"github.com/DanielChachagua/GestionCar/pkg/services"
	"gorm.io/gorm"
)

// type Dependency struct {
// 	Repository *repositories.Repository
// }

// func NewDependency(db *gorm.DB) *Dependency {

// 	repo := &repositories.Repository{
// 		DB: db,
// 	}

// 	return &Dependency{
// 		Repository: repo,
// 	}
// }

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