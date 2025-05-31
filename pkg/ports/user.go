package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type UserService interface {
	Create(user *models.UserCreate) (id string, err error)
	Update(user *models.UserUpdate) (err error)
}

type UserRepository interface {
	Create(user *models.UserCreate) (id string, err error)
	// Update(user *user.UserUpdate) (err error)
	GetByIdentifier(identifier string) (user *models.User, err error) 
	GetUserByID(id string) (user *models.User, err error) 
	GetUserByUsername(username string) (user *models.User, err error) 
	GetByEmail(email string) (user *models.User, err error) 
	ExistUser(identifier string, email string) (exist bool, err error) 
}