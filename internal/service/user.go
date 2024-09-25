package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
)

func (s *Service) CreateUser(ctx context.Context, u models.User) error {
	if err := u.Validate(); err != nil {
		s.logger.Error().Err(err).Msg("validation failed")
		return errors.Join(errs.ErrValidationFailed, err)
	}

	u.PasswordHash = models.GeneratePasswordHash(u.PasswordHash)
	err := s.repo.CreateUser(ctx, u)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to create user")
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}
	return nil
}

func (s *Service) Authenticate(ctx context.Context, username, password string) (*models.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, errors.Join(errs.ErrUserNotFound, err)
	}

	if !user.ComparePassword(password) {
		return nil, errs.ErrUserInvalidPassword
	}

	return user, nil
}
