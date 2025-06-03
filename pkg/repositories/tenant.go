package repositories

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *MainRepository) TenantGetByID(userID, tenantID string) (*models.Tenant, error) {
	var tenant models.Tenant
	err := r.DB.
		Preload("UserTenants", "user_id = ? AND tenant_id = ?", userID, tenantID).
		Where("id = ?", tenantID).
		First(&tenant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Tenant not found", err)
		}
		return nil, models.ErrorResponse(500, "Error retrieving tenant", err)
	}

	if !tenant.IsActive {
		return nil, models.ErrorResponse(403, "Tenant is inactive", nil)
	}

	if len(tenant.UserTenants) == 0 {
		return nil, models.ErrorResponse(403, "You do not have permission to access this tenant", nil)
	}

	if !tenant.UserTenants[0].IsActive {
		return nil, models.ErrorResponse(403, "You do not have permission to access this tenant", nil)

	}

	return &tenant, nil
}

func (r *MainRepository) TenantGetAll(userID string) (*[]models.TenantResponse, error) {
	var tenants []models.TenantResponse

	err := r.DB.
		Model(&models.Tenant{}).
		Select(`tenants.id, tenants.name, tenants.address, tenants.phone,
                tenants.email, tenants.is_active,
                user_tenants.is_active AS user_is_active,
                tenants.created_at, tenants.updated_at`).
		Joins("JOIN user_tenants ON tenants.id = user_tenants.tenant_id").
		Where("user_tenants.user_id = ?", userID).
		Scan(&tenants).Error

	if err != nil {
		return nil, models.ErrorResponse(500, "Error retrieving tenants", err)
	}

	return &tenants, nil
}

func (r *MainRepository) TenantUserCreate(tenantUserCreate *models.TenantUserCreate) (string, error) {
	tx := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tenantName := strings.ReplaceAll(tenantUserCreate.TenantCreate.Name, " ", "_")
	uri := fmt.Sprintf("%s%s_%s.db%s", os.Getenv("URI_PATH"), tenantName, uuid.NewString(), os.Getenv("URI_CONFIG"))
	connection, err := utils.Encrypt(uri)
	if err != nil {
		return "", err
	}

	tenant := &models.Tenant{
		ID:         uuid.NewString(),
		Name:       tenantUserCreate.TenantCreate.Name,
		Address:    tenantUserCreate.TenantCreate.Address,
		Phone:      tenantUserCreate.TenantCreate.Phone,
		Email:      tenantUserCreate.TenantCreate.Email,
		CuitPdv:    tenantUserCreate.TenantCreate.CuitPdv,
		Connection: connection,
	}

	if err := tx.Create(tenant).Error; err != nil {
		tx.Rollback()
		return "", models.ErrorResponse(500, "Error creating tenant", err)
	}

	pass, err :=utils.HashPassword(tenantUserCreate.UserCreate.Password)
	if err != nil {
		return "", models.ErrorResponse(500, "Error hashed password", err)
	}

	user := &models.User{
		ID:       uuid.NewString(),
		FirstName: tenantUserCreate.UserCreate.FirstName,
		LastName:  tenantUserCreate.UserCreate.LastName,
		Username: tenantUserCreate.UserCreate.Username,
		Email:    tenantUserCreate.UserCreate.Email,
		Password: pass,
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Create(&models.UserTenant{
		UserID:    user.ID,
		TenantID:  tenant.ID,
		IsAdmin:   true,
	}).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	err = database.PrepareDB(uri, user.ID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return tenant.ID, nil
}

func (r *MainRepository) TenantUpdate(userID string, tenant *models.TenantUpdate) error {
	var userTenant models.UserTenant

	err := r.DB.First(&userTenant, "user_id = ? AND tenant_id = ?", userID, tenant.ID).Error
	if err != nil {
		return models.ErrorResponse(404, "User-tenant relationship not found", err)
	}

	if !userTenant.IsAdmin {
		return models.ErrorResponse(403, "You do not have permission to update the tenant", nil)
	}

	if err := r.DB.Model(&models.Tenant{}).Updates(tenant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Tenant not found", err)
		}
		return models.ErrorResponse(500, "Error updating tenant", err)
	}

	return nil
}
