package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"time"
)

func (s *Service) CreateComment(ctx context.Context, c models.Comment) (*models.Comment, error) {
	return s.repo.CreateComment(ctx, c)
}

func (s *Service) GetCommentByID(ctx context.Context, commentID int) (*models.Comment, error) {
	comment, err := s.repo.GetCommentByID(ctx, commentID)
	if err != nil {
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, commentID int, c models.Comment) (*models.Comment, error) {
	return s.repo.UpdateComment(ctx, commentID, c)
}

func (s *Service) DeleteComment(ctx context.Context, commentID int) error {
	tNow := time.Now()
	toUpdate := models.Comment{
		ID:        uint(commentID),
		DeletedAt: &tNow,
	}

	err := s.repo.DeleteComment(ctx, toUpdate)
	if err != nil {
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return nil
}
