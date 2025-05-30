package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type MovementTypeLaundry struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	IsIncome bool   `gorm:"not null" json:"is_income"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`	
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MovementTypeWorkshop struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	IsIncome bool   `gorm:"not null" json:"is_income"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MovementTypeCreate struct {
	Name string `json:"name"`
	IsIncome bool   `json:"is_income"`
}

func (m *MovementTypeCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type MovementTypeUpdate struct {
	ID string `gojson:"id"`
	Name string `json:"name"`
	IsIncome bool   `json:"is_income"`
}

func (m *MovementTypeUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}