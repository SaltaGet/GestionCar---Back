package models

import "time"

type ExpenseLaundry struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Data string `gorm:"not null;size:100000" json:"data"`
	Date time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ExpenseWorkshop struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Data string `gorm:"not null;size:100000" json:"data"`
	Date time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ExpenseCreate struct {
	ID        string    `json:"id"`
	PurchaseOrderID string `json:"purchase_order_id"`
	Detail string `json:"detail"`
	Amount string `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
