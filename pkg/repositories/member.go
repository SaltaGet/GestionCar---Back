package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (t *TenantRepository) MemberGetAll() (*[]models.Member, error) {
	var members []models.Member
	if err := t.DB.Preload("Role").Find(&members).Error; err != nil {
		return nil, err
	}

	return &members, nil
}

func (t *TenantRepository) MemeberGetPermissionByUserID(userID string) (*models.Member, error) {
	var member models.Member
	if err := t.DB.Preload("Role").Preload("Role.Permissions").Where("user_id = ?", userID).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Miembro no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar miembro", err)
	}
	return &member, nil
}