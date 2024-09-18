package repository

import (
	"context"
	"mini_blog/internal/models"
)

func (r *Repository) CreatePostRepo(ctx context.Context, p models.Post) (*models.Post, error) {
	err := r.DB.WithContext(ctx).Create(&p).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *Repository) GetPostByIDRepo(ctx context.Context, postID int) (*models.Post, error) {
	post := models.Post{}
	err := r.DB.WithContext(ctx).First(&post, postID).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *Repository) GetPostsRepo(ctx context.Context) ([]*models.Post, error) {
	posts := []*models.Post{}
	err := r.DB.WithContext(ctx).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}
 
func (r *Repository) UpdatePostRepo(ctx context.Context, postID int, p models.Post) (*models.Post, error) {
	err := r.DB.WithContext(ctx).Model(&models.Post{}).Where("id = ?", postID).Updates(&p).Error
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *Repository) DeletePostRepo(ctx context.Context, postID int) error {
	err := r.DB.WithContext(ctx).Delete(&models.Post{}, postID).Error
	if err != nil {
		return err
	}

	return nil
}
