package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func (r *TenantRepository) RoleGetAll(roleName string) (*[]models.Role, error) {
	var currentRole models.Role

	if err := r.DB.Where("name = ?", roleName).First(&currentRole).Error; err != nil {
		return nil, err
	}

	var allRoles []models.Role
	if err := r.DB.Find(&allRoles).Error; err != nil {
		return nil, err
	}
	return &allRoles, nil
}
