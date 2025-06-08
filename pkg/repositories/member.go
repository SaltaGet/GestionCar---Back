package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (t *TenantRepository) MemberGetAll() (*[]models.Member, error) {
	var members []models.Member
	if err := t.DB.Preload("Role").Find(&members).Error; err != nil {
		return nil, err
	}

	return &members, nil
}

func (t *TenantRepository) MemberGetPermissionByUserID(userID string) (*models.Member, error) {
	var member models.Member
	if err := t.DB.Preload("Role").Preload("Role.Permissions").Where("user_id = ?", userID).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Miembro no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar miembro", err)
	}
	return &member, nil
}

func (t *TenantRepository) MemberGetByID(id string) (*models.Member, error) {
	var member models.Member
	if err := t.DB.Preload("Role").Where("id = ?", id).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Miembro no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar miembro", err)
	}
	return &member, nil
}

func (t *TenantRepository) MemberAdd(memberAdd *models.MemberAdd) (id string, err error) {
	newID := uuid.NewString()
	if err := t.DB.Create(&models.Member{
		ID: newID,
		UserID: memberAdd.UserID,
		RoleID: memberAdd.RoleID,
	}).Error; err != nil {
		return "", err
	}

	return newID, nil
}


func (t *TenantRepository) MemberDelete(id string) error {
	if err := t.DB.Delete(&models.Member{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}