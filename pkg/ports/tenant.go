package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type TenantRepository interface {
	TenantGetByID(id string) (tenant *models.Tenant, err error) 
}

type TenantService interface {
	TenantGetByID(id string) (name string, err error) 
}