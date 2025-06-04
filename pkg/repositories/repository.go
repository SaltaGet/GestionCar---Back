package repositories

import (
	"gorm.io/gorm"
)

// type Repository struct {
// 	DB *gorm.DB
// }

type MainRepository struct {
	DB *gorm.DB
}

type TenantRepository struct {
	DB *gorm.DB
}
