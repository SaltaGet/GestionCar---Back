package models

type UserTenant struct {
	UserID   string `gorm:"primaryKey" json:"user_id"`
	TenantID string `gorm:"primaryKey" json:"tenant_id"`
	IsActive   bool   `gorm:"not null;default:true" json:"is_active"`
	IsAdmin     bool   `gorm:"not null;default:false" json:"is_admin"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Tenant Tenant `gorm:"foreignKey:TenantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tenant"`
}
