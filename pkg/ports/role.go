package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type RoleService interface {
	GetRoleAll(role string) (roles *[]models.Role, err error)
}

type RoleRepository interface {
	GetRoleAll(role string) (roles *[]models.Role, err error)
}