package services

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func (r *RoleService) RoleGetAll() (*[]models.Role, error) {
	roles, err := r.RoleRepository.RoleGetAll()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los roles", err)
	}
	return roles, nil
}

func (r *RoleService) RoleCreate(roleCrate *models.RoleCreate) (id string, err error) {
	id, err = r.RoleRepository.RoleCreate(roleCrate)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear el role", err)
	}
	return id, nil
}