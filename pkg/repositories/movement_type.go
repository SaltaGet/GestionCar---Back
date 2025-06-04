package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
)

func (r *TenantRepository) GetMovementTypeByID(id string) (*models.MovementType, error) {
		var movementType models.MovementType
		if err := r.DB.Where("id = ?", id).First(&movementType).Error; err != nil {
			return nil, err
		}
		return &movementType, nil
}

func (r *TenantRepository) GetAllMovementTypes(isIncome bool) (*[]models.MovementType, error) {
		var movementTypes []models.MovementType
		if err := r.DB.Where("is_income = ?", isIncome).Find(&movementTypes).Error; err != nil {
			return nil, err
		}
		return &movementTypes, nil
}

func (r *TenantRepository) CreateMovementType(movementType *models.MovementTypeCreate) (string, error) {
	newID := uuid.NewString()
			if err := r.DB.Create(&models.MovementType{
				ID: newID,
				Name: movementType.Name,
				IsIncome: movementType.IsIncome,
			}).Error; err != nil {
				return "", err
			}
			return newID, nil
}

func (r *TenantRepository) UpdateMovementType(movementTypeUpdate *models.MovementTypeUpdate) error {
			if err := r.DB.Model(&models.MovementType{}).Where("id = ?", movementTypeUpdate.ID).Updates(&models.MovementType{
				Name: movementTypeUpdate.Name,
				IsIncome: movementTypeUpdate.IsIncome,
			}).Error; err != nil {
				return err
			}
			return nil
}

func (r *TenantRepository) DeleteMovementType(id string) error {
		var movementType models.MovementType
		if err := r.DB.Where("id = ?", id).Delete(&movementType).Error; err != nil {
			return err
		}
		return nil
}