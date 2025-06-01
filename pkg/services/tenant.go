package services

// import (
// 	"github.com/DanielChachagua/GestionCar/pkg/models"
// 	"github.com/DanielChachagua/GestionCar/pkg/repositories"
// )

func (t *TenantService) TenantGetByID(id string) (string, error) {
	tenant, err := t.TenantRepository.TenantGetByID(id)
	if err != nil {
		return "", err
	}
	return tenant.Name, nil
}

// func GetWorkplaceAll(role string) (*[]models.Workplace, error) {
// 	workplaces, err := repositories.Repo.GetWorkplaceAll(role)
// 	if err != nil {
// 		return nil, models.ErrorResponse(500, "Error al buscar los lugares de trabajo", err)
// 	}
// 	return workplaces, nil
// }