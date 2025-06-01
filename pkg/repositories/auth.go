package repositories

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
	"gorm.io/gorm"
)

func (r *Repository) AuthLogin(username string, password string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, models.ErrorResponse(401, "Credenciales incorrectas", nil)
	}

	return &user, nil
}