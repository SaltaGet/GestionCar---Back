package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Vehicle struct {
	ID   string `gorm:"primaryKey;size:36" json:"id"`
	Brand string `gorm:"not null" json:"brand"`
	Model string ` json:"model"`
	Color string `gorm:"not null" json:"color"`
	Year  string `json:"year"`
	Domain string `gorm:"not null;unique" json:"domain"`
	ClientID string `gorm:"not null;size:36" json:"client_id"`
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
	err := validate.Struct(v)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()

	return fmt.Errorf("campo %s es invalido, revisar: (%s)", field, tag)
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
	err := validate.Struct(v)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()

	return fmt.Errorf("campo %s es invalido, revisar: (%s)", field, tag)
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