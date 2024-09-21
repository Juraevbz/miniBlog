package repository

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"time"

	"gorm.io/gorm"
)

func (r *Repository) CreatePost(ctx context.Context, p models.Post) (*models.Post, error) {
	err := r.DB.WithContext(ctx).Create(&p).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *Repository) GetPostByID(ctx context.Context, postID int) (*models.Post, error) {
	post := models.Post{}
	err := r.DB.WithContext(ctx).Where("deleted_at IS NULL").First(&post, postID).Error
	if err != nil {
		return nil, handleError(err)
	}

	return &post, nil
}

func (r *Repository) GetPosts(ctx context.Context) ([]*models.Post, error) {
	posts := []*models.Post{}
	err := r.DB.WithContext(ctx).Where("deleted_at IS NULL").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *Repository) UpdatePost(ctx context.Context, postID int, p models.Post) (*models.Post, error) {
	err := r.DB.WithContext(ctx).Model(&models.Post{}).Where("id = ?", postID).Updates(&p).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *Repository) DeletePost(ctx context.Context, postID int) error {
	tNow := time.Now()
	post := models.Post{
		DeletedAt: &tNow,
	}
	
	err := r.DB.WithContext(ctx).Where("id = ?", postID).Updates(&post).Error
	if err != nil {
		return err
	}

	return nil
}

func handleError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}
	return err
}