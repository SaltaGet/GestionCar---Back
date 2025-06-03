package services

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
)

func (u *UserService) UserCreate(userCreate *models.UserCreate) (string, error) {
	existingUser, err := u.UserRepository.UserGetExistByUsernameEmail(userCreate.Username, userCreate.Email)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al buscar el usuario", err)
	}
	if existingUser {
		return "", models.ErrorResponse(400, "El username o el email ya existe", nil)
	}

	
	userCreate.Password, err = utils.HashPassword(userCreate.Password)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al hashear la contraseña", err)
	}

	id, err := u.UserRepository.UserCreate(userCreate)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear el usuario", err)
	}

	return id, nil
}