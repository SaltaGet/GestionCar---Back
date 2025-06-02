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


type MemberController struct {
	MemberService ports.MemberService
}

type RoleController struct {
	RoleService ports.RoleService
}

type PermissionController struct {
	PermissionService ports.PermissionService
}