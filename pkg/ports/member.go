package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type MemberRepository interface {
	MemberGetRolePermissions(userID string) (member *models.Member, role *models.Role, permissions *[]string, err error)
}

type MemberService interface {
	MemberGetRolePermissions(user *models.User, tenant *models.Tenant) (token string, err error)
}