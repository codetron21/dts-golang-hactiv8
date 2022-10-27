package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             int        `gorm:"column:id;primaryKey;not null" json:"id"`
	Name           string     `gorm:"column:name;not null" json:"name" valid:"required~Name must not be empty"`
	SocialMediaUrl string     `gorm:"column:social_media_url;not null" json:"social_media_url" valid:"required~Url must not be empty"`
	UserID         int        `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"updated_at" json:"updated_at"`
}

func (sm SocialMedia) BeforeCreate(db *gorm.DB) error {
	isValid, err := govalidator.ValidateStruct(sm)
	if err != nil {
		return err
	}

	if !isValid {
		return errors.New("invalid social media data")
	}

	if !govalidator.IsURL(sm.SocialMediaUrl) {
		return errors.New("social media url invalid")
	}

	return nil
}
