package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ExpenseResumeLaundry struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Data      string    `gorm:"not null;size:100000" json:"data"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ExpenseResumeWorkshop struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Data      string    `gorm:"not null;size:100000" json:"data"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ExpenseResumeCreate struct {
	Data string `json:"data" validate:"required"`
	Date time.Time `json:"date" validate:"required"`
}

func (e *ExpenseResumeCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

type ExpenseResumeUpdate struct {
	ID   string `json:"id" validate:"required"`
	Data string `json:"data" validate:"required"`
}

func (e *ExpenseResumeUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}



