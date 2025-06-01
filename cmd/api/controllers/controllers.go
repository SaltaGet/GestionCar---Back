package controllers

import "github.com/DanielChachagua/GestionCar/pkg/ports"

type AuthController struct {
	AuthService ports.AuhtService
}

type UserController struct {
	UserService ports.UserService
}

type TenantController struct {
	TenantService ports.TenantService
}