package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Expense struct {
	ID                  string              `gorm:"primaryKey;size:36" json:"id"`
	Details             string              `json:"details"`
	SupplierID          string              `gorm:"not null;size:36" json:"supplier_id"`
	MovementTypeID      string              `gorm:"not null;size:36" json:"movement_type_id"`
	Amount              float32             `gorm:"not null" json:"amount"`
	CreatedAt           time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	Supplier            Supplier     `gorm:"foreignKey:SupplierID" json:"supplier"`
	MovementTypeLaundry MovementType `gorm:"foreignKey:MovementTypeID;references:ID" json:"movement_type_laundry"`
}

type ExpenseCreate struct {
	Details        string  `json:"details" validate:"required"`
	SupplierID     string  `json:"supplier_id"`
	MovementTypeID string  `json:"movement_type_id" validate:"required"`
	Amount         float32 `json:"amount" validate:"required"`
}

func (e *ExpenseCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

type ExpenseUpdate struct {
	ID             string  `json:"id"`
	Details        string  `json:"details" validate:"required"`
	SupplierID     string  `json:"supplier_id"`
	MovementTypeID string  `json:"movement_type_id" validate:"required"`
	Amount         float32 `json:"amount" validate:"required"`
}

func (e *ExpenseUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}
