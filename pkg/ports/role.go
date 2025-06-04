package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type RoleService interface {
	RoleGetAll(role string) (roles *[]models.Role, err error)
}

type RoleRepository interface {
	RoleGetAll(role string) (roles *[]models.Role, err error)
}