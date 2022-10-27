package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        int        `gorm:"column:id;primaryKey;not null" json:"id"`
	UserID    int        `gorm:"column:user_id;not null" json:"user_id"`
	PhotoID   int        `gorm:"column:photo_id;not null" json:"photo_id"`
	Message   string     `gorm:"column:message;not null" json:"message" valid:"required~Message must not be empty"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updated_at"`
}

func (c Comment) BeforeCreate(db *gorm.DB) error {
	isValid, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	if !isValid {
		return errors.New("invalid comment data")
	}

	return nil
}
