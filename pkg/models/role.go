package models

type Role struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string    `gorm:"not null;unique" json:"name"`
	Permissions      []Permission   `gorm:"many2many:user_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"skills"`
}