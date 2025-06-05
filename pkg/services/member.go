package services

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func (m *MemberService) MemberGetAll() (*[]models.Member, error) {
	members, err := m.MemberRepository.MemberGetAll()
	if err != nil {
		return nil, err
	}

	return members, nil
}
