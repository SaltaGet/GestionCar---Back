package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"not null;size:30" json:"first_name"`
	LastName  string    `gorm:"not null;size:30" json:"last_name"`
	Username  string    `gorm:"unique;size:30;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email" validate:"email"`
	Password  string    `gorm:"not null" json:"password"`
	IsAdmin      bool    `gorm:"not null;default:false" json:"is_admin"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Tenants      []Tenant   `gorm:"many2many:user_tenants;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tenants"`
}


type UserDTO struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreate struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (u *UserCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type UserUpdate struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Role      string `json:"role" validate:"required,oneof= admin admin_laundry admin_workshop employee_laundry employee_workshop"`
}

func (u *UserUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}