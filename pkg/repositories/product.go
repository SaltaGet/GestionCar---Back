package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *TenantRepository) ProductGetByID(id string) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil

}

func (r *TenantRepository) ProductGetByIdentifier(identifier string) (*[]models.Product, error) {
	var product []models.Product
	if err := r.DB.Where("identifier LIKE ?", "%"+identifier+"%").Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *TenantRepository) ProductGetByName(name string) (*[]models.Product, error) {
	var products []models.Product
	if err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (r *TenantRepository) ProductGetAll() (*[]models.Product, error) {

	var products []models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil

}

func (r *TenantRepository) ProductCreate(element *models.ProductCreate) (string, error) {
	newID := uuid.NewString()
	if err := r.DB.Create(&models.Product{
		ID:         newID,
		Identifier: element.Identifier,
		Name:       element.Name,
		Stock:      0,
	}).Error; err != nil {
		return "", err
	}
	return newID, nil

}

func (r *TenantRepository) ProductUpdate(element *models.ProductUpdate) error {
	if err := r.DB.Model(&models.Product{}).Where("id = ?", element.ID).Updates(&models.Product{
		Identifier: element.Identifier,
		Name:       element.Name,
	}).Error; err != nil {
		return err
	}
	return nil

}

func (r *TenantRepository) UpdateStock(stockUpdate *models.StockUpdate) error {
	if err := r.DB.Model(&models.Product{}).
		Where("id = ?", stockUpdate.ID).
		Update("stock", stockUpdate.Stock).Error; err != nil {
		return err
	}
	return nil

}

func (r *TenantRepository) AddToStock(stockUpdate *models.StockUpdate) error {
	return r.DB.Model(&models.Product{}).
		Where("id = ?", stockUpdate.ID).
		UpdateColumn("stock", gorm.Expr("stock + ?", stockUpdate.Stock)).Error
}

func (r *TenantRepository) SubtractFromStockToStock(stockUpdate *models.StockUpdate) error {
	return r.DB.Model(&models.Product{}).
		Where("id = ?", stockUpdate.ID).
		UpdateColumn("stock", gorm.Expr("stock - ?", stockUpdate.Stock)).Error
}

func (r *TenantRepository) ProductDelete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}
