package models

type IncomeService struct {
	ID string `gorm:"primaryKey" json:"id"`
	IncomeLaundryID string `gorm:"not null" json:"income_laundry_id"`
	ServiceID string `gorm:"not null" json:"service_id"`
	IncomeLaundry Income `gorm:"foreignKey:IncomeLaundryID;references:ID" json:"income_laundry"`
	Service Service `gorm:"foreignKey:ServiceID;references:ID" json:"service"`
}
