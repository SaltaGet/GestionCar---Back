package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type PurchaseOrder struct {
	ID                      string            `gorm:"not null;primaryKey;size:36" json:"id"`
	OrderNumber             string            `gorm:"not null" json:"order_number"`
	OrderDate               string            `gorm:"not null" json:"order_date"`
	Amount                  float32           `gorm:"not null" json:"amount"`
	SupplierID              string            `gorm:"not null;size:36" json:"supplier_id"`
	CreatedAt               time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt               time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
	Supplier                Supplier          `gorm:"foreignKey:SupplierID;references:ID" json:"supplier"`
	PurchaseProducts []PurchaseProduct `gorm:"foreignKey:PurchaseOrderID;references:ID" json:"purchase_products"`
}

type PurchaseOrderCreate struct {
	OrderNumber            string                  `json:"order_number" validate:"required"`
	OrderDate              string                  `json:"order_date" validate:"required"`
	Amount                 float32                 `json:"amount" validate:"required"`
	SupplierID             string                  `json:"supplier_id"`
	PurchaseProductCreates []PurchaseProductCreate `json:"purchase_products" validate:"required,gt=0"`
}

func (p *PurchaseOrderCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()
	param := validationErr.Param()

	return fmt.Errorf("campo %s es invalido, revisar: (%s) (%s)", field, tag, param)
}

type PurchaseOrderUpdate struct {
	ID                     string                  `json:"id" validate:"required"`
	OrderNumber            string                  `json:"order_number" validate:"required"`
	OrderDate              string                  `json:"order_date" validate:"required"`
	Amount                 float32                 `json:"amount" validate:"required"`
	SupplierID             string                  `json:"supplier_id"`
	PurchaseProductUpdates []PurchaseProductUpdate `json:"purchase_products" validate:"required,gt=0"`
}

func (p *PurchaseOrderUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()
	param := validationErr.Param()

	return fmt.Errorf("campo %s es invalido, revisar: (%s) (%s)", field, tag, param)
}
