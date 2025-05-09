package models

import "time"

type IncomeLaundry struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Data      string    `gorm:"not null;size:100000" json:"data"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type IncomeWorkshop struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Data      string    `gorm:"not null;size:100000" json:"data"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type IncomeCreateLaundry struct {
	ID         string           `json:"id"`
	Services   []Service `json:"services"`
	Details    string           `json:"details"`
	ClientID   string           `json:"client_id"`
	VehicleID  string           `json:"vehicle_id"`
	EmployeeID string           `json:"employee_id"`
	Amount     float32          `json:"amount"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}
