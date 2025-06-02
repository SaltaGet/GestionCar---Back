package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
	"github.com/google/uuid"
)

func (t *TenantService) TenantGetByID(id string) (string, error) {
	tenant, err := t.TenantRepository.TenantGetByID(id)
	if err != nil {
		return "", err
	}
	return tenant.Name, nil
}

func (t *TenantService) TenantGetAll(userID string) (*[]models.TenantResponse, error) {
	tenants, err := t.TenantRepository.TenantGetAll(userID)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

func (t *TenantService) TenantCreate(tenantCreate *models.TenantCreate) (string, error) {
	tenantName := strings.ReplaceAll(tenantCreate.Name, " ", "_")
	uri := fmt.Sprintf("%s%s_%s.db%s",os.Getenv("URI_PATH"),tenantName,uuid.NewString(),os.Getenv("URI_CONFIG"))
	connection, err := utils.Encrypt(uri)
	if err != nil {
		return "", err
	}

	id, err := t.TenantRepository.TenantCreate(&models.Tenant{
		ID:        uuid.NewString(),
		Name:      tenantCreate.Name,
		Address:   tenantCreate.Address,
		Phone:     tenantCreate.Phone,
		Email:     tenantCreate.Email,
		CUITPDV:   tenantCreate.CUITPDV,
		Connection: connection,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t *TenantService) TenantUpdate(userID string, tenant *models.TenantUpdate) error {
	return t.TenantRepository.TenantUpdate(userID, tenant)
}
