package services

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
)

func GetRoleAll(role string,workplace string) (*[]models.Role, error) {
	roles, err := repositories.Repo.GetAllRoles(role, workplace)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los roles", err)
	}
	return roles, nil
}