package models

type IncomeServiceLaundry struct {
	ID string `gorm:"primaryKey" json:"id"`
	IncomeLaundryID string `gorm:"not null" json:"income_laundry_id"`
	ServiceID string `gorm:"not null" json:"service_id"`
	IncomeLaundry IncomeLaundry `gorm:"foreignKey:IncomeLaundryID;references:ID" json:"income_laundry"`
	Service ServiceLaundry `gorm:"foreignKey:ServiceID;references:ID" json:"service"`
}

type IncomeServiceWorkshop struct {
	ID string `gorm:"primaryKey" json:"id"`
	IncomeWorkshopID string `gorm:"not null" json:"income_workshop_id"`
	ServiceID string `gorm:"not null" json:"service_id"`
	IncomeWorkshop IncomeWorkshop `gorm:"foreignKey:IncomeWorkshopID;references:ID" json:"income_workshop"`
	Service ServiceWorkshop `gorm:"foreignKey:ServiceID;references:ID" json:"service"`
}