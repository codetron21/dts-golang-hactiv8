package repository

import (
	"errors"
	"final_project/datasource"
	"final_project/model"
	"fmt"
)

type UserRepository struct {
	database *datasource.Database
}

func (repo UserRepository) CreateUser(user *model.User) error {
	err := repo.database.DB.Create(user).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error create user")
	}

	return nil
}

func (repo UserRepository) UpdateUser(user *model.User) error {
	err := repo.database.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error update user")
	}

	err = repo.database.DB.First(user, user.ID).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error update user")
	}

	return nil
}

func (repo UserRepository) FindUserById(userId int) error {
	user := model.User{ID: userId}
	err := repo.database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error find user by id")
	}

	return nil
}

func (repo UserRepository) FindUserByEmail(user *model.User, email string) error {
	err := repo.database.DB.First(user, "email = ?", email).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error find user by email")
	}

	return nil
}

func (repo UserRepository) DeleteUserById(userId int) error {
	user := model.User{ID: userId}
	err := repo.database.DB.Delete(&user).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error delete user by id")
	}

	return nil
}
