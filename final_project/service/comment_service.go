package service

import (
	"errors"
	"final_project/model"
	"final_project/repository"
)

type CommentService struct {
	userRepo    repository.UserRepository
	commentRepo repository.CommentRepository
	photoRepo   repository.PhotoRepository
}

func (cs CommentService) GetComments(userId int) ([]model.Comment, error) {
	err := cs.userRepo.FindUserById(userId)
	if err != nil {
		return []model.Comment{}, err
	}

	comments, err := cs.commentRepo.GetComments()
	if err != nil {
		return []model.Comment{}, err
	}

	return comments, nil
}

func (cs CommentService) CreateComment(userId int, comment *model.Comment) error {
	err := cs.userRepo.FindUserById(userId)
	if err != nil {
		return err
	}

	ok := cs.photoRepo.CheckPhotoIsExistByID(comment.PhotoID)
	if !ok {
		return errors.New("photo id not found")
	}

	comment.UserID = userId
	err = cs.commentRepo.CreateComment(comment)
	if err != nil {
		return err
	}

	return nil
}

func (cs CommentService) UpdateCommentById(userId int, commentId int, comment *model.Comment) error {
	err := cs.userRepo.FindUserById(userId)
	if err != nil {
		return err
	}

	comment.UserID = userId
	comment.ID = commentId
	err = cs.commentRepo.UpdateCommentById(comment)
	if err != nil {
		return err
	}

	return nil
}

func (cs CommentService) DeleteCommentById(userId int, commentId int) error {
	err := cs.userRepo.FindUserById(userId)
	if err != nil {
		return err
	}

	err = cs.commentRepo.DeleteCommentById(commentId)
	if err != nil {
		return err
	}

	return nil
}
