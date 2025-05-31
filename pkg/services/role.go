package services

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func (r *RoleService) GetRoleAll(role string) (*[]models.Role, error) {
	roles, err := r.RoleRepository.GetRoleAll(role)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al buscar los roles", err)
	}
	return roles, nil
}