package repository

import (
	"errors"
	"final_project/datasource"
	"final_project/model"
	"fmt"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	database *datasource.Database
}

func (repo PhotoRepository) CreatePhoto(photo *model.Photo) error {
	return repo.database.DB.Create(photo).Error
}

func (repo PhotoRepository) GetAllPhotoById(userId int) (*[]model.Photo, error) {
	photos := []model.Photo{}

	err := repo.database.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Email")
		}).
		Find(&photos).
		Error

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error get all photo by id")
	}

	return &photos, nil
}

func (repo PhotoRepository) DeletePhotoById(photoId int) error {
	if !repo.CheckPhotoIsExistByID(photoId) {
		return errors.New("photo not found")
	}

	err := repo.database.DB.Where("id = ?", photoId).Delete(&model.Photo{}).Error
	if err != nil {
		fmt.Println("error delete photo id", err)
		return errors.New("error delete photo")
	}

	return nil
}

func (repo PhotoRepository) CheckPhotoIsExistByID(photoId int) bool {
	err := repo.database.DB.Where("id = ?", photoId).First(&model.Photo{}).Error
	if err != nil {
		fmt.Println("error check photo by id", err)
		return false
	}

	return true
}

func (repo PhotoRepository) CheckPhotoIsExistByUserID(photo *model.Photo) bool {
	err := repo.database.DB.Where("user_id = ? AND id = ?", photo.UserID, photo.ID).First(&photo).Error
	if err != nil {
		fmt.Println("error check photo by user id", err)
		return false
	}

	return true
}

func (repo PhotoRepository) UpdatePhoto(photo *model.Photo) error {
	if !repo.CheckPhotoIsExistByUserID(photo) {
		return errors.New("photo with user id not found")
	}

	err := repo.database.DB.Updates(photo).Error
	if err != nil {
		fmt.Println("error update photo", err)
		return errors.New("error update photo")
	}

	return nil
}
