package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type MemberRepository interface {
	MemberGetByID(id string) (member *models.Member, err error)
	MemberGetPermissionByUserID(userID string) (member *models.Member, err error)
	MemberGetAll() (members *[]models.Member, err error)
	MemberCreate(memeberCreate *models.MemberCreate, user *models.AuthenticatedUser) (id string, err error)
	MemberDelete(id string) (err error)
}

type MemberService interface {
	MemberGetByID(id string) (member *models.Member, err error)
	MemberGetPermissionByUserID(userID string) (permission *models.Member, err error)
	MemberGetAll() (members *[]models.MemberResponse, err error)
	MemberCreate(memeberCreate *models.MemberCreate, user *models.AuthenticatedUser) (id string, err error)
	MemberDelete(id string) (err error)
}