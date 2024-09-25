package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"time"
)

func (s *Service) CreateComment(ctx context.Context, c models.Comment) (*models.Comment, error) {
	if err := c.Validate(); err != nil {
		s.logger.Error().Err(err).Msg("failed to validate comment")
		return nil, errors.Join(errs.ErrValidationFailed, err)
	}

	comment, err := s.repo.CreateComment(ctx, c)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to create comment")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return comment, nil
}

func (s *Service) GetCommentByID(ctx context.Context, commentID int) (*models.Comment, error) {
	comment, err := s.repo.GetCommentByID(ctx, commentID)
	if err != nil {
		s.logger.Error().Err(err).Int("commentID", commentID).Msg("failed to get comment by id")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, commentID int, c models.Comment) (*models.Comment, error) {
	comment, err := s.repo.UpdateComment(ctx, commentID, c)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to update comment")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return comment, nil
}

func (s *Service) DeleteComment(ctx context.Context, commentID int) error {
	tNow := time.Now()
	toUpdate := models.Comment{
		ID:        commentID,
		DeletedAt: &tNow,
	}

	err := s.repo.DeleteComment(ctx, toUpdate)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to delete comment")
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return nil
}
