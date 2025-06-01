package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type AuthRepository interface {
	AuthLogin(username string, password string) (user *models.User, err error)
}

type AuhtService interface {
	AuthLogin(username string , password string) (token string, err error)
}
