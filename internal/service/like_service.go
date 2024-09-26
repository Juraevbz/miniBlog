package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
)

func (s *Service) CreateLike(ctx context.Context, l models.Like) (*models.Like, error) {
	if err := l.Validate(); err != nil {
		s.logger.Error().Err(err).Msg("failed to validate like")
		return nil, errors.Join(errs.ErrValidationFailed, err)
	}
	
	like, err := s.repo.CreateLike(ctx, l)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to create like")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return like, nil
}

func (s *Service) GetLikeByID(ctx context.Context, likeID int, userID int) (*models.Like, error) {
	like, err := s.repo.GetLikeByID(ctx, likeID)
	if err != nil {
		s.logger.Error().Err(err).Int("likeID", likeID).Msg("failed to get like by id")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return like, nil
}

func (s *Service) DeleteLike(ctx context.Context, likeID int, userID int) error {
	err := s.repo.DeleteLike(ctx, likeID)
	if err != nil {
		s.logger.Error().Err(err).Int("likeID", likeID).Msg("failed to delete like")
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return nil
}
