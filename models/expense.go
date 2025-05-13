package models

import "time"

type ExpenseLaundry struct {
	ID                    string              `gorm:"primaryKey" json:"id"`
	Details               string              `json:"details"`
	SupplierID            string              `json:"supplier_id"`
	MovementTypeID            string          `gorm:"not null" json:"employee_id"`
	Amount                float32             `gorm:"not null" json:"amount"`
	CreatedAt             time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	Supplier              SupplierLaundry     `gorm:"foreignKey:ClientID" json:"client"`
	MovementTypeLaundry   MovementTypeLaundry `gorm:"foreignKey:MovementTypeLaundryID;references:ID" json:"movement_type_laundry"`
}

type ExpenseWorkshop struct {
	ID                    string              `gorm:"primaryKey" json:"id"`
	Details               string              `json:"details"`
	SupplierID            string              `json:"supplier_id"`
	MovementTypeID            string          `gorm:"not null" json:"employee_id"`
	Amount                float32             `gorm:"not null" json:"amount"`
	CreatedAt             time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	Supplier              SupplierWorkshop     `gorm:"foreignKey:SupplierID" json:"supplier"`
	MovementTypeWorkshop   MovementTypeWorkshop `gorm:"foreignKey:MovementTypeWorkshopID;references:ID" json:"movement_type_workshop"`
}

type ExpenseCreate struct {
	Details        string   `json:"details" validate:"required"`
	SupplierID            string              `json:"supplier_id"`
	MovementTypeID string   `json:"movement_type_id" validate:"required"`
	Amount         float32  `json:"amount" validate:"required"`
}

type ExpenseUpdate struct {
	ID             string   `json:"id"`
	Details        string   `json:"details" validate:"required"`
	SupplierID            string              `json:"supplier_id"`
	MovementTypeID string   `json:"movement_type_id" validate:"required"`
	Amount         float32  `json:"amount" validate:"required"`
}
