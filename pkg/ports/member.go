package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type MemberRepository interface {
	MemeberGetPermissionByUserID(userID string) (member *models.Member, err error)
	MemberGetAll() (members *[]models.Member, err error)
}

type MemberService interface {
	MemeberGetPermissionByUserID(userID string) (permission *models.Member, err error)
	MemberGetAll() (members *[]models.MemberResponse, err error)
}