package services

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
)

func GetWorkplaceAll(role string) (*[]models.Workplace, error) {
	workplaces, err := repositories.Repo.GetWorkplaceAll(role)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los lugares de trabajo", err)
	}
	return workplaces, nil
}