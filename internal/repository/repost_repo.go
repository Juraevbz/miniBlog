package repository

import (
	"context"
	"mini_blog/internal/models"
)

func (r *Repository) CreateRepost(ctx context.Context, repost models.Repost) (*models.Repost, error) {
	err := r.DB.WithContext(ctx).Create(&repost).Error
	if err != nil {
		return nil, err
	}

	return &repost, nil
}

func (r *Repository) GetRepostByID(ctx context.Context, repostID int) (*models.Repost, error) {
	var repost models.Repost
	err := r.DB.WithContext(ctx).First(&repost, repostID).Error
	if err != nil {
		return nil, err
	}

	return &repost, nil
}

func (r *Repository) DeleteRepost(ctx context.Context, repostID int, repost models.Repost) error {
	err := r.DB.WithContext(ctx).Where("id = ?", repostID).Updates(&repost).Error
	if err != nil {
		return err
	}

	return nil
}
