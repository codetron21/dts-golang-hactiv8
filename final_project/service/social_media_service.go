package service

import "final_project/repository"

type SocialMediaService struct {
	userRepo        repository.UserRepository
	socialMediaRepo repository.SocialMediaRepository
}
