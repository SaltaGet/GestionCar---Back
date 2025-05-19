package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"gorm.io/gorm"
)

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
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByUsernameEmail(username string, email string) (bool, error) {
	err := r.DB.Where("email = ? OR username = ?", email, username).First(&models.User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *Repository) CreateUser(user *models.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}