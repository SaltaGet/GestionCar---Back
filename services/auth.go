package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"github.com/DanielChachagua/GestionCar/utils"
	"gorm.io/gorm"
)

func AuthLogin(username, password string) (string, error) {
	user, err := repositories.Repo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return "", models.ErrorResponse(500, "Error al  buscar usuario", err)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", models.ErrorResponse(401, "Credenciales incorrectas", nil)
	}

	token, err := utils.GenerateUserToken(user)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al generar token", err)
	}

	return token, nil
}

func AuthWorkplace(id string) (string, error) {
	workplace, err := repositories.Repo.GetWorkplaceByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", models.ErrorResponse(404, "Lugar de trabajo no encontrado", err)
		}
		return "", models.ErrorResponse(500, "Error al buscar lugar de trabajo", err)
	}

	token, err := utils.GenerateWorkplaceToken(workplace)

	if err != nil {
		return "", models.ErrorResponse(500, "Error al generar token", err)
	}

	return token, nil
}

func CurrentUser(userId string) (*models.User, error) {
	user, err := repositories.Repo.GetUserByID(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar usuario", err)
	}

	return user, nil
}

func CurrentWorkplace(workplaceId string) (*models.Workplace, error) {
	workplace, err := repositories.Repo.GetWorkplaceByID(workplaceId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Usuario no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al buscar usuario", err)
	}

	return workplace, nil
}

// func GetWorkplaceByRole(role string) (*models.Workplace, error) {
// 	workplace, err := repositories.Repo.GetWorkplaceByRole(role)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, models.ErrorResponse(404, "Rol no encontrado", err)
// 		}
// 		return nil, models.ErrorResponse(500, "Error al buscar rol", err)
// 	}

// 	return workplace, nil
// }