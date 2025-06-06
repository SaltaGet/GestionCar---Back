package services

import "github.com/DanielChachagua/GestionCar/pkg/models"

func (p *PermissionService) PermissionByRoleID(roleID string) (*[]string, error) {
	permissions, err := p.PermissionRepository.PermissionByRoleID(roleID)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener permisos", err)
	}
	return permissions, nil
}