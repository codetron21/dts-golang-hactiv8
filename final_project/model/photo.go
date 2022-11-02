package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int        `gorm:"column:id;primaryKey;not null" json:"id"`
	Title     string     `gorm:"column:title;not null" json:"title" valid:"required~title must not be empty"`
	Caption   string     `gorm:"column:caption" json:"caption"`
	PhotoUrl  string     `gorm:"column:photo_url;not null" json:"photo_url" valid:"required~photo url must not be empty"`
	UserID    int        `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updated_at,omitempty"`
	Comments  []Comment  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"comments,omitempty"`
	User      *User      `json:"User,omitempty"`
}

func (p Photo) BeforeCreate(db *gorm.DB) error {
	isValid, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	if !isValid {
		return errors.New("invalid photo data")
	}

	if !govalidator.IsURL(p.PhotoUrl) {
		return errors.New("photo url invalid")
	}

	return nil
}
