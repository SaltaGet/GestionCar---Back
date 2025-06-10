package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"


type PermissionService interface {
	PermissionByRoleID(roleID string) (permissions *[]string, err error)
	PermissionGetAll() (permissions *[]models.Permission, err error)
	PermissionGetToMe(roleID string) (permissions *[]models.Permission, err error)
}

type PermissionRepository interface {
	PermissionByRoleID(roleID string) (permissions *[]string, err error)
	PermissionGetAll() (permissions *[]models.Permission, err error)
	PermissionGetToMe(roleID string) (permissions *[]models.Permission, err error)
}