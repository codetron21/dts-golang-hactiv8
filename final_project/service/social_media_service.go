package service

import "final_project/repository"

type SocialMediaService struct {
	userRepo  repository.UserRepository
	photoRepo repository.PhotoRepository
}
