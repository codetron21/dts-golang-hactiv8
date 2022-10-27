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

func (repo UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	currentUser := model.User{}
	repo.database.DB.First(&currentUser, "id = ?", user.ID)

	currentUser.Email = user.Email
	currentUser.Username = user.Username

	err := repo.database.DB.Save(&currentUser).Error

	if err != nil {
		return nil, err
	}

	return &currentUser, nil
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
