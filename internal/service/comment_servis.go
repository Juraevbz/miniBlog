package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
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
