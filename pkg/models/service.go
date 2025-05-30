package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ServiceWorkshop struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ServiceLaundry struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Service struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
}

type ServiceCreate struct {
	Name string `json:"name" validate:"required"`
}

func (s *ServiceCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

type ServiceUpdate struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
}

func (s *ServiceUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}