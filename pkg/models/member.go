package models

import "time"

type Member struct {
	ID          string `gorm:"primaryKey" json:"id"`
	UserID        string `gorm:"not null" json:"name"`
	RoleID        string `gorm:"not null;foreignKey:RoleID" json:"role_id"`
	IsActive    bool   `gorm:"not null;default:true" json:"is_active"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}