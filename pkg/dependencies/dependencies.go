package dependencies

import (
	"github.com/DanielChachagua/GestionCar/pkg/repositories"
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
	EntityController *entCtrl.Controller
	UserController *userCtrl.Controller
	AuthController *authCtrl.Controller
	EstablishmentController *estCtrl.Controller
}

func NewApplication(db *gorm.DB) *Application {
	entityRepo := &entRep.Repository{DB: db,}
	entityServ := &entServ.Service{EntityRepository: entityRepo}

	userRepo := &userRep.Repository{DB: db,}
	userServ := &userServ.Service{UserRepository: userRepo}

	establishmentRepo := &estRep.Repository{DB: db,}
	establishmentServ := &estServ.Service{EstablishmentRepository: establishmentRepo, EntityRepository: entityRepo}
	
	authRepo := &authRep.Repository{DB: db,}
	authServ := &authServ.Service{AuthRepository: authRepo, UserRepository: userRepo, EstablishmentRepository: establishmentRepo}

	return &Application{
		EntityController: &entCtrl.Controller{EntityService: entityServ},
		UserController: &userCtrl.Controller{UserService: userServ},
		AuthController: &authCtrl.Controller{AuthService: authServ},
		EstablishmentController: &estCtrl.Controller{EstablishmentService: establishmentServ},
	}
}

func (app *Application) SetDBRepository(db *sql.DB) {
	entityRepo := &entRep.Repository{DB: db,}
	entityServ := &entServ.Service{EntityRepository: entityRepo}

	userRepo := &userRep.Repository{DB: db,}
	userServ := &userServ.Service{UserRepository: userRepo}

	authRepo := &authRep.Repository{DB: db,}
	authServ := &authServ.Service{AuthRepository: authRepo, UserRepository: userRepo}

	establishmentRepo := &estRep.Repository{DB: db,}
	establishmentServ := &estServ.Service{EstablishmentRepository: establishmentRepo, EntityRepository: entityRepo}

	app.UserController.UserService = userServ
	app.EntityController.EntityService = entityServ
	app.AuthController.AuthService = authServ
	app.EstablishmentController.EstablishmentService = establishmentServ
}