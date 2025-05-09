package models

import "time"

type DailyMovementsLaundry struct {
	ID            string      `gorm:"primaryKey" json:"id"`
	IsIncome  bool				`gorm:"not null" json:"is_income"`
	Detail        AnyMovement `gorm:"not null;type:json" json:"detail"`
	Amount        float32     `gorm:"not null" json:"amount"`
	PaymentMethod string      `gorm:"not null" json:"payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	MovementTypeID string      `gorm:"not null" json:"movement_type_id"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	MovementTypeLaundry MovementTypeLaundry `gorm:"foreignKey:MovementTypeID;references:ID" json:"movement_type_laundry"`
}

type DailyMovementsWorkshop struct {
	ID            string      `gorm:"primaryKey" json:"id"`
	IsIncome  bool				`gorm:"not null" json:"is_income"`
	Detail        AnyMovement `gorm:"not null;type:json" json:"detail"`
	Amount        float32     `gorm:"not null" json:"amount"`
	PaymentMethod string      `gorm:"not null" json:"payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	MovementTypeID string      `gorm:"not null" json:"movement_type_id"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	MovementTypeLaundry MovementTypeWorkshop `gorm:"foreignKey:MovementTypeID;references:ID" json:"movement_type_laundry"`
}