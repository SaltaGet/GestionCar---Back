package services

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func (r *RoleService) RoleGetAll(role string) (*[]models.Role, error) {
	roles, err := r.RoleRepository.RoleGetAll(role)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los roles", err)
	}
	return roles, nil
}