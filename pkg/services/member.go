package services

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
)

func (m *MemberService) MemberGetRolePermissions(user *models.User, tenant *models.Tenant) (string, error) {
	member, role, permissions, err := m.MemberRepository.MemberGetRolePermissions(user.ID)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user, tenant.ID, member.ID, role, permissions)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al generar token", err)
	}

	return token, nil
}
