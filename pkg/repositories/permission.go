package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

func (t *TenantRepository) PermissionByRoleID(roleID string) (permissions *[]string, err error) {
	var permission []string
	err = t.DB.Model(&models.Permission{}).
		Select("permissions.name").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Where("role_permissions.role_id = ?", roleID).
		Pluck("permissions.name", &permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}