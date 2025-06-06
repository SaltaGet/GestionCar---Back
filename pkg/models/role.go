package models

type Role struct {
	ID          string       `gorm:"primaryKey;size:36" json:"id"`
	Name        string       `gorm:"not null;unique" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"permissions"`
}
