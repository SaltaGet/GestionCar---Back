package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	ID        string    `gorm:"primaryKey;size:36" json:"id"`
	Name      string    `gorm:"not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Incomes []Income `gorm:"many2many:income_services;" json:"incomes"`
}

type ServiceCreate struct {
	Name string `json:"name" validate:"required"`
}

func (s *ServiceCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()

	return fmt.Errorf("campo %s es invalido, revisar: (%s)", field, tag)
}

type ServiceUpdate struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
}

func (s *ServiceUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()

	return fmt.Errorf("campo %s es invalido, revisar: (%s)", field, tag)
}