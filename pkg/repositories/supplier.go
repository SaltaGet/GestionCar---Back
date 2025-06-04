package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
)

func (r *TenantRepository) SupplierGetByID(id string) (*models.Supplier, error) {
		var supplier models.Supplier
		if err := r.DB.Where("id = ?", id).First(&supplier).Error; err != nil {
			return nil, err
		}
		return &supplier, nil
}

func (r *TenantRepository) SupplierGetAll() ([]models.Supplier, error) {
		var suppliers []models.Supplier
		if err := r.DB.Find(&suppliers).Error; err != nil {
			return nil, err
		}
		return suppliers, nil
}

func (r *TenantRepository) SupplierGetByName(name string) (*[]models.Supplier, error) {
		var supplier []models.Supplier
		if err := r.DB.Where("name LIKE ?", "%"+name +"%").Find(&supplier).Error; err != nil {
			return nil, err
		}
		return &supplier, nil
}

func (r *TenantRepository) SupplierCreate(supplierCreate *models.SupplierCreate) (string, error) {
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

func (r *TenantRepository) SupplierUpdate(supplierUpdate *models.SupplierUpdate) error {
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

func (r *TenantRepository) SupplierDelete(id string) error {
		var supplier models.Supplier
		if err := r.DB.Where("id = ?", id).Delete(&supplier).Error; err != nil {
			return err
		}
	return nil
}
