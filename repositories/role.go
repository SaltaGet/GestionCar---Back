package repositories

import (
	"github.com/DanielChachagua/GestionCar/models"
)

func (r *Repository) GetAllRoles(roleName string, workplace string) (*[]models.Role, error) {
	var currentRole models.Role

	if err := r.DB.Where("name = ?", roleName).First(&currentRole).Error; err != nil {
    return nil, err
	}

	var allRoles []models.Role
	if workplace == "laundry" {
		if err := r.DB.Where("hierarchy > ? AND workplace = ?", currentRole.Hierarchy, "laundry").Find(&allRoles).Error; err != nil {
			return nil, err
		}
	} else if workplace == "workshop" {
		if err := r.DB.Where("hierarchy > ? AND workplace = ?", currentRole.Hierarchy, "workshop").Find(&allRoles).Error; err != nil {
			return nil, err
		}
	}
	return &allRoles, nil
}