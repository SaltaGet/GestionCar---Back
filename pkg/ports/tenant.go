package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type TenantRepository interface {
	TenantGetByID(id string) (tenant *models.Tenant, err error) 
	TenantGetAll(userID string) (tenants *[]models.TenantResponse, err error)
	TenantCreate(tenant *models.Tenant) (id string, err error)
	TenantUpdate(userID string, tenant *models.TenantUpdate) (err error)
}

type TenantService interface {
	TenantGetByID(id string) (name string, err error) 
	TenantGetAll(userID string) (tenants *[]models.TenantResponse, err error)
	TenantCreate(tenant *models.TenantCreate) (id string, err error)
	TenantUpdate(userID string, tenant *models.TenantUpdate) (err error)
}