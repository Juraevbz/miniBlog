package repository

import (
	"context"
	"mini_blog/internal/models"
)

func (r *Repository) CreateLike(ctx context.Context, l models.Like) (*models.Like, error) {
	err := r.DB.WithContext(ctx).Create(&l).Error
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *Repository) GetLikeByID(ctx context.Context, likeID int) (*models.Like, error) {
	like := models.Like{}
	err := r.DB.WithContext(ctx).Where("deleted_at IS NULL").First(&like, likeID).Error
	if err != nil {
		return nil, handleError(err)
	}

	return &like, nil
}

func (r *Repository) GetLikesByPostID(ctx context.Context, postID int) ([]*models.Like, error) {
	likes := []*models.Like{}
	err := r.DB.WithContext(ctx).Where("post_id = ? AND deleted_at IS NULL", postID).Find(&likes).Error
	if err != nil {
		return nil, handleError(err)
	}

	return likes, nil
}

func (r *Repository) CountLikes(ctx context.Context, postID int) (int, error) {
	like := models.Like{}
	var count int64
	err := r.DB.WithContext(ctx).Model(&like).Where("post_id = ? AND deleted_at IS NULL", postID).Count(&count).Error
	if err != nil {
		return 0, handleError(err)
	}

	return int(count), nil
}

func (r *Repository) DeleteLike(ctx context.Context, likeID int) error {
	err := r.DB.WithContext(ctx).Delete(&models.Like{}, likeID).Error
	if err != nil {
		return err
	}

	return nil
}