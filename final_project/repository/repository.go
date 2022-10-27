package repository

import (
	"final_project/datasource"
)

type Repository struct {
	UserRepository        UserRepository
	PhotoRepository       PhotoRepository
	CommentRepository     CommentRepository
	SocialMediaRepository SocialMediaRepository
}

func New(db *datasource.Database) Repository {
	return Repository{
		UserRepository: UserRepository{
			database: db,
		},
	}
}
