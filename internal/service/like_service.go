package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
)

func (s *Service) CreateLike(ctx context.Context, l models.Like) (*models.Like, error) {
	return s.repo.CreateLike(ctx, l)
}

func (s *Service) GetLikeByID(ctx context.Context, likeID int) (*models.Like, error) {
	like, err := s.repo.GetLikeByID(ctx, likeID)
	if err != nil {
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return like, nil
}
