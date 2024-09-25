package repository

import (
	"context"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"

	"gorm.io/gorm"
)

func (r *Repository) CreateUser(ctx context.Context, u models.User) error {
	err := r.DB.WithContext(ctx).Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user *models.User
	err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return nil, errs.ErrUserNotFound
		}
		return nil, err.Error
	}
	return user, nil
}