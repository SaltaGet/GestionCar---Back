package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
)

func (r *Repository) GetSupplierByID(id string) (*models.Supplier, error) {
		var supplier models.Supplier
		if err := r.DB.Where("id = ?", id).First(&supplier).Error; err != nil {
			return nil, err
		}
		return &supplier, nil
}

func (r *Repository) GetAllSuppliers() ([]models.Supplier, error) {
		var suppliers []models.Supplier
		if err := r.DB.Find(&suppliers).Error; err != nil {
			return nil, err
		}
		return suppliers, nil
}

func (r *Repository) CreateSupplier(supplierCreate *models.SupplierCreate) (string, error) {
	var supplierID string
			supplier := models.Supplier{
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
	return supplierID, nil
}

func (r *Repository) UpdateSupplier(supplierUpdate *models.SupplierUpdate) error {
		var supplierLaundry models.Supplier
		if err := r.DB.Where("id = ?", supplierUpdate.ID).First(&supplierLaundry).Error; err != nil {
			return err
		}
		supplierLaundry.Name = supplierUpdate.Name
		supplierLaundry.Address = supplierUpdate.Address
		supplierLaundry.Phone = supplierUpdate.Phone
		supplierLaundry.Email = supplierUpdate.Email
		if err := r.DB.Save(&supplierLaundry).Error; err != nil {
			return err
		}
	return nil
}

func (r *Repository) DeleteSupplierByID(id string) error {
		var supplier models.Supplier
		if err := r.DB.Where("id = ?", id).Delete(&supplier).Error; err != nil {
			return err
		}
	return nil
}

func (r *Repository) GetSupplierByName(name string) (*[]models.Supplier, error) {
		var supplier []models.Supplier
		if err := r.DB.Where("name LIKE ?", "%"+name +"%").Find(&supplier).Error; err != nil {
			return nil, err
		}
		return &supplier, nil
}

