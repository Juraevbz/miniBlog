package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"time"
)

func (s *Service) CreateRepost(ctx context.Context, r models.Repost) (*models.Repost, error) {
	return s.repo.CreateRepost(ctx, r)
}

func (s *Service) GetRepostByID(ctx context.Context, repostID int) (*models.Repost, error) {
	repost, err := s.repo.GetRepostByID(ctx, repostID)
	if err != nil {
		return nil, err
	}
	return repost, nil
}

func (s *Service) DeleteRepost(ctx context.Context, repostID int) error {
	tNow := time.Now()
	toUpdate := models.Repost{
		ID:        uint(repostID),
		DeletedAt: &tNow,
	}

	err := s.repo.DeleteRepost(ctx, toUpdate)
	if err != nil {
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return nil
}
