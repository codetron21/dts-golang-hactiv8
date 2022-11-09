package service

import "final_project/repository"

type Service struct {
	UserService        UserService
	PhotoService       PhotoService
	CommentService     CommentService
	SocialMediaService SocialMediaService
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
		CommentService: CommentService{
			userRepo:    repo.UserRepository,
			commentRepo: repo.CommentRepository,
			photoRepo:   repo.PhotoRepository,
		},
		SocialMediaService: SocialMediaService{
			userRepo:        repo.UserRepository,
			socialMediaRepo: repo.SocialMediaRepository,
		},
	}
}
