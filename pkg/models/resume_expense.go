package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ResumeExpense struct {
	ID        string    `gorm:"primaryKey;size:36" json:"id"`
	Data      string    `gorm:"not null;size:100000" json:"data"`
	Date      time.Time `gorm:"not null" json:"date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type ResumeExpenseCreate struct {
	Data string `json:"data" validate:"required"`
	Date time.Time `json:"date" validate:"required"`
}

func (e *ResumeExpenseCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

type ResumeExpenseUpdate struct {
	ID   string `json:"id" validate:"required"`
	Data string `json:"data" validate:"required"`
}

func (e *ResumeExpenseUpdate) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}



