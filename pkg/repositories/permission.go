package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

func (t *TenantRepository) PermissionByRoleID(roleID string) (*[]string, error) {
	var permission []string
	err := t.DB.Model(&models.Permission{}).
		Select("permissions.code").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", roleID).
		Pluck("permissions.code", &permission).Error
	if err != nil {
		return nil, models.ErrorResponse(500, "Error interno al obtener permisos", err)
	}
	return &permission, nil
}

func (t *TenantRepository) PermissionGetAll() (*[]models.Permission, error) {
	var permission []models.Permission
	err := t.DB.Find(&permission).Error
	if err != nil {
		return nil, models.ErrorResponse(500, "Error interno al obtener permisos", err)
	}
	return &permission, nil
}

func (t *TenantRepository) PermissionGetToMe(roleID string) (*[]models.Permission, error) {
	var permissions []models.Permission
	err := t.DB.
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", roleID).
		Find(&permissions).Error
	if err != nil {
		return nil, models.ErrorResponse(500, "Error interno al obtener permisos", err)
	}
	return &permissions, nil
}
