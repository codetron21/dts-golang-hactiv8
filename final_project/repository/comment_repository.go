package repository

import (
	"errors"
	"final_project/datasource"
	"final_project/model"
	"fmt"

	"gorm.io/gorm"
)

type CommentRepository struct {
	database *datasource.Database
}

func (repo CommentRepository) CreateComment(data *model.Comment) error {
	err := repo.database.DB.Save(&data).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error create comment")
	}

	return nil
}

func (repo CommentRepository) GetComments() ([]model.Comment, error) {
	comments := []model.Comment{}
	err := repo.database.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Email")
		}).
		Preload("Photo", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Title", "Caption", "PhotoUrl", "UserID")
		}).
		Find(&comments).Error

	if err != nil {
		fmt.Println(err)
		return comments, errors.New("error get comments")
	}

	return comments, nil
}

func (repo CommentRepository) UpdateCommentById(data *model.Comment) error {
	var comment model.Comment
	db := repo.database.DB
	err := db.Where("id = ?", data.ID).First(&comment).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("comment not found")
	}

	data.PhotoID = comment.PhotoID
	err = db.Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error update comment")
	}

	return nil
}

func (repo CommentRepository) DeleteCommentById(commentId int) error {
	db := repo.database.DB

	err := db.Where("id = ?", commentId).First(&model.Comment{}).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("comment not found")
	}

	err = db.Where("id = ?", commentId).Delete(&model.Comment{}).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("error delete comment")
	}

	return nil
}
