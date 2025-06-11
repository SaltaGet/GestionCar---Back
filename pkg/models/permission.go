package models

type Permission struct {
	ID    string    `gorm:"primaryKey;size:36" json:"id"`
	Code  string `gorm:"not null,unique" json:"code"`
	Details string `gorm:"not null" json:"detail"`
	Group string `gorm:"not null" json:"group" validate:"oneof=income expense user supplier product purchase attendance client vehicle resume"`
	Roles []Role `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"roles"`
}
