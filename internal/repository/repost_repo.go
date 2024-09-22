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
	err := r.DB.WithContext(ctx).Preload("Post").First(&repost, repostID).Error
	if err != nil {
		return nil, err
	}

	return &repost, nil
}

func (r *Repository) DeleteRepost(ctx context.Context, repos models.Repost) error {
	err := r.DB.WithContext(ctx).Where("id = ?", repos.ID).Updates(&r).Error
	if err != nil {
		return err
	}

	return nil
}
