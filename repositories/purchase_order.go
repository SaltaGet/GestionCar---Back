package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
)

func (r *Repository) GetPurchaseOrderByID(id string, workplace string) (*models.PurchaseOrderLaundry, *models.PurchaseOrderWorkshop, error) {
	if workplace == "laundry" {
		var purchaseOrder models.PurchaseOrderLaundry
		if err := r.DB.Where("id = ?", id).First(&purchaseOrder).Error; err != nil {
			return nil, nil, err
		}
		return &purchaseOrder, nil, nil
	} else if workplace == "workshop" {
		var purchaseOrder models.PurchaseOrderWorkshop
		if err := r.DB.Where("id = ?", id).First(&purchaseOrder).Error; err != nil {
			return nil, nil, err
		}
		return nil, &purchaseOrder, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetAllPurchaseOrders(workplace string) ([]models.PurchaseOrderLaundry, []models.PurchaseOrderWorkshop, error) {
	if workplace == "laundry" {
		var purchaseOrders []models.PurchaseOrderLaundry
		if err := r.DB.Find(&purchaseOrders).Error; err != nil {
			return nil, nil, err
		}
		return purchaseOrders, nil, nil
	} else if workplace == "workshop" {
		var purchaseOrders []models.PurchaseOrderWorkshop
		if err := r.DB.Find(&purchaseOrders).Error; err != nil {
			return nil, nil, err
		}
		return nil, purchaseOrders, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) CreatePurchaseOrder(purchaseOrder interface{}) (string, error) {
	switch p := purchaseOrder.(type) {
	case *models.PurchaseOrderLaundry:
		if err := r.DB.Create(p).Error; err != nil {
			return "", err
		}
		return p.ID, nil
	case *models.PurchaseOrderWorkshop:
		if err := r.DB.Create(p).Error; err != nil {
			return "", err
		}
		return p.ID, nil
	default:
		return "", fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdatePurchaseOrder(purchaseOrder interface{}) error {
	switch p := purchaseOrder.(type) {
	case *models.PurchaseOrderLaundry:
		if err := r.DB.Save(p).Error; err != nil {
			return err
		}
		return nil
	case *models.PurchaseOrderWorkshop:
		if err := r.DB.Save(p).Error; err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) DeletePurchaseOrderByID(id string, workplace string) error {
	if workplace == "laundry" {
		var purchaseOrder models.PurchaseOrderLaundry
		if err := r.DB.Where("id = ?", id).Delete(&purchaseOrder).Error; err != nil {
			return err
		}
		return nil
	} else if workplace == "workshop" {
		var purchaseOrder models.PurchaseOrderWorkshop
		if err := r.DB.Where("id = ?", id).Delete(&purchaseOrder).Error; err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("tipo de espacio no soportado")
}


// var order PurchaseOrderWorkshop

// err := db.Preload("Supplier").
//           Preload("PurchaseParts.PartWorkshop").
//           First(&order, "id = ?", someID).Error

// if err != nil {
//     // manejar error
// }