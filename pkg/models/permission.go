package models

type Permission struct {
	ID    int    `gorm:"primaryKey;size:36" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Group string `gorm:"not null" json:"group" validate:"oneof=income expense user supplier product purchase attendance client vehicle resume"`
	Roles []Role `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"roles"`
}
