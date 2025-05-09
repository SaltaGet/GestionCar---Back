package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Proveedor
type SupplierLaundry struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Address   string `gorm:"not null" json:"address"`
	Phone     string `gorm:"not null" json:"phone"`
	Email     string `gorm:"not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}


type SupplierWorkshop struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Address   string `gorm:"not null" json:"address"`
	Phone     string `gorm:"not null" json:"phone"`
	Email     string `gorm:"not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type SupplierCreate struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

func (s *SupplierCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

type SupplierUpdate struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

func (s *SupplierUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}