package model

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID           int        `gorm:"column:id;primaryKey;not null" json:"id"`
	Username     string     `gorm:"column:username;type:varchar(50);uniqueIndex;not null" json:"username" valid:"required~Username must not be empty"`
	Email        string     `gorm:"column:email;uniqueIndex;not null" json:"email" valid:"email~Email invalid,required~Email must not be empty"`
	Password     string     `gorm:"column:password;not null" json:"password" valid:"required~Password must not be empty,minstringlength(6)~Password length at least have 6 characters"`
	Age          int        `gorm:"column:age;not null" json:"age" valid:"required~Age must not be empty,min(8)~Age min greater than 8"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
}

func (u User) BeforeCreate(db *gorm.DB) error {
	govalidator.ParamTagMap["min"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		minimum := params[0]
		age, err := strconv.Atoi(str)
		if err != nil {
			return false
		}

		minAge, err := strconv.Atoi(minimum)
		if err != nil {
			return false
		}

		return age > minAge
	})

	govalidator.ParamTagRegexMap["min"] = regexp.MustCompile(`^min\\((\\d+)\\)$`)

	isValid, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	if !isValid {
		return errors.New("invalid user data")
	}

	return nil
}
