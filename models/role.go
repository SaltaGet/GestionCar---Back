package models

type Role struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string    `gorm:"not null" json:"name"`
	Hierarchy int    `gorm:"not null" json:"hierarchy"`
	Workplace string `gorm:"not null" json:"workplace" validate:"oneof=all laundry workshop"`
}