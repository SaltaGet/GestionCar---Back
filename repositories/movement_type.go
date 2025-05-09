package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
)

func (r *Repository) GetMovementTypeByID(id string, workplace string) (*models.MovementTypeLaundry, *models.MovementTypeWorkshop, error) {
	if workplace == "laundry" {
		var movementType models.MovementTypeLaundry
		if err := r.DB.Where("id = ?", id).First(&movementType).Error; err != nil {
			return nil, nil, err
		}
		return &movementType, nil, nil
	} else if workplace == "workshop" {
		var movementType models.MovementTypeWorkshop
		if err := r.DB.Where("id = ?", id).First(&movementType).Error; err != nil {
			return nil, nil, err
		}
		return nil, &movementType, nil
	}
	return nil, nil, fmt.Errorf("tipo de espacio no soportado")
}

func (r *Repository) GetAllMovementTypes(workplace string) ([]models.MovementTypeLaundry, []models.MovementTypeWorkshop, error) {
	if workplace == "laundry" {
		var movementTypes []models.MovementTypeLaundry
		if err := r.DB.Find(&movementTypes).Error; err != nil {
			return nil, nil, err
		}
		return movementTypes, nil, nil
	} else if workplace == "workshop" {
		var movementTypes []models.MovementTypeWorkshop
		if err := r.DB.Find(&movementTypes).Error; err != nil {
			return nil, nil, err
		}
		return nil, movementTypes, nil
	}
	return nil, nil, fmt.Errorf("tipo de espacio no soportado")
}

func (r *Repository) CreateMovementType(movementType *models.MovementTypeCreate, workplace string) (string, error) {
	newID := uuid.NewString()
	switch workplace{
		case "laundry":
			if err := r.DB.Create(&models.MovementTypeLaundry{
				ID: newID,
				Name: movementType.Name,
				IsIncome: movementType.IsIncome,
			}).Error; err != nil {
				return "", err
			}
			return newID, nil
		case "workshop":
			if err := r.DB.Create(&models.MovementTypeWorkshop{
				ID: newID,
				Name: movementType.Name,
				IsIncome: movementType.IsIncome,
			}).Error; err != nil {
				return "", err
			}
			return newID, nil
		default:
			return "", fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdateMovementType(movementTypeUpdate *models.MovementTypeUpdate, workplace string) error {
	switch workplace{
		case "laundry":
			if err := r.DB.Model(&models.MovementTypeLaundry{}).Where("id = ?", movementTypeUpdate.ID).Updates(&models.MovementTypeLaundry{
				Name: movementTypeUpdate.Name,
				IsIncome: movementTypeUpdate.IsIncome,
			}).Error; err != nil {
				return err
			}
			return nil
		case "workshop":
			if err := r.DB.Model(&models.MovementTypeWorkshop{}).Where("id = ?", movementTypeUpdate.ID).Updates(&models.MovementTypeWorkshop{
				Name: movementTypeUpdate.Name,
				IsIncome: movementTypeUpdate.IsIncome,
			}).Error; err != nil {
				return err
			}
			return nil
		default:
			return fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) DeleteMovementType(id string, workplace string) error {
	if workplace == "laundry" {
		var movementType models.MovementTypeLaundry
		if err := r.DB.Where("id = ?", id).Delete(&movementType).Error; err != nil {
			return err
		}
		return nil
	} else if workplace == "workshop" {
		var movementType models.MovementTypeWorkshop
		if err := r.DB.Where("id = ?", id).Delete(&movementType).Error; err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("tipo de espacio no soportado")
}