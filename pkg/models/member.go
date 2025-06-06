package models

import "time"

type Member struct {
	ID        string    `gorm:"primaryKey;size:36" json:"id"`
	UserID    string    `gorm:"not null;size:36" json:"user_id"`
	RoleID    string    `gorm:"not null;size:36" json:"role_id"`
	IsActive  bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Role Role `gorm:"foreignKey:RoleID;references:ID" json:"role"`
}

type MemberResponse struct {
	ID       string  `json:"id"`
	UserID   string  `json:"user_id"`
	RoleID   string  `json:"role_id"`
	IsActive bool    `json:"is_active"`
	Role     Role    `json:"role"`
	UserData UserDTO `json:"user_data"`
}
