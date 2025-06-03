package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type AuthRepository interface {
	AuthLogin(username string, password string) (user *models.User, err error)
	AuthGetTenant(userID string, tenantID string) (tenant *models.Tenant, err error)
	CurrentUser(userID string) (user *models.User, err error)
	UserGetRolePermissions(connection, userID string) (member *models.Member, role *models.Role, permissions *[]string, err error)
}

type AuhtService interface {
	AuthLogin(username string , password string) (token string, err error)
	AuthGetTenant(user *models.User, tenantID string) (token string, err error)
	CurrentUser(userID string) (user *models.User, err error)
}
