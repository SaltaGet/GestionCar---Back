package models

import "time"

type ExpenseResumeLaundry struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Data      string    `gorm:"not null;size:100000" json:"data"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ExpenseResumeWorkshop struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Data      string    `gorm:"not null;size:100000" json:"data"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}