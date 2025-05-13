package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Client struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"not null;size:30" json:"first_name"`
	LastName  string    `gorm:"not null;size:30" json:"last_name"`
	CUIL      string    `gorm:"unique;size:30" json:"cuil"`
	DNI       string    `gorm:"unique;size:30" json:"dni"`
	Email     string    `gorm:"unique" json:"email" validate:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Vehicles  []Vehicle  `gorm:"foreignKey:ClientID" json:"vehicles"`
}

type ClientCreate struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	CUIL      string `json:"cuil" validate:"required"`		
	DNI       string `json:"dni" validate:""`
	Email     string `json:"email" validate:"required,email"`
}

func (c *ClientCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

type ClientUpdate struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	CUIL      string `json:"cuil" validate:"required"`		
	DNI       string `json:"dni" validate:""`
	Email     string `json:"email" validate:"required,email"`
}

func (c *ClientUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}