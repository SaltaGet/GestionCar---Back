package models

import "time"

type PurchaseProductLaundry struct {
	ID        string  `gorm:"primaryKey" json:"id"`
	ProductID string  `gorm:"not null" json:"product_id"`
	PurchaseOrderID string `gorm:"not null" json:"purchase_order_id"`
	ExpiredAt string  `gorm:"not null" json:"expired_at"`
	UnitPrice  float32 `gorm:"not null" json:"unit_price"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	TotalPrice float32 `gorm:"not null" json:"total_price"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Product   ProductLaundry `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	PurchaseOrder PurchaseOrderLaundry `gorm:"foreignKey:PurchaseOrderID;references:ID" json:"purchase_order"`
}

type PurchasePartWorkshop struct {
	ID        string  `gorm:"primaryKey" json:"id"`
	PartID string  `gorm:"not null" json:"part_id"`
	PurchaseOrderID string `gorm:"not null" json:"purchase_order_id"`
	ExpiredAt string  `gorm:"not null" json:"expired_at"`
	UnitPrice  float32 `gorm:"not null" json:"unit_price"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	TotalPrice float32 `gorm:"not null" json:"total_price"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	PartWorkshop PartWorkshop `gorm:"foreignKey:PartID;references:ID" json:"part"`
	PurchaseOrderWorkshop PurchaseOrderWorkshop `gorm:"foreignKey:PurchaseOrderID;references:ID" json:"purchase_order"`
}