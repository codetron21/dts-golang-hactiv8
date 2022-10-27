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
		fmt.Println("Create User (Service):", err.Error())
		return errors.New("error creating user")
	}

	return nil
}

func (us UserService) LoginUser(user *model.User) (jwt string, err error) {
	// find user
	userInDb := model.User{}
	err = us.repository.FindUserByEmail(&userInDb, user.Email)
	if err != nil {
		fmt.Println("Login User (Service) find email:", err.Error())
		err = errors.New("user cannot be found")
		return
	}

	// check password
	isPassMatch := helpers.ComparePassword([]byte(userInDb.Password), []byte(user.Password))
	if !isPassMatch {
		err = errors.New("incorrect password")
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

func (us UserService) UpdateUserById(user *model.User, userIdFromClaim int, userId int) (*model.User, error) {
	if userIdFromClaim != userId {
		return nil, errors.New("invalid id")
	}

	err := us.repository.FindUserById(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.ID = userId
	updatedUser, err := us.repository.UpdateUser(user)
	if err != nil {
		return nil, errors.New("error update user")
	}

	return updatedUser, nil
}

func (us UserService) DeleteUserById(userIdFromClaim int, userId int) error {
	if userIdFromClaim != userId {
		return errors.New("invalid id")
	}

	err := us.repository.FindUserById(userId)
	if err != nil {
		return errors.New("user not found")
	}

	err = us.repository.DeleteUserById(userId)
	if err != nil {
		return err
	}

	return nil
}
