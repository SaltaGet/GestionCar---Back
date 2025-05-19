package services

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"github.com/DanielChachagua/GestionCar/utils"
	"github.com/google/uuid"
)

func UserCreate(user *models.UserCreate) (string, error) {
	// Check if the user already exists
	existingUser, err := repositories.Repo.GetUserByUsernameEmail(user.Username, user.Email)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al buscar el usuario", err)
	}
	if existingUser {
		return "", models.ErrorResponse(400, "El username o el email ya existe", nil)
	}

	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al hashear la contraseña", err)
	}
	// Create the new user
	newUser := &models.User{
		ID: uuid.NewString(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Username: user.Username,
		Password: pass,
	}
	err = repositories.Repo.CreateUser(newUser)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear el usuario", err)
	}

	return newUser.ID, nil
}