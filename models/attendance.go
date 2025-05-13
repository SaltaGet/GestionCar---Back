package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Asistencia empleados
type AttendanceLaundry struct {
	ID         string          `gorm:"primaryKey" json:"id"`
	EmployeeID string          `gorm:"not null" json:"employee_id"`
	Attendance string          `gorm:"not null" json:"role" validate:"oneof=presente tarde parcial ausente"`
	Hours      int             `gorm:"not null;" json:"hours" validate:"max=24"`
	Date       string          `gorm:"not null" json:"date"`
	Amount     float32          `gorm:"not null" json:"amount"`
	IsHoliday  bool            `gorm:"not null;default:false" json:"is_holiday"`
	CreatedAt  time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	Employee   EmployeeLaundry `gorm:"foreignKey:EmployeeID;references:ID" json:"employee"`
}

type AttendanceWorkshop struct {
	ID         string          `gorm:"primaryKey" json:"id"`
	EmployeeID string          `gorm:"not null" json:"employee_id"`
	Attendance string          `gorm:"not null" json:"role" validate:"oneof=presente tarde parcial ausente"`
	Hours      int             `gorm:"not null;" json:"hours" validate:"max=24"`
	Date       string          `gorm:"not null" json:"date"`
	Amount     float32          `gorm:"not null" json:"amount"`
	IsHoliday  bool            `gorm:"not null;default:false" json:"is_holiday"`
	CreatedAt  time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	Employee   EmployeeWorkshop `gorm:"foreignKey:EmployeeID;references:ID" json:"employee"`
}

type AttendanceCreate struct {
	EmployeeID string          `json:"employee_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Attendance string          `json:"role" validate:"oneof=presente tarde parcial ausente"`
	Hours      int             `gjson:"hours" validate:"max=24"`
	Date       string          `json:"date" validate:"required" example:"2022-01-01"`
	Amount     float32          `json:"amount" validate:"required" example:"1234.56"`
	IsHoliday  bool            `json:"is_holiday" default:"false" example:"false"`
}

func (e *AttendanceCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

type AttendanceUpdate struct {
	ID         string          `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" validate:"required"`
	EmployeeID string          `json:"employee_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Attendance string          `json:"role" validate:"oneof=presente tarde parcial ausente"`
	Hours      int             `gjson:"hours" validate:"max=24"`
	Date       string          `json:"date" validate:"required" example:"2022-01-01"`
	Amount     float32          `json:"amount" validate:"required" example:"1234.56"`
	IsHoliday  bool            `json:"is_holiday" default:"false" example:"false"`
}

func (e *AttendanceUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

type DateBetween struct {
    DateFrom string `json:"date_from" validate:"required,datetime=2006-01-02" example:"2022-01-01"`
    DateTo   string `json:"date_to" validate:"required,datetime=2006-01-02" example:"2022-01-01"`
}

func (e *DateBetween) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}