package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Vehicle struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Brand string `gorm:"not null" json:"brand"`
	Model string ` json:"model"`
	Color string `gorm:"not null" json:"color"`
	Year  string `json:"year"`
	Domain string `gorm:"not null;unique" json:"domain"`
	ClientID string `gorm:"not null" json:"client_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Client Client `gorm:"foreignKey:ClientID" json:"client"`
}

type VehicleCreate struct {
	Brand string `json:"brand" validate:"required" example:"Toyota"`
	Model string `json:"model" example:"Corolla or null"`
	Color string `json:"color" validate:"required" example:"Red"`
	Year  string `json:"year" example:"2020"`
	Domain string `json:"domain" validate:"required" example:"ABC123"`
	ClientID string `json:"client_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (v *VehicleCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(v)
}

type VehicleUpdate struct {
	ID    string `json:"id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Brand string `json:"brand" example:"Toyota"`
	Model string `json:"model" example:"Corolla"`
	Color string `json:"color" example:"Red"`
	Year  string `json:"year" example:"2020"`
	Domain string `json:"domain" example:"ABC123"`
	ClientID string `json:"client_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
}

func (v *VehicleUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(v)
}

type VehicleDTO struct {
	ID   string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Color string `json:"color"`
	Year  string `json:"year"`
	Domain string `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}