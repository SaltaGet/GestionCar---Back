package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type EmployeeLaundry struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Phone string `gorm:"not null" json:"phone"`
	Email string `gorm:"not null" json:"email"`
	Address string `gorm:"not null" json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type EmployeeWorkshop struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Phone string `gorm:"not null" json:"phone"`
	Email string `gorm:"not null" json:"email"`
	Address string `gorm:"not null" json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type EmployeeCreate struct {
	Name string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
} 

func (e *EmployeeCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

type EmployeeUpdate struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
}

func (e *EmployeeUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}
