package service

import (
	"errors"
	"final_project/helpers"
	"final_project/model"
	"final_project/repository"
	"fmt"
)

type UserService struct {
	repository repository.UserRepository
}

func (us UserService) CreateUser(user *model.User) error {
	user.Password = helpers.HashPassword(user.Password)
	err := us.repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us UserService) LoginUser(user *model.User) (jwt string, err error) {
	// find user
	userInDb := model.User{}
	err = us.repository.FindUserByEmail(&userInDb, user.Email)
	if err != nil {
		fmt.Println("Login User (Service) find email:", err.Error())
		err = errors.New("incorrect email or password")
		return
	}

	// check password
	isPassMatch := helpers.ComparePassword([]byte(userInDb.Password), []byte(user.Password))
	if !isPassMatch {
		err = errors.New("incorrect email or password")
		return
	}

	// generate token
	jwt, err = helpers.GenerateToken(userInDb.ID, userInDb.Email)
	if err != nil {
		fmt.Println("Login User (Service) generate token", err.Error())
		err = errors.New("error generate jwt token")
		return
	}

	return
}

func (us UserService) UpdateUserById(user *model.User, userIdFromClaim int) error {
	err := us.repository.FindUserById(userIdFromClaim)
	if err != nil {
		return err
	}

	user.ID = userIdFromClaim
	err = us.repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (us UserService) DeleteUserById(userIdFromClaim int) error {
	err := us.repository.FindUserById(userIdFromClaim)
	if err != nil {
		return err
	}

	err = us.repository.DeleteUserById(userIdFromClaim)
	if err != nil {
		return err
	}

	return nil
}
