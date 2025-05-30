package models

type Permission struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Group string `gorm:"not null" json:"group" validate:"oneof=income expense user supplier product purchase attendance client vehicle resume"`
}