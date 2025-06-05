package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
)

func (r *TenantRepository) PurchaseProdcutGetByID(id string) (*models.PurchaseProduct, error) {
	var purchaseProduct models.PurchaseProduct
	if err := r.DB.Where("id = ?", id).First(&purchaseProduct).Error; err != nil {
		return nil, err
	}
	return &purchaseProduct, nil
}

func (r *TenantRepository) PurchaseProductGetByPurchaseID(purchaseID string) (*[]models.PurchaseProduct, error) {
	var purchaseProduct []models.PurchaseProduct
	if err := r.DB.Where("purchase_order_id = ?", purchaseID).Find(&purchaseProduct).Error; err != nil {
		return nil, err
	}
	return &purchaseProduct, nil
}

func (r *TenantRepository) PurchaseProductGetAll() ([]models.PurchaseProduct, error) {
	var purchaseProducts []models.PurchaseProduct
	if err := r.DB.Find(&purchaseProducts).Error; err != nil {
		return nil, err
	}
	return purchaseProducts, nil
}

func (r *TenantRepository) PurchaseProductCreate(element *models.PurchaseProductCreate) (string, error) {
	newID := uuid.NewString()
	if err := r.DB.Create(&models.PurchaseProduct{
		ID:         newID,
		ProductID:  element.ProductID,
		ExpiredAt:  element.ExpiredAt,
		UnitPrice:  element.UnitPrice,
		Quantity:   element.Quantity,
		TotalPrice: element.UnitPrice * float32(element.Quantity),
	}).Error; err != nil {
		return "", err
	}
	return newID, nil

}

func (r *TenantRepository) PurchaseProductUpdate(element *models.PurchaseProductUpdate) error {
	if err := r.DB.Where("id = ?", element.ID).Updates(&models.PurchaseProduct{
		ProductID:  element.ProductID,
		ExpiredAt:  element.ExpiredAt,
		UnitPrice:  element.UnitPrice,
		Quantity:   element.Quantity,
		TotalPrice: element.UnitPrice * float32(element.Quantity),
	}).Error; err != nil {
		return err
	}
	return nil

}

func (r *TenantRepository) PurchaseProductDelete(id string) error {
	var purchaseProduct models.PurchaseProduct
	if err := r.DB.Where("id = ?", id).Delete(&purchaseProduct).Error; err != nil {
		return err
	}

	return nil
}
