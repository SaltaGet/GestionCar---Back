package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

// Asistencia empleados
type Attendance struct {
	ID         string    `gorm:"primaryKey;size:36" json:"id"`
	EmployeeID string    `gorm:"size:36" json:"employee_id"`
	Attendance string    `gorm:"not null" json:"attendance" validate:"oneof=presente tarde parcial ausente"`
	Hours      int       `gorm:"not null;" json:"hours" validate:"max=24"`
	Date       string    `gorm:"not null" json:"date"`
	Amount     float32   `gorm:"not null" json:"amount"`
	IsHoliday  bool      `gorm:"not null;default:false" json:"is_holiday"`
	IsPaid     bool      `gorm:"not null;default:false" json:"is_paid"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Employee Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;" json:"employee,omitzero"`
}

type AttendanceCreate struct {
	EmployeeID string  `json:"employee_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Attendance string  `json:"attendance" validate:"oneof=presente tarde parcial ausente"`
	Hours      int     `json:"hours" validate:"max=24"`
	Date       string  `json:"date" validate:"required" example:"2022-01-01"`
	Amount     float32 `json:"amount" validate:"required" example:"1234.56"`
	IsHoliday  bool    `json:"is_holiday" default:"false" example:"false"`
}

func (e *AttendanceCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()
	param := validationErr.Param()

	return fmt.Errorf("campo %s es invalido, revisar: (%s) (%s)", field, tag, param)
}

type AttendanceUpdate struct {
	ID         string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" validate:"required"`
	EmployeeID string  `json:"employee_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Attendance string  `json:"attendance" validate:"oneof=presente tarde parcial ausente"`
	Hours      int     `json:"hours" validate:"max=24"`
	Date       string  `json:"date" validate:"required" example:"2022-01-01"`
	Amount     float32 `json:"amount" validate:"required" example:"1234.56"`
	IsHoliday  bool    `json:"is_holiday" default:"false" example:"false"`
}

func (e *AttendanceUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()
	param := validationErr.Param()

	return fmt.Errorf("campo %s es invalido, revisar: (%s) (%s)", field, tag, param)
}

type DateBetween struct {
	DateFrom string `json:"date_from" validate:"required,datetime=2006-01-02" example:"2022-01-01"`
	DateTo   string `json:"date_to" validate:"required,datetime=2006-01-02" example:"2022-01-01"`
}

func (e *DateBetween) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()
	param := validationErr.Param()

	return fmt.Errorf("campo %s es invalido, revisar: (%s) (%s)", field, tag, param)
}

type UpdatePay struct {
	ListIDs []string `json:"list_ids" validate:"required,min=1" example:"123e4567-e89b-12d3-a456-426614174000,123e4567-e89b-12d3-a456-426614174000"`
}

func (e *UpdatePay) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()
	param := validationErr.Param()

	return fmt.Errorf("campo %s es invalido, revisar: (%s) (%s)", field, tag, param)
}

type AttendanceDTO struct {
	ID         string    `json:"id"`
	EmployeeID string    `json:"employee_id"`
	Attendance string    `json:"attendance"`
	Hours      int       `json:"hours"`
	Date       string    `json:"date"`
	Amount     float32   `json:"amount"`
	IsHoliday  bool      `json:"is_holiday"`
	IsPaid     bool      `json:"is_paid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}



// type AttendanceResponse struct {
// 	ID         string    `json:"id"`
// 	EmployeeID string    `json:"employee_id"`
// 	Attendance string    `json:"attendance"`
// 	Hours      int       `json:"hours" validate:"max=24"`
// 	Date       string    `json:"date"`
// 	Amount     float32   `json:"amount"`
// 	IsHoliday  bool      `json:"is_holiday"`
// 	IsPaid     bool      `json:"is_paid"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// }
