package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID           int           `gorm:"column:id;primaryKey;not null" json:"id"`
	Username     string        `gorm:"column:username;type:varchar(50);uniqueIndex;not null" json:"username" valid:"required~username must not be empty"`
	Email        string        `gorm:"column:email;uniqueIndex;not null" json:"email" valid:"email~Email invalid,required~email must not be empty"`
	Password     string        `gorm:"column:password;not null" json:"password,omitempty" valid:"required~password must not be empty,minstringlength(6)~password length at least have 6 characters"`
	Age          int           `gorm:"column:age;not null" json:"age,omitempty" valid:"required~age must not be empty"`
	CreatedAt    *time.Time    `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt    *time.Time    `gorm:"column:updated_at" json:"updated_at,omitempty"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photos,omitempty"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"comments:photos,omitempty"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"social_medias:photos,omitempty"`
}

func (u User) BeforeCreate(db *gorm.DB) error {
	isValid, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	if !isValid {
		return errors.New("invalid user data")
	}

	if u.Age < 9 {
		return errors.New("age must greater than 8")
	}

	return nil
}
