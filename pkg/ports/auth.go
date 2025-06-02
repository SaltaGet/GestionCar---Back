package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type AuthRepository interface {
	AuthLogin(username string, password string) (user *models.User, err error)
	AuthTenant(userID string, tenantID string) (user *models.Tenant, err error)
	CurrentUser(userID string) (user *models.User, err error)
	GetUserRolePermissions(connection, userID string) (user *models.User, role*models.Role, permissions *[]models.Permission, err error)
}

type AuhtService interface {
	AuthLogin(username string , password string) (token string, err error)
	AuthTenant(userID string, tenantID string) (token string, err error)
}
