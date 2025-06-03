package repositories

import "github.com/DanielChachagua/GestionCar/pkg/models"

func (t *TenantRepository) MemberGetAll() (*[]models.Member, error) {
	var members []models.Member
	if err := t.DB.Find(&members).Error; err != nil {
		return nil, err
	}

	return &members, nil
}