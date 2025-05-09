package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
)

func (r *Repository) GetPurchaseElementByID(id string, workplace string) (*models.PurchaseProductLaundry, *models.PurchasePartWorkshop, error) {
	if workplace == "laundry" {
		var purchaseProduct models.PurchaseProductLaundry
		if err := r.DB.Where("id = ?", id).First(&purchaseProduct).Error; err != nil {
			return nil, nil, err
		}
		return &purchaseProduct, nil, nil
	} else if workplace == "workshop" {
		var purchasePart models.PurchasePartWorkshop
		if err := r.DB.Where("id = ?", id).First(&purchasePart).Error; err != nil {
			return nil, nil, err
		}
		return nil, &purchasePart, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetAllPurchaseElements(workplace string) ([]models.PurchaseProductLaundry, []models.PurchasePartWorkshop, error) {
	if workplace == "laundry" {
		var purchaseProducts []models.PurchaseProductLaundry
		if err := r.DB.Find(&purchaseProducts).Error; err != nil {
			return nil, nil, err
		}
		return purchaseProducts, nil, nil
	} else if workplace == "workshop" {
		var purchaseParts []models.PurchasePartWorkshop
		if err := r.DB.Find(&purchaseParts).Error; err != nil {
			return nil, nil, err
		}
		return nil, purchaseParts, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) CreatePurchaseElement(element interface{}) (string, error) {
	switch e := element.(type) {
	case *models.PurchaseProductLaundry:
		if err := r.DB.Create(e).Error; err != nil {
			return "", err
		}
		return e.ID, nil
	case *models.PurchasePartWorkshop:
		if err := r.DB.Create(e).Error; err != nil {
			return "", err
		}
		return e.ID, nil
	default:
		return "", fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdatePurchaseElement(element interface{}) error {
	switch e := element.(type) {
	case *models.PurchaseProductLaundry:
		if err := r.DB.Save(e).Error; err != nil {
			return err
		}
	case *models.PurchasePartWorkshop:
		if err := r.DB.Save(e).Error; err != nil {
			return err
		}
	default:
		return fmt.Errorf("tipo de movimiento no soportado")
	}
	return nil
}

func (r *Repository) DeletePurchaseElementByID(id string, workplace string) error {
	if workplace == "laundry" {
		var purchaseProduct models.PurchaseProductLaundry
		if err := r.DB.Where("id = ?", id).Delete(&purchaseProduct).Error; err != nil {
			return err
		}
	} else if workplace == "workshop" {
		var purchasePart models.PurchasePartWorkshop
		if err := r.DB.Where("id = ?", id).Delete(&purchasePart).Error; err != nil {
			return err
		}
	} else {
		return fmt.Errorf("tipo de movimiento no soportado")
	}
	return nil
}