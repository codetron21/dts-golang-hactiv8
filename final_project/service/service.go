package service

import "final_project/repository"

type Service struct {
	UserService UserService
}

func New(repo *repository.Repository) Service {
	return Service{
		UserService: UserService{
			repository: repo.UserRepository,
		},
	}
}
