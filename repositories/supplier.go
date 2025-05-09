package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/google/uuid"
)

func (r *Repository) GetSupplierByID(id string, workplace string) (*models.SupplierLaundry, *models.SupplierWorkshop, error) {
	if workplace == "laundry" {
		var supplier models.SupplierLaundry
		if err := r.DB.Where("id = ?", id).First(&supplier).Error; err != nil {
			return nil, nil, err
		}
		return &supplier, nil, nil
	} else if workplace == "workshop" {
		var supplier models.SupplierWorkshop
		if err := r.DB.Where("id = ?", id).First(&supplier).Error; err != nil {
			return nil, nil, err
		}
		return nil, &supplier, nil
	}
	return nil, nil, fmt.Errorf("tipo de espacio no soportado")
}

func (r *Repository) GetAllSuppliers(workplace string) ([]models.SupplierLaundry, []models.SupplierWorkshop, error) {
	if workplace == "laundry" {
		var suppliers []models.SupplierLaundry
		if err := r.DB.Find(&suppliers).Error; err != nil {
			return nil, nil, err
		}
		return suppliers, nil, nil
	}
	if workplace == "workshop" {
		var suppliers []models.SupplierWorkshop
		if err := r.DB.Find(&suppliers).Error; err != nil {
			return nil, nil, err
		}
		return nil, suppliers, nil
	}
	return nil, nil, fmt.Errorf("tipo de espacio no soportado")
}

func (r *Repository) CreateSupplier(supplierCreate *models.SupplierCreate, workplaceType string) (string, error) {
	var supplierID string
	switch workplaceType {
	case "laundry":
			supplier := models.SupplierLaundry{
					ID:      uuid.NewString(),
					Name:    supplierCreate.Name,
					Address: supplierCreate.Address,
					Phone:   supplierCreate.Phone,
					Email:   supplierCreate.Email,
			}
			if err := r.DB.Create(&supplier).Error; err != nil {
					return "", err
			}
			supplierID = supplier.ID
	case "workshop":
			supplier := models.SupplierWorkshop{
					ID:      uuid.NewString(),
					Name:    supplierCreate.Name,
					Address: supplierCreate.Address,
					Phone:   supplierCreate.Phone,
					Email:   supplierCreate.Email,
			}
			if err := r.DB.Create(&supplier).Error; err != nil {
					return "", err
			}
			supplierID = supplier.ID
	default:
			return "", errors.New("tipo de workplace no soportado")
	}

	return supplierID, nil
}

func (r *Repository) UpdateSupplier(id string, supplierUpdate *models.SupplierUpdate, workplace string) error {
	switch workplace {
	case "laundry":
		var supplierLaundry models.SupplierLaundry
		if err := r.DB.Where("id = ?", id).First(&supplierLaundry).Error; err != nil {
			return err
		}
		supplierLaundry.Name = supplierUpdate.Name
		supplierLaundry.Address = supplierUpdate.Address
		supplierLaundry.Phone = supplierUpdate.Phone
		supplierLaundry.Email = supplierUpdate.Email
		if err := r.DB.Save(&supplierLaundry).Error; err != nil {
			return err
		}
	case "workshop":
		var supplierWorkshop models.SupplierWorkshop
		if err := r.DB.Where("id = ?", id).First(&supplierWorkshop).Error; err != nil {
			return err
		}
		supplierWorkshop.Name = supplierUpdate.Name
		supplierWorkshop.Address = supplierUpdate.Address
		supplierWorkshop.Phone = supplierUpdate.Phone
		supplierWorkshop.Email = supplierUpdate.Email
		if err := r.DB.Save(&supplierWorkshop).Error; err != nil {
			return err
		}
	default:
			return errors.New("tipo de workplace no soportado")
	}
	return nil
}

func (r *Repository) DeleteSupplier(supplier interface{}) error {
	switch s := supplier.(type) {
	case *models.SupplierLaundry:
		if err := r.DB.Delete(s).Error; err != nil {
			return err
		}
	case *models.SupplierWorkshop:
		if err := r.DB.Delete(s).Error; err != nil {
			return err
		}
	default:
		return fmt.Errorf("tipo de espacio no soportado")
	}
	return nil
}

func (r *Repository) DeleteSupplierByID(id string, workplace string) error {
	if workplace == "laundry" {
		var supplier models.SupplierLaundry
		if err := r.DB.Where("id = ?", id).Delete(&supplier).Error; err != nil {
			return err
		}
	} else if workplace == "workshop" {
		var supplier models.SupplierWorkshop
		if err := r.DB.Where("id = ?", id).Delete(&supplier).Error; err != nil {
			return err
		}
	} else {
		return fmt.Errorf("tipo de espacio no soportado")
	}
	return nil
}

func (r *Repository) GetSupplierByName(name string, workplace string) (*models.SupplierLaundry, *models.SupplierWorkshop, error) {
	if workplace == "laundry" {
		var supplier models.SupplierLaundry
		if err := r.DB.Where("name LIKE ?", "%"+name +"%").First(&supplier).Error; err != nil {
			return nil, nil, err
		}
		return &supplier, nil, nil
	} else if workplace == "workshop" {
		var supplier models.SupplierWorkshop
		if err := r.DB.Where("name LIKE ?", "%"+name +"%").First(&supplier).Error; err != nil {
			return nil, nil, err
		}
		return nil, &supplier, nil
	}
	return nil, nil, fmt.Errorf("tipo de espacio no soportado")
}

