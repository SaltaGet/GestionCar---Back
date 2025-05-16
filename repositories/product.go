package repositories

import (
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetElementByID(id string, workplace string) (*models.ProductLaundry, *models.PartWorkshop, error) {
	if workplace == "laundry" {
		var product models.ProductLaundry
		if err := r.DB.Where("id = ?", id).First(&product).Error; err != nil {
			return nil, nil, err
		}
		return &product, nil, nil
	} else if workplace == "workshop" {
		var part models.PartWorkshop
		if err := r.DB.Where("id = ?", id).First(&part).Error; err != nil {
			return nil, nil, err
		}
		return nil, &part, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetElementsByIdentifier(identifier string, workplace string) (*[]models.ProductLaundry, *[]models.PartWorkshop, error) {
	if workplace == "laundry" {
		var product []models.ProductLaundry
		if err := r.DB.Where("identifier LIKE ?", "%"+identifier+"%").Find(&product).Error; err != nil {
			return nil, nil, err
		}
		return &product, nil, nil
	} else if workplace == "workshop" {
		var part []models.PartWorkshop
		if err := r.DB.Where("identifier LIKE ?", "%"+identifier+"%").Find(&part).Error; err != nil {
			return nil, nil, err
		}
		return nil, &part, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetAllElementsByName(name string, workplace string) (*[]models.ProductLaundry, *[]models.PartWorkshop, error) {
	if workplace == "laundry" {
		var products []models.ProductLaundry
		if err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
			return nil, nil, err
		}
		return &products, nil, nil
	} else if workplace == "workshop" {
		var parts []models.PartWorkshop
		if err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&parts).Error; err != nil {
			return nil, nil, err
		}
		return nil, &parts, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) GetAllElements(workplace string) (*[]models.ProductLaundry, *[]models.PartWorkshop, error) {
	if workplace == "laundry" {
		var products []models.ProductLaundry
		if err := r.DB.Find(&products).Error; err != nil {
			return nil, nil, err
		}
		return &products, nil, nil
	} else if workplace == "workshop" {
		var parts []models.PartWorkshop
		if err := r.DB.Find(&parts).Error; err != nil {
			return nil, nil, err
		}
		return nil, &parts, nil
	}
	return nil, nil, fmt.Errorf("tipo de movimiento no soportado")
}

func (r *Repository) CreateElement(element *models.ProductCreate, workplace string) (string, error) {
	newID := uuid.NewString()
	switch workplace {
	case "laundry":
		if err := r.DB.Create(&models.ProductLaundry{
			ID:         newID,
			Identifier: element.Identifier,
			Name:       element.Name,
			Stock: 0,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	case "workshop":
		if err := r.DB.Create(&models.PartWorkshop{
			ID:         newID,
			Identifier: element.Identifier,
			Name:       element.Name,
		}).Error; err != nil {
			return "", err
		}
		return newID, nil
	default:
		return "", fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdateElement(element *models.ProductUpdate, workplace string) error {
	switch workplace {
	case "laundry":
		if err := r.DB.Model(&models.ProductLaundry{}).Where("id = ?", element.ID).Updates(&models.ProductLaundry{
			Identifier: element.Identifier,
			Name:       element.Name,
		}).Error; err != nil {
			return err
		}
		return nil
	case "workshop":
		if err := r.DB.Model(&models.PartWorkshop{}).Where("id = ?", element.ID).Updates(&models.PartWorkshop{
			Identifier: element.Identifier,
			Name:       element.Name,
		}).Error; err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) UpdateStock(stock int32, id string, workplace string) error {
	switch workplace {
	case "laundry":
		if err := r.DB.Model(&models.ProductLaundry{}).
			Where("id = ?", id).
			Update("stock", stock).Error; err != nil {
			return err
		}
		return nil
	case "workshop":
		if err := r.DB.Model(&models.PartWorkshop{}).
			Where("id = ?", id).
			Update("stock", stock).Error; err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) AddToStock(id string, cantidad int32, workplace string) error {
	switch workplace {
	case "laundry":
			return r.DB.Model(&models.ProductLaundry{}).
					Where("id = ?", id).
					UpdateColumn("stock", gorm.Expr("stock + ?", cantidad)).Error
	case "workshop":
			return r.DB.Model(&models.PartWorkshop{}).
					Where("id = ?", id).
					UpdateColumn("stock", gorm.Expr("stock + ?", cantidad)).Error
	default:
			return fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) SubtractFromStockToStock(id string, cantidad int32, workplace string) error {
	switch workplace {
	case "laundry":
			return r.DB.Model(&models.ProductLaundry{}).
					Where("id = ?", id).
					UpdateColumn("stock", gorm.Expr("stock - ?", cantidad)).Error
	case "workshop":
			return r.DB.Model(&models.PartWorkshop{}).
					Where("id = ?", id).
					UpdateColumn("stock", gorm.Expr("stock - ?", cantidad)).Error
	default:
			return fmt.Errorf("tipo de movimiento no soportado")
	}
}

func (r *Repository) DeleteElement(id string, workplace string) error {
	switch workplace {
	case "laundry":
		if err := r.DB.Where("id = ?", id).Delete(&models.ProductLaundry{}).Error; err != nil {
			return err
		}
	case "workshop":
		if err := r.DB.Where("id = ?", id).Delete(&models.PartWorkshop{}).Error; err != nil {
			return err
		}
	default:
		return fmt.Errorf("tipo de movimiento no soportado")
	}
	return nil
}
