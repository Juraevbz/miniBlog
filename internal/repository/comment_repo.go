package repository

import (
	"context"
	"mini_blog/internal/models"
)

func (r *Repository) CreateComment(ctx context.Context, c models.Comment) (*models.Comment, error) {
	err := r.DB.WithContext(ctx).Create(&c).Error
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *Repository) GetCommentByID(ctx context.Context, commentID int) (*models.Comment, error) {
	comment := models.Comment{}
	err := r.DB.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		return nil, handleError(err)
	}

	return &comment, nil
}

func (r *Repository) GetCommentsByPostID(ctx context.Context, postID int) ([]*models.Comment, error) {
	comments := []*models.Comment{}
	err := r.DB.WithContext(ctx).Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		return nil, handleError(err)
	}

	return comments, nil
}

func (r *Repository) CountComments(ctx context.Context, postID int) (int, error) {
	comment := models.Comment{}
	var count int64
	err := r.DB.WithContext(ctx).Model(&comment).Where("post_id = ? AND deleted_at IS NULL", postID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
