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

func NewApplication(db *gorm.DB) *Application {
	repo := &repositories.Repository{DB: db,}
	
	authServ := &services.AuthService{AuthRepository: repo, UserRepository: repo}
	userServ := &services.UserService{UserRepository: repo}
	tenantServ := &services.TenantService{TenantRepository: repo}

	return &Application{
		AuthController: &controllers.AuthController{AuthService: authServ},
		UserController: &controllers.UserController{UserService: userServ},
		TenantController: &controllers.TenantController{TenantService: tenantServ},
	}
}

func (app *Application) SetDBRepository(db *gorm.DB) {
	repo := &repositories.Repository{DB: db,}
	
	authServ := &services.AuthService{AuthRepository: repo, UserRepository: repo}
	userServ := &services.UserService{UserRepository: repo}
	tenantServ := &services.TenantService{TenantRepository: repo}

	app.AuthController.AuthService = authServ
	app.UserController.UserService = userServ
	app.TenantController.TenantService = tenantServ
}