package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetPurchaseOrderByID(id string) (*models.PurchaseOrder, error) {
	var purchaseOrder models.PurchaseOrder
	if err := r.DB.Where("id = ?", id).First(&purchaseOrder).Error; err != nil {
		return nil, err
	}
	return &purchaseOrder, nil

}

func (r *Repository) GetAllPurchaseOrders() (*[]models.PurchaseOrder, error) {
	var purchaseOrders []models.PurchaseOrder
	if err := r.DB.Find(&purchaseOrders).Error; err != nil {
		return nil, err
	}
	return &purchaseOrders, nil

}

func (r *Repository) CreatePurchaseOrder(purchaseOrder *models.PurchaseOrderCreate) (string, error) {
	newID := uuid.NewString()
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&models.PurchaseOrder{
			ID:          newID,
			OrderNumber: purchaseOrder.OrderNumber,
			OrderDate:   purchaseOrder.OrderDate,
			Amount:      purchaseOrder.Amount,
			SupplierID:  purchaseOrder.SupplierID,
		}).Error; err != nil {
			return err
		}
		for _, element := range purchaseOrder.PurchaseProductCreates {
			if err := tx.Create(&models.PurchaseProduct{
				ID:              uuid.NewString(),
				ProductID:       element.ProductID,
				PurchaseOrderID: newID,
				ExpiredAt:       element.ExpiredAt,
				UnitPrice:       element.UnitPrice,
				Quantity:        element.Quantity,
				TotalPrice:      element.UnitPrice * float32(element.Quantity),
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return newID, nil
}

// func (r *Repository) UpdatePurchaseOrder(purchaseOrder *models.PurchaseOrderUpdate, workplace string) error {
// 	return r.DB.Transaction(func(tx *gorm.DB) error {
// 		switch workplace {
// 		case "laundry":
// 			if err := r.DB.Where("id = ?", purchaseOrder.ID).Updates(&models.PurchaseOrderLaundry{
// 				OrderNumber: purchaseOrder.OrderNumber,
// 				OrderDate: purchaseOrder.OrderDate,
// 				Amount: purchaseOrder.Amount,
// 				SupplierID: purchaseOrder.SupplierID,
// 			}).Error; err != nil {
// 				return err
// 			}

// 			return nil
// 		case "workshop":
// 			if err := r.DB.Where("id = ?", purchaseOrder.ID).Updates(&models.PurchaseOrderWorkshop{
// 				OrderNumber: purchaseOrder.OrderNumber,
// 				OrderDate: purchaseOrder.OrderDate,
// 				Amount: purchaseOrder.Amount,
// 				SupplierID: purchaseOrder.SupplierID,
// 			}).Error; err != nil {
// 				return err
// 			}
// 			return nil
// 		default:
// 			return fmt.Errorf("tipo de movimiento no soportado")
// 		}
// 	})
// }

func (r *Repository) UpdatePurchaseOrder(purchaseOrder *models.PurchaseOrderUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", purchaseOrder.ID).Updates(&models.PurchaseOrder{
			OrderNumber: purchaseOrder.OrderNumber,
			OrderDate:   purchaseOrder.OrderDate,
			Amount:      purchaseOrder.Amount,
			SupplierID:  purchaseOrder.SupplierID,
		}).Error; err != nil {
			return err
		}

		var existingProducts []models.PurchaseProduct
		if err := tx.Where("purchase_order_id = ?", purchaseOrder.ID).Find(&existingProducts).Error; err != nil {
			return err
		}
		existingIDs := map[string]bool{}
		for _, p := range existingProducts {
			existingIDs[p.ID] = true
		}

		receivedIDs := map[string]bool{}
		for _, prod := range purchaseOrder.PurchaseProductUpdates {
			receivedIDs[prod.ID] = true
		}

		for _, p := range existingProducts {
			if !receivedIDs[p.ID] {
				if err := tx.Delete(&models.PurchaseProduct{}, "id = ?", p.ID).Error; err != nil {
					return err
				}
			}
		}

		for _, prod := range purchaseOrder.PurchaseProductUpdates {
			if prod.ID == "" || !existingIDs[prod.ID] {
				newProd := models.PurchaseProduct{
					ID:              uuid.NewString(),
					ProductID:       prod.ProductID,
					PurchaseOrderID: purchaseOrder.ID,
					ExpiredAt:       prod.ExpiredAt,
					UnitPrice:       prod.UnitPrice,
					Quantity:        prod.Quantity,
					TotalPrice:      prod.UnitPrice * float32(prod.Quantity),
				}
				if err := tx.Create(&newProd).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Model(&models.PurchaseProduct{}).
					Where("id = ?", prod.ID).
					Updates(map[string]interface{}{
						"product_id":  prod.ProductID,
						"expired_at":  prod.ExpiredAt,
						"unit_price":  prod.UnitPrice,
						"quantity":    prod.Quantity,
						"total_price": prod.UnitPrice * float32(prod.Quantity),
					}).Error; err != nil {
					return err
				}
			}
		}
		return nil

	})
}

func (r *Repository) DeletePurchaseOrderByID(id string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("purchase_order_id = ?", id).Delete(&models.PurchaseProduct{}).Error; err != nil {
			return err
		}
		if err := tx.Where("id = ?", id).Delete(&models.PurchaseOrder{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// var order PurchaseOrderWorkshop

// err := db.Preload("Supplier").
//           Preload("PurchaseParts.PartWorkshop").
//           First(&order, "id = ?", someID).Error

// if err != nil {
//     // manejar error
// }
