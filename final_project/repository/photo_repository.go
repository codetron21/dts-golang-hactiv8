package repository

import "gorm.io/gorm"

type PhotoRepository struct {
	DB *gorm.DB
}
