package service

import "final_project/repository"

type Service struct {
	UserService  UserService
	PhotoService PhotoService
}

func New(repo *repository.Repository) Service {
	return Service{
		UserService: UserService{
			repository: repo.UserRepository,
		},
		PhotoService: PhotoService{
			userRepo:  repo.UserRepository,
			photoRepo: repo.PhotoRepository,
		},
	}
}
