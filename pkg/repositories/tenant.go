package repositories

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func (r *Repository) TenantGetByID(id string) (*models.Tenant, error) {
	var tenant models.Tenant
		err := r.DB.First(&tenant, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

// func (r *Repository) GetWorkplaceLaundry() (*models.Workplace, error) {
// 	var workplace models.Workplace
// 		err := r.DB.First(&workplace, "identifier = ?", "laundry").Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &workplace, nil
// }

// func (r *Repository) GetWorkplaceWorkshop() (*models.Workplace, error) {
// 	var workplace models.Workplace
// 		err := r.DB.First(&workplace, "identifier = ?", "workshop").Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &workplace, nil
// }

// func (r *Repository) GetWorkplaceAll(role string) (*[]models.Workplace, error) {
// 	var workplaces []models.Workplace
// 	if utils.Contains([]string{"admin_laundry", "employee_laundry"}, role) {
// 		err := r.DB.Where("identifier = ?", "laundry").Find(&workplaces).Error
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &workplaces, nil
// 	} else if utils.Contains([]string{"admin_workshop", "employee_workshop"}, role) {
// 		err := r.DB.Where("identifier = ?", "workshop").Find(&workplaces).Error
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &workplaces, nil
// 	} else {
// 		err := r.DB.Find(&workplaces).Error
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &workplaces, nil
// 	}
// }
