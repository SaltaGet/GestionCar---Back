package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) UserGetByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "id = ?", id).Error
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UserGetByUsername(username string) (*models.User, error) {
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

func (r *Repository) UserGetExistByUsernameEmail(username string, email string) (bool, error) {
	err := r.DB.Where("email = ? OR username = ?", email, username).First(&models.User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *Repository) UserCreate(user *models.UserCreate) (string, error) {
	newID := uuid.NewString()
	err := r.DB.Create(&models.User{Username: user.Username, Email: user.Email, Password: user.Password}).Error
	if err != nil {
		return "" ,err
	}
	return newID, nil
}