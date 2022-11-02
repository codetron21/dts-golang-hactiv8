package repository

import (
	"final_project/datasource"
	"final_project/model"
)

type UserRepository struct {
	database *datasource.Database
}

func (repo UserRepository) CreateUser(user *model.User) error {
	return repo.database.DB.Create(user).Error
}

func (repo UserRepository) UpdateUser(user *model.User) error {
	err := repo.database.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}

	err = repo.database.DB.First(user, user.ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) FindUserById(userId int) error {
	user := model.User{ID: userId}
	return repo.database.DB.First(&user, "id = ?", userId).Error
}

func (repo UserRepository) FindUserByEmail(user *model.User, email string) error {
	return repo.database.DB.First(user, "email = ?", email).Error
}

func (repo UserRepository) DeleteUserById(userId int) error {
	user := model.User{ID: userId}
	return repo.database.DB.Delete(&user).Error
}
