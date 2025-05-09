package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
)

func (r *Repository) GetMovementByID(id string, workspace string) (*models.DailyMovementsLaundry, *models.DailyMovementsWorkshop, error) {
	if workspace == "laundry" {
		var movement models.DailyMovementsLaundry
		if err := r.DB.Where("id = ?", id).First(&movement).Error; err != nil {
			return nil, nil, err
		}
		return &movement, nil, nil
	} else if workspace == "workshop" {
		var movement models.DailyMovementsWorkshop
		if err := r.DB.Where("id = ?", id).First(&movement).Error; err != nil {
			return nil, nil, err
		}
		return nil, &movement, nil
	}
	return nil, nil, nil
}

func (r *Repository) GetAllMovements(id string, workspace string) ([]models.DailyMovementsLaundry, []models.DailyMovementsWorkshop, error) {
	if workspace == "laundry" {
		var movements []models.DailyMovementsLaundry
		if err := r.DB.Where("id = ?", id).Find(&movements).Error; err != nil {
			return nil, nil, err
		}
		return movements, nil, nil
	} else if workspace == "workshop" {
		var movements []models.DailyMovementsWorkshop
		if err := r.DB.Where("id = ?", id).Find(&movements).Error; err != nil {
			return nil, nil, err
		}
		return nil, movements, nil
	}
	return nil, nil, nil
}

func (r *Repository) CreateMovement(movement interface{}) (string, error) {
	switch m := movement.(type) {
	case *models.DailyMovementsLaundry:
		if err := r.DB.Create(m).Error; err != nil {
			return "", err
		}
		return m.ID, nil
	case *models.DailyMovementsWorkshop:
		if err := r.DB.Create(m).Error; err != nil {
			return "", err
		}
		return m.ID, nil
	default:
		return "", fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdateMovement(movement interface{}) error {
	switch m := movement.(type) {
	case *models.DailyMovementsLaundry:
		if err := r.DB.Save(m).Error; err != nil {
			return err
		}
		return nil
	case *models.DailyMovementsWorkshop:
		if err := r.DB.Save(m).Error; err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) DeleteMovement(id string, workspace string) error {
	if workspace == "laundry" {
		var movement models.DailyMovementsLaundry
		if err := r.DB.Where("id = ?", id).Delete(&movement).Error; err != nil {
			return err
		}
		return nil
	} else if workspace == "workshop" {
		var movement models.DailyMovementsWorkshop
		if err := r.DB.Where("id = ?", id).Delete(&movement).Error; err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) EndDay(id string, workspace string) error {
	if workspace == "laundry" {
		var movement models.DailyMovementsLaundry
		if err := r.DB.Where("id = ?", id).Delete(&movement).Error; err != nil {
			return err
		}
		return nil
	} else if workspace == "workshop" {
		var movement models.DailyMovementsWorkshop
		if err := r.DB.Where("id = ?", id).Delete(&movement).Error; err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("tipo de movimiento no soportado")
}