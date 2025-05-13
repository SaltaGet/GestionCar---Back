package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
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

func (r *Repository) GetPurchaseElementByPurchaseID(purchaseID string, workplace string) (*[]models.PurchaseProductLaundry, *[]models.PurchasePartWorkshop, error) {
	if workplace == "laundry" {
		var purchaseProduct []models.PurchaseProductLaundry
		if err := r.DB.Where("purchase_order_id = ?", purchaseID).Find(&purchaseProduct).Error; err != nil {
			return nil, nil, err
		}
		return &purchaseProduct, nil, nil
	} else if workplace == "workshop" {
		var purchasePart []models.PurchasePartWorkshop
		if err := r.DB.Where("purchase_order_id = ?", purchaseID).Find(&purchasePart).Error; err != nil {
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

func (r *Repository) CreatePurchaseElement(element *models.PurchaseProductCreate, workplace string) (string, error) {
	newID := uuid.NewString()
	switch workplace {
	case "laundry":
		if err := r.DB.Create(&models.PurchaseProductLaundry{
			ID: newID,
			ProductID: element.ProductID,
			ExpiredAt: element.ExpiredAt,
			UnitPrice: element.UnitPrice,
			Quantity: element.Quantity,
			TotalPrice: element.UnitPrice * float32(element.Quantity),
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	case "workshop":
		if err := r.DB.Create(&models.PurchasePartWorkshop{
			ID: newID,
			PartID: element.ProductID,
			ExpiredAt: element.ExpiredAt,
			UnitPrice: element.UnitPrice,
			Quantity: element.Quantity,
			TotalPrice: element.UnitPrice * float32(element.Quantity),
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	default:
		return "", fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdatePurchaseElement(element *models.PurchaseProductUpdate, workplace string) error {
	switch workplace {
	case "laundry":
		if err := r.DB.Where("id = ?", element.ID).Updates(&models.PurchaseProductLaundry{
			ProductID: element.ProductID,
			ExpiredAt: element.ExpiredAt,
			UnitPrice: element.UnitPrice,
			Quantity: element.Quantity,
			TotalPrice: element.UnitPrice * float32(element.Quantity),
		}).Error; err != nil {
			return err
		}
		return nil
	case "workshop":
		if err := r.DB.Where("id = ?", element.ID).Updates(&models.PurchasePartWorkshop{
			PartID: element.ProductID,
			ExpiredAt: element.ExpiredAt,
			UnitPrice: element.UnitPrice,
			Quantity: element.Quantity,
			TotalPrice: element.UnitPrice * float32(element.Quantity),
		}).Error; err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("tipo de movimiento no soportado")
	}
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