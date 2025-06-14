package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Expense struct {
	ID                  string       `gorm:"primaryKey;size:36" json:"id"`
	Details             string       `json:"details"`
	SupplierID          string       `gorm:"not null;size:36" json:"supplier_id"`
	MovementTypeID      string       `gorm:"not null;size:36" json:"movement_type_id"`
	Amount              float32      `gorm:"not null" json:"amount"`
	CreatedAt           time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	Supplier            Supplier     `gorm:"foreignKey:SupplierID" json:"supplier"`
	MovementType MovementType `gorm:"foreignKey:MovementTypeID;references:ID" json:"movement_type"`
}

type ExpenseCreate struct {
	Details        string  `json:"details" validate:"required"`
	SupplierID     string  `json:"supplier_id"`
	MovementTypeID string  `json:"movement_type_id" validate:"required"`
	Amount         float32 `json:"amount" validate:"required"`
}

func (e *ExpenseCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()
	param := validationErr.Param()

	return fmt.Errorf("campo %s es invalido, revisar: (%s) (%s)", field, tag, param)
}

type ExpenseUpdate struct {
	ID             string  `json:"id"`
	Details        string  `json:"details" validate:"required"`
	// SupplierID     string  `json:"supplier_id"` sacar
	MovementTypeID string  `json:"movement_type_id" validate:"required"`
	Amount         float32 `json:"amount" validate:"required"`
}

func (e *ExpenseUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()
	param := validationErr.Param()

	return fmt.Errorf("campo %s es invalido, revisar: (%s) (%s)", field, tag, param)
}
