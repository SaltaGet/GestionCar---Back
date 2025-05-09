package models

import "time"

type PurchaseOrderLaundry struct {
	ID            string `gorm:"not null;primaryKey" json:"id"`
	OrderNumber   string `gorm:"not null" json:"order_number"`
	OrderDate     string `gorm:"not null" json:"order_date"`
	Amount        string `gorm:"not null" json:"amount"`
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
	Amount        string `gorm:"not null" json:"amount"`
	SupplierID string  `gorm:"not null" json:"supplier_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Supplier      SupplierWorkshop `gorm:"foreignKey:SupplierID;references:ID" json:"supplier"`
	PurchasePartWorkshops []PurchasePartWorkshop `gorm:"foreignKey:PurchaseOrderID;references:ID" json:"purchase_parts"`
}