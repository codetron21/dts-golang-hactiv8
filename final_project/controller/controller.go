package controller

import "final_project/service"

type Controller struct {
	UserController        UserController
	CommentController     CommentController
	PhotoController       PhotoController
	SocialMediaController SocialMediaController
}

func New(service *service.Service) Controller {
	return Controller{
		UserController: UserController{
			service: service.UserService,
		},
		PhotoController: PhotoController{
			service: service.PhotoService,
		},
		CommentController: CommentController{
			service: service.CommentService,
		},
		SocialMediaController: SocialMediaController{
			service: service.SocialMediaService,
		},
	}
}
