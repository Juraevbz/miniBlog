package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"time"
)

func (s *Service) CreateRepost(ctx context.Context, postID int) (*models.Repost, error) {
	post, err := s.repo.GetPostByID(ctx, postID)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to get post by id")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	countComments, err := s.repo.CountComments(ctx, int(post.ID))
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		s.logger.Error().Err(err).Msg("failed to count comments")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	countLikes, err := s.repo.CountLikes(ctx, int(post.ID))
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		s.logger.Error().Err(err).Msg("failed to count likes")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	toCreate := models.Repost{
		PostID:   postID,
		Title:    post.Title,
		Content:  post.Content,
		ImageURL: post.ImageURL,
		Comments: countComments,
		Likes:    countLikes,
	}

	repost, err := s.repo.CreateRepost(ctx, toCreate)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to create repost")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return repost, nil
}

func (s *Service) GetRepostByID(ctx context.Context, repostID int) (*models.Repost, error) {
	repost, err := s.repo.GetRepostByID(ctx, repostID)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to get repost by id")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return repost, nil
}

func (s *Service) DeleteRepost(ctx context.Context, repostID int) error {
	tNow := time.Now()
	toDelete := models.Repost{
		DeletedAt: &tNow,
	}

	err := s.repo.DeleteRepost(ctx, repostID, toDelete)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to delete repost")
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return nil
}
