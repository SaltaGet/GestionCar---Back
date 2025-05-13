package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (r *Repository) GetAllPurchaseOrders(workplace string) (*[]models.PurchaseOrderLaundry, *[]models.PurchaseOrderWorkshop, error) {
	if workplace == "laundry" {
		var purchaseOrders []models.PurchaseOrderLaundry
		if err := r.DB.Find(&purchaseOrders).Error; err != nil {
			return nil, nil, err
		}
		return &purchaseOrders, nil, nil
	} else if workplace == "workshop" {
		var purchaseOrders []models.PurchaseOrderWorkshop
		if err := r.DB.Find(&purchaseOrders).Error; err != nil {
			return nil, nil, err
		}
		return nil, &purchaseOrders, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) CreatePurchaseOrder(purchaseOrder *models.PurchaseOrderCreate, workplace string) (string, error) {
	newID := uuid.NewString()
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		switch workplace {
		case "laundry":
			if err := tx.Create(&models.PurchaseOrderLaundry{
				ID:          newID,
				OrderNumber: purchaseOrder.OrderNumber,
				OrderDate:   purchaseOrder.OrderDate,
				Amount:      purchaseOrder.Amount,
				SupplierID:  purchaseOrder.SupplierID,
			}).Error; err != nil {
				return err
			}
			for _, element := range purchaseOrder.PurchaseProductCreates {
				if err := tx.Create(&models.PurchaseProductLaundry{
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
		case "workshop":
			if err := tx.Create(&models.PurchaseOrderWorkshop{
				ID:          newID,
				OrderNumber: purchaseOrder.OrderNumber,
				OrderDate:   purchaseOrder.OrderDate,
				Amount:      purchaseOrder.Amount,
				SupplierID:  purchaseOrder.SupplierID,
			}).Error; err != nil {
				return err
			}
			for _, element := range purchaseOrder.PurchaseProductCreates {
				if err := tx.Create(&models.PurchasePartWorkshop{
					ID:              uuid.NewString(),
					PartID:          element.ProductID,
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
		default:
			return fmt.Errorf("tipo de movimiento no soportado")
		}
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

func (r *Repository) UpdatePurchaseOrder(purchaseOrder *models.PurchaseOrderUpdate, workplace string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		switch workplace {
		case "laundry":
			if err := tx.Where("id = ?", purchaseOrder.ID).Updates(&models.PurchaseOrderLaundry{
				OrderNumber: purchaseOrder.OrderNumber,
				OrderDate:   purchaseOrder.OrderDate,
				Amount:      purchaseOrder.Amount,
				SupplierID:  purchaseOrder.SupplierID,
			}).Error; err != nil {
				return err
			}

			var existingProducts []models.PurchaseProductLaundry
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
					if err := tx.Delete(&models.PurchaseProductLaundry{}, "id = ?", p.ID).Error; err != nil {
						return err
					}
				}
			}

			for _, prod := range purchaseOrder.PurchaseProductUpdates {
				if prod.ID == "" || !existingIDs[prod.ID] {
					newProd := models.PurchaseProductLaundry{
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
					if err := tx.Model(&models.PurchaseProductLaundry{}).
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
		case "workshop":
			if err := tx.Where("id = ?", purchaseOrder.ID).Updates(&models.PurchaseOrderWorkshop{
				OrderNumber: purchaseOrder.OrderNumber,
				OrderDate:   purchaseOrder.OrderDate,
				Amount:      purchaseOrder.Amount,
				SupplierID:  purchaseOrder.SupplierID,
			}).Error; err != nil {
				return err
			}

			var existingProducts []models.PurchasePartWorkshop
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
					if err := tx.Delete(&models.PurchasePartWorkshop{}, "id = ?", p.ID).Error; err != nil {
						return err
					}
				}
			}

			for _, prod := range purchaseOrder.PurchaseProductUpdates {
				if prod.ID == "" || !existingIDs[prod.ID] {
					newProd := models.PurchasePartWorkshop{
						ID:              uuid.NewString(),
						PartID:       prod.ProductID,
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
					if err := tx.Model(&models.PurchasePartWorkshop{}).
						Where("id = ?", prod.ID).
						Updates(map[string]interface{}{
							"part_id":  prod.ProductID,
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
		default:
			return fmt.Errorf("tipo de movimiento no soportado")
		}
	})
}

func (r *Repository) DeletePurchaseOrderByID(id string, workplace string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		switch workplace {
		case "laundry":
			if err := tx.Where("purchase_order_id = ?", id).Delete(&models.PurchaseProductLaundry{}).Error; err != nil {
				return err
			}
			if err := tx.Where("id = ?", id).Delete(&models.PurchaseOrderLaundry{}).Error; err != nil {
				return err
			}
		case "workshop":
			if err := tx.Where("purchase_order_id = ?", id).Delete(&models.PurchasePartWorkshop{}).Error; err != nil {
				return err
			}
			if err := tx.Where("id = ?", id).Delete(&models.PurchaseOrderWorkshop{}).Error; err != nil {
				return err
			}
		default:
			return fmt.Errorf("tipo de espacio no soportado")
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
