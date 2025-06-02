package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

func (r *TenantRepository) MemberGetRolePermissions(userID string) (*models.Member, *models.Role, *[]string, error) {
	var member models.Member
	if err := r.DB.Where("user_id = ?", userID).First(&member).Error; err != nil {
		return nil, nil, nil, models.ErrorResponse(404, "User not found", err)
	}

	var role models.Role
	if err := r.DB.Where("id = ?", member.RoleID).First(&role).Error; err != nil {
		return nil, nil, nil, models.ErrorResponse(404, "Role not found", err)
	}

	var permissions []string
	err := r.DB.Model(&models.Permission{}).
		Select("permissions.name").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", role.ID).
		Pluck("permissions.name", &permissions).Error
	if err != nil {
		return nil, nil, nil, models.ErrorResponse(500, "Error al obtener los permisos", err)
	}

	return &member, &role, &permissions, nil
}