package repositories

import "github.com/DanielChachagua/GestionCar/models"

func (r *Repository) GetVehicleByID(id string) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	if err := r.DB.Where("id = ?", id).First(&vehicle).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *Repository) GetVehicleByDomain(domain string) (*[]models.Vehicle, error) {
	var vehicles []models.Vehicle
	if err := r.DB.Preload("Client").Where("domain LIKE ?", "%"+domain+"%").Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return &vehicles, nil
}

func (r *Repository) CreateVehicle(vehicle *models.Vehicle) (string, error) {
	if err := r.DB.Create(vehicle).Error; err != nil {
		return "", err
	}
	return vehicle.ID, nil
}

func (r *Repository) UpdateVehicle(vehicle *models.Vehicle) error {
	var existing models.Vehicle
	if err := r.DB.First(&existing, "id = ?", vehicle.ID).Error; err != nil {
		return err 
	}

	if err := r.DB.Updates(vehicle).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteVehicle(id string) error {
	var vehicle models.Vehicle
	if err := r.DB.Where("id = ?", id).Delete(&vehicle).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	if err := r.DB.Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (r *Repository) GetVehicleByClientID(clientID string) ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	if err := r.DB.Where("client_id = ?", clientID).Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}

// func (r *Repository) GetVehicleByClientIDAndDomain(clientID, domain string) (*models.Vehicle, error) {
// 	var vehicle models.Vehicle
// 	if err := r.DB.Where("client_id = ? AND domain = ?", clientID, domain).First(&vehicle).Error; err != nil {
// 		return nil, err
// 	}
// 	return &vehicle, nil
// }

// func (r *Repository) GetVehicleByClientIDAndDomainLike(clientID, domain string) ([]models.Vehicle, error) {
// 	var vehicles []models.Vehicle
// 	if err := r.DB.Where("client_id = ? AND domain LIKE ?", clientID, "%"+domain+"%").Find(&vehicles).Error; err != nil {
// 		return nil, err
// 	}
// 	return vehicles, nil
// }
