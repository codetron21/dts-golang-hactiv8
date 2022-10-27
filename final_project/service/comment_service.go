package service

import "final_project/repository"

type CommentService struct {
	userRepo  repository.UserRepository
	photoRepo repository.PhotoRepository
}
