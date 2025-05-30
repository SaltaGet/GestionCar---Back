package models

import "time"

type Workplace struct {
	ID   string    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Address string `gorm:"not null" json:"address"`
	Phone string `gorm:"not null" json:"phone"`
	Email string `gorm:"not null" json:"email"`
	Identifier string `gorm:"not null;unique" validate:"oneof=laundry workshop" json:"identifier"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}