package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ProductLaundry struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	Identifier string    `gorm:"not null;unique" json:"identifier"`
	Name       string    `gorm:"not null" json:"name"`
	Stock      int32     `gorm:"not null;min:0;default:0" json:"stock"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type PartWorkshop struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	Identifier string    `gorm:"not null;unique" json:"identifier"`
	Name       string    `gorm:"not null" json:"name"`
	Stock      int32     `gorm:"not null;min:0;default:0" json:"stock"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ProductCreate struct {
	Identifier string  `json:"identifier" validate:"required"`
	Name       string  `json:"name" validate:"required"`
}

func (p *ProductCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type ProductUpdate struct {
	ID         string  `json:"id" validate:"required"`
	Identifier string  `json:"identifier"`
	Name       string  `json:"name" validate:"required"`
}

func (p *ProductUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type StockUpdate struct {
	Stock int32 `json:"stock" validate:"required"`
}

func (p *StockUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
