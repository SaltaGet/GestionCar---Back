package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetElementByID(id string) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil

}

func (r *Repository) GetElementsByIdentifier(identifier string) (*[]models.Product, error) {
	var product []models.Product
	if err := r.DB.Where("identifier LIKE ?", "%"+identifier+"%").Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) GetAllElementsByName(name string) (*[]models.Product, error) {
	var products []models.Product
	if err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (r *Repository) GetAllElements() (*[]models.Product, error) {

	var products []models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil

}

func (r *Repository) CreateElement(element *models.ProductCreate) (string, error) {
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

func (r *Repository) UpdateElement(element *models.ProductUpdate) error {
	if err := r.DB.Model(&models.Product{}).Where("id = ?", element.ID).Updates(&models.Product{
		Identifier: element.Identifier,
		Name:       element.Name,
	}).Error; err != nil {
		return err
	}
	return nil

}

func (r *Repository) UpdateStock(stock int32, id string) error {
	if err := r.DB.Model(&models.Product{}).
		Where("id = ?", id).
		Update("stock", stock).Error; err != nil {
		return err
	}
	return nil

}

func (r *Repository) AddToStock(id string, cantidad int32) error {
	return r.DB.Model(&models.Product{}).
		Where("id = ?", id).
		UpdateColumn("stock", gorm.Expr("stock + ?", cantidad)).Error
}

func (r *Repository) SubtractFromStockToStock(id string, cantidad int32) error {
	return r.DB.Model(&models.Product{}).
		Where("id = ?", id).
		UpdateColumn("stock", gorm.Expr("stock - ?", cantidad)).Error
}

func (r *Repository) DeleteElement(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}
