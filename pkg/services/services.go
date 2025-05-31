package services

import (
	"github.com/DanielChachagua/GestionCar/pkg/ports"
)

type AttendanceService struct {
	AttendanceRepository ports.AttendanceRespository
}

type AuthService struct {
	AuthRepository ports.AuthRepository
}

type VehicleService struct {
	VehicleRepository ports.VehicleRepository
}

type ClientService struct {
	ClientRepository ports.ClientRepository
}
