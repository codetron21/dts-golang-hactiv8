package repository

import (
	"final_project/datasource"
	"final_project/model"
)

type PhotoRepository struct {
	database *datasource.Database
}

func (repo PhotoRepository) CreatePhoto(photo *model.Photo) error {
	return repo.database.DB.Create(photo).Error
}

func (repo PhotoRepository) GetAllPhotoById(userId int) (*[]model.Photo, error) {
	photos := []model.Photo{}
	err := repo.database.DB.Where("user_id = ?", userId).Find(&photos).Error
	return &photos, err
}
