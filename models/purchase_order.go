package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type PurchaseOrderLaundry struct {
	ID            string `gorm:"not null;primaryKey" json:"id"`
	OrderNumber   string `gorm:"not null" json:"order_number"`
	OrderDate     string `gorm:"not null" json:"order_date"`
	Amount        float32 `gorm:"not null" json:"amount"`
	SupplierID string  `gorm:"not null" json:"supplier_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Supplier      SupplierLaundry `gorm:"foreignKey:SupplierID;references:ID" json:"supplier"`
	PurchaseProductLaundrys []PurchaseProductLaundry `gorm:"foreignKey:PurchaseOrderID;references:ID" json:"purchase_products"`
}

type PurchaseOrderWorkshop struct {
	ID            string `gorm:"not null;primaryKey" json:"id"`
	OrderNumber   string `gorm:"not null" json:"order_number"`
	OrderDate     string `gorm:"not null" json:"order_date"`
	Amount        float32 `gorm:"not null" json:"amount"`
	SupplierID string  `gorm:"not null" json:"supplier_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Supplier      SupplierWorkshop `gorm:"foreignKey:SupplierID;references:ID" json:"supplier"`
	PurchasePartWorkshops []PurchasePartWorkshop `gorm:"foreignKey:PurchaseOrderID;references:ID" json:"purchase_parts"`
}

type PurchaseOrderCreate struct {
	OrderNumber   string `json:"order_number" validate:"required"`
	OrderDate     string `json:"order_date" validate:"required"`
	Amount        float32 `json:"amount" validate:"required"`
	SupplierID string  `json:"supplier_id"`
	PurchaseProductCreates []PurchaseProductCreate `json:"purchase_products" validate:"required,gt=0"`
}

func (p *PurchaseOrderCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type PurchaseOrderUpdate struct {
	ID            string `json:"id" validate:"required"`
	OrderNumber   string `json:"order_number" validate:"required"`
	OrderDate     string `json:"order_date" validate:"required"`
	Amount        float32 `json:"amount" validate:"required"`
	SupplierID string  `json:"supplier_id"`
	PurchaseProductUpdates []PurchaseProductUpdate `json:"purchase_products" validate:"required,gt=0"`
}

func (p *PurchaseOrderUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}