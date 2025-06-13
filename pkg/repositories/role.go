package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
)

func (r *RoleRepository) RoleGetAll() (*[]models.Role, error) {
	var allRoles []models.Role
	if err := r.DB.Find(&allRoles).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al obtener roles", err)
	}
	return &allRoles, nil
}

func (t *RoleRepository) RoleCreate(roleCreate *models.RoleCreate) (string, error) {
	newID := uuid.NewString()
	err := t.DB.Create(&models.Role{ID: newID, Name: roleCreate.Name}).Error
	if err != nil {
		return "", models.ErrorResponse(500, "Error interno al crear el rol", err)
	}
	return newID, nil
}
