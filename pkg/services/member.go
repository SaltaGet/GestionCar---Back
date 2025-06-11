package services

import (
	"log"
	"runtime/debug"

	"github.com/DanielChachagua/GestionCar/pkg/models"
)

func (m *MemberService) MemberGetAll() (*[]models.MemberResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic atrapado en MemberGetAll: %v", r)
			debug.PrintStack()
		}
	}()
	members, err := m.MemberRepository.MemberGetAll()
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0, len(*members))
	for _, member := range *members {
		ids = append(ids, member.ID)
	}

	users, err := m.UserRepository.UserGetByListID(ids)
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]models.UserDTO)
	for _, user := range *users {
		userMap[user.ID] = user
	}

	membersResponse := make([]models.MemberResponse, 0, len(*members))
	for _, member := range *members {
		user, exists := userMap[member.ID]
		if !exists {
			continue 
		}
		membersResponse = append(membersResponse, models.MemberResponse{
			ID:       member.ID,
			RoleID:   member.RoleID,
			IsActive: member.IsActive,
			Role:     member.Role,
			UserData: user,
		})
	}

	return &membersResponse, nil
}


func (m *MemberService) MemberGetPermissionByUserID(userID string) (*models.Member, error) {
	member, err := m.MemberRepository.MemberGetPermissionByUserID(userID)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (m *MemberService) MemberGetByID(id string) (*models.Member, error) {
	member, err := m.MemberRepository.MemberGetByID(id)
	if err != nil {
		return nil, err
	}
	
	return member, nil
}

func (m *MemberService) MemberCreate(memberCreate *models.MemberCreate, user *models.AuthenticatedUser) (string, error) {
	id, err := m.MemberRepository.MemberCreate(memberCreate, user)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (m *MemberService) MemberDelete(memberID string) error {
	if err := m.MemberRepository.MemberDelete(memberID); err != nil {
		return err
	}
	return nil
}