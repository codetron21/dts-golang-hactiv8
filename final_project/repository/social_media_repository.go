package repository

import "gorm.io/gorm"

type SocialMediaRepository struct {
	DB *gorm.DB
}
