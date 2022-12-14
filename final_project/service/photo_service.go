package service

import (
	"final_project/model"
	"final_project/repository"
)

type PhotoService struct {
	userRepo  repository.UserRepository
	photoRepo repository.PhotoRepository
}

func (ps PhotoService) CreatePhoto(userId int, photo *model.Photo) error {
	err := ps.userRepo.FindUserById(userId)
	if err != nil {
		return err
	}

	photo.UserID = userId
	err = ps.photoRepo.CreatePhoto(photo)
	if err != nil {
		return err
	}

	return nil
}

func (ps PhotoService) GetPhotos(userId int) ([]model.Photo, error) {
	err := ps.userRepo.FindUserById(userId)
	if err != nil {
		return nil, err
	}

	photos, err := ps.photoRepo.GetAllPhotoById(userId)
	if err != nil {
		return nil, err
	}

	return *photos, nil
}

func (ps PhotoService) UpdatePhotoById(photo *model.Photo, photoId int, userId int) error {
	photo.ID = photoId
	photo.UserID = userId

	err := ps.userRepo.FindUserById(photo.UserID)
	if err != nil {
		return err
	}

	err = ps.photoRepo.UpdatePhoto(photo)
	if err != nil {
		return err
	}

	return nil
}

func (ps PhotoService) DeletePhotoById(photoId int, userId int) error {
	err := ps.userRepo.FindUserById(userId)
	if err != nil {
		return err
	}

	err = ps.photoRepo.DeletePhotoById(photoId)
	if err != nil {
		return err
	}

	return nil
}
