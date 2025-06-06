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
		ids = append(ids, member.UserID)
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
		user, exists := userMap[member.UserID]
		if !exists {
			continue 
		}
		membersResponse = append(membersResponse, models.MemberResponse{
			ID:       member.ID,
			UserID:   member.UserID,
			RoleID:   member.RoleID,
			IsActive: member.IsActive,
			Role:     member.Role,
			UserData: user,
		})
	}

	return &membersResponse, nil
}


func (m *MemberService) MemeberGetPermissionByUserID(userID string) (*models.Member, error) {
	member, err := m.MemberRepository.MemeberGetPermissionByUserID(userID)
	if err != nil {
		return nil, err
	}

	return member, nil
}