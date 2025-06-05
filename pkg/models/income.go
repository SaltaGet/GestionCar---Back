package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Income struct {
	ID                  string              `gorm:"primaryKey" json:"id"`
	Ticket              string              `json:"ticket"`
	Details             string              `json:"details"`
	ClientID            string              `gorm:"not null" json:"client_id"`
	VehicleID           string              `gorm:"not null" json:"vehicle_id"`
	EmployeeID          string              `json:"employee_id"`
	Amount              float32             `gorm:"not null" json:"amount"`
	MovementTypeID      string              `gorm:"not null" json:"movement_type_id"`
	CreatedAt           time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
	Client              Client              `gorm:"foreignKey:ClientID" json:"client"`
	Vehicle             Vehicle             `gorm:"foreignKey:VehicleID" json:"vehicle"`
	Employee     Employee     `gorm:"foreignKey:EmployeeID" json:"employee"`
	MovementType MovementType `gorm:"foreignKey:MovementTypeID;references:ID" json:"movement_type"`
	Services      []Service   `gorm:"many2many:income_services;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"services"`
}

type IncomeCreate struct {
	Ticket         string   `json:"ticket" validate:"required"`
	ServicesID     []string `json:"services_id" validate:"required,gt=0"`
	Details        string   `json:"details" validate:"required"`
	ClientID       string   `json:"client_id" validate:"required"`
	VehicleID      string   `json:"vehicle_id" validate:"required"`
	EmployeeID     string   `json:"employee_id"`
	MovementTypeID string   `json:"movement_type_id" validate:"required"`
	Amount         float32  `json:"amount" validate:"required"`
}

func (i *IncomeCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}

type IncomeUpdate struct {
	ID             string   `json:"id"`
	Ticket         string   `json:"ticket" validate:"required"`
	ServicesID     []string `json:"services_id" validate:"required,gt=0"`
	Details        string   `json:"details"`
	ClientID       string   `json:"client_id" validate:"required"`
	VehicleID      string   `json:"vehicle_id" validate:"required"`
	EmployeeID     string   `json:"employee_id"`
	MovementTypeID string   `json:"movement_type_id" validate:"required"`
	Amount         float32  `json:"amount" validate:"required"`
}

func (i *IncomeUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}
