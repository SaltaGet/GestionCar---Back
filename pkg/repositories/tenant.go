package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"gorm.io/gorm"
)

func (r *Repository) TenantGetByID(id string) (*models.Tenant, error) {
	var tenant models.Tenant
	err := r.DB.First(&tenant, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Tenant not found", err)
		}
		return nil, models.ErrorResponse(500, "Error retrieving tenant", err)
	}
	return &tenant, nil
}

func (r *Repository) TenantGetAll(userID string) (*[]models.TenantResponse, error) {
	var tenants []models.TenantResponse

	err := r.DB.
		Model(&models.Tenant{}).
		Select(`tenants.id, tenants.name, tenants.address, tenants.phone,
                tenants.email, tenants.identifier, tenants.is_active,
                user_tenants.active AS user_is_active,
                tenants.created_at, tenants.updated_at`).
		Joins("JOIN user_tenants ON tenants.id = user_tenants.tenant_id").
		Where("user_tenants.user_id = ?", userID).
		Scan(&tenants).Error

	if err != nil {
		return nil, models.ErrorResponse(500, "Error retrieving tenants", err)
	}

	return &tenants, nil
}

func (r *Repository) TenantCreate(tenant *models.Tenant) (string, error) {
	if err := r.DB.Create(&tenant).Error; err != nil {
		return "", models.ErrorResponse(500, "Error creating tenant", err)
	}

	return tenant.ID, nil
}

func (r *Repository) TenantUpdate(userID string, tenant *models.TenantUpdate) error {
	var userTenant models.UserTenant

	err := r.DB.First(&userTenant, "user_id = ? AND tenant_id = ?", userID, tenant.ID).Error
	if err != nil {
		return models.ErrorResponse(404, "User-tenant relationship not found", err)
	}

	if userTenant.IsAdmin {
		return models.ErrorResponse(403, "You do not have permission to update the tenant", nil)
	}

	if err := r.DB.Model(&userTenant).Updates(tenant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Tenant not found", err)
		}
		return models.ErrorResponse(500, "Error updating tenant", err)
	}

	return nil
}

