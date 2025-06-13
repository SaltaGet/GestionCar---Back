package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (t *MemberRepository) MemberGetAll() (*[]models.Member, error) {
	var members []models.Member
	if err := t.DB.Preload("Role").Find(&members).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar miembros", err)
	}

	return &members, nil
}

func (t *MemberRepository) MemberGetPermissionByUserID(userID string) (*models.Member, error) {
	var member models.Member
	if err := t.DB.Preload("Role").Preload("Role.Permissions").Where("user_id = ?", userID).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Miembro no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error  internoal buscar miembro", err)
	}
	return &member, nil
}

func (t *MemberRepository) MemberGetByID(id string) (*models.Member, error) {
	var member models.Member
	if err := t.DB.Preload("Role").Preload("Role.Permissions").Where("id = ?", id).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Miembro no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error interno al buscar miembro", err)
	}
	return &member, nil
}

func (t *MemberRepository) MemberCreate(memeberCreate *models.MemberCreate, user *models.AuthenticatedUser) (id string, err error) {
	newID := uuid.NewString()

	// pass, err := utils.GenerateRandomString(10)
	// if err != nil {
	// 	return "", err
	// }
	hashPass, err := utils.HashPassword("1234")
	if err != nil {
		return "", models.ErrorResponse(500, "Error al hashear la contraseña", err)
	}

	if err := t.DB.Create(&models.Member{
		ID:        newID,
		FirstName: memeberCreate.FirstName,
		LastName:  memeberCreate.LastName,
		Username:  fmt.Sprintf("%s@%s", memeberCreate.Username, *user.Identifier),
		Email:     memeberCreate.Email,
		RoleID:    memeberCreate.RoleID,
		Password:  hashPass,
	}).Error; err != nil {
		return "", models.ErrorResponse(500, "Error interno al crear miembro", err)
	}
	return newID, nil
}

func (t *MemberRepository) MemberDelete(id string) error {
	if err := t.DB.Delete(&models.Member{}, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Miembro no encontrado", err)
		}
		return models.ErrorResponse(500, "Error interno al eliminar miembro", err)
	}
	return nil
}