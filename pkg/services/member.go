package services

import (
	"fmt"
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


func (m *MemberService) MemberGetPermissionByUserID(userID string) (*models.Member, error) {
	member, err := m.MemberRepository.MemberGetPermissionByUserID(userID)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (m *MemberService) MemberAdd(memberAdd *models.MemberAdd, tenantID, userID string) (string, error) { 
	id, err := m.MemberRepository.MemberAdd(memberAdd) 
	if err != nil {
		return "", err
	}

	if err := m.UserRepository.UserTenantAdd(memberAdd.UserID, tenantID); err != nil {
		if delErr := m.MemberRepository.MemberDelete(id); delErr != nil {
        return "", fmt.Errorf("UserTenantAdd failed: %v, rollback also failed: %v", err, delErr)
    }
		return "", err
	}

	return id, nil
}

func (m *MemberService) MemberGetByID(id string) (*models.MemberResponse, error) {
	member, err := m.MemberRepository.MemberGetByID(id)
	if err != nil {
		return nil, err
	}
	
	user, err := m.UserRepository.UserGetByID(member.UserID)
	if err != nil {
		return nil, err
	}

	return &models.MemberResponse{
		ID:       member.ID,
		UserID:   member.UserID,
		RoleID:   member.RoleID,
		IsActive: member.IsActive,
		Role:     member.Role,
		UserData: models.UserDTO{
			ID:       user.ID,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}