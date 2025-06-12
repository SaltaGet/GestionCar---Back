package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *TenantRepository) ProductGetByID(id string) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Producto no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error interno al buscar producto", err)
	}
	return &product, nil

}

func (r *TenantRepository) ProductGetByIdentifier(identifier string) (*[]models.Product, error) {
	var product []models.Product
	if err := r.DB.Where("identifier LIKE ?", "%"+identifier+"%").Find(&product).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar producto", err)
	}
	return &product, nil
}

func (r *TenantRepository) ProductGetByName(name string) (*[]models.Product, error) {
	var products []models.Product
	if err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar productos", err)
	}
	return &products, nil
}

func (r *TenantRepository) ProductGetAll() (*[]models.Product, error) {
	var products []models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, models.ErrorResponse(500, "Error interno al buscar productos", err)
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
		return "", models.ErrorResponse(500, "Error interno al crear producto", err)
	}
	return newID, nil

}

func (r *TenantRepository) ProductUpdate(element *models.ProductUpdate) error {
	if err := r.DB.Model(&models.Product{}).Where("id = ?", element.ID).Updates(&models.Product{
		Identifier: element.Identifier,
		Name:       element.Name,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Producto no encontrado", err)
		}
		return models.ErrorResponse(500, "Error interno al actualizar producto", err)
	}
	return nil

}

func (r *TenantRepository) UpdateStock(stockUpdate *models.StockUpdate) error {
	if err := r.DB.Model(&models.Product{}).
		Where("id = ?", stockUpdate.ID).
		Update("stock", stockUpdate.Stock).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Producto no encontrado", err)
		}
		return models.ErrorResponse(500, "Error interno al actualizar stock", err)
	}
	return nil

}

func (r *TenantRepository) AddToStock(stockUpdate *models.StockUpdate) error {
	if err := r.DB.Model(&models.Product{}).
		Where("id = ?", stockUpdate.ID).
		UpdateColumn("stock", gorm.Expr("stock + ?", stockUpdate.Stock)).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.ErrorResponse(404, "Producto no encontrado", err)
			}
			return models.ErrorResponse(500, "Error interno al actualizar stock", err)
		}

	return nil
}

func (r *TenantRepository) SubtractFromStockToStock(stockUpdate *models.StockUpdate) error {
	if err := r.DB.Model(&models.Product{}).
		Where("id = ?", stockUpdate.ID).
		UpdateColumn("stock", gorm.Expr("stock - ?", stockUpdate.Stock)).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.ErrorResponse(404, "Producto no encontrado", err)
			}
			return models.ErrorResponse(500, "Error interno al actualizar stock", err)
		}

	return nil
}

func (r *TenantRepository) ProductDelete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Producto no encontrado", err)
			
		}
		return models.ErrorResponse(500, "Error interno al eliminar producto", err)
	}
	return nil
}
