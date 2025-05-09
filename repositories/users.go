package repositories

import "github.com/DanielChachagua/GestionCar/models"

func (r *Repository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "id = ?", id).Error
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "username = ?", username).Error
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(user *models.User) error {
	err := r.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}