package services

import (
	"github.com/DanielChachagua/GestionCar/pkg/ports"
)

type AttendanceService struct {
	AttendanceRepository ports.AttendanceRespository
}

type AuthService struct {
	AuthRepository ports.AuthRepository
	UserRepository ports.UserRepository
}

type ClientService struct {
	ClientRepository ports.ClientRepository
}

type EmployeeService struct {
	EmployeeRepository ports.EmployeeRepository
}

type ExpenseService struct {
	ExpenseRepository ports.ExpenseRepository
}

type IncomeService struct {
	IncomeRepository ports.IncomeRepository
}

type MovementTypeService struct {
	MovementTypeRepository ports.MovementTypeRepository
}

type ProductService struct {
	ProductRepository ports.ProductRepository
}

type VehicleService struct {
	VehicleRepository ports.VehicleRepository
}
