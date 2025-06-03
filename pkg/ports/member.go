package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type MemberRepository interface {
	MemberGetAll() (members *[]models.Member, err error)
}

type MemberService interface {
	MemberGetAll() (members *[]models.Member, err error)
}