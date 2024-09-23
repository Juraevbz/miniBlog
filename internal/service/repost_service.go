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
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	countComments, err := s.repo.CountComments(ctx, int(post.ID))
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	countLikes, err := s.repo.CountLikes(ctx, int(post.ID))
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	toCreate := models.Repost{
		PostID:   uint(postID),
		Title:    post.Title,
		Content:  post.Content,
		ImageURL: post.ImageURL,
		Comments: countComments,
		Likes:    countLikes,
	}

	repost, err := s.repo.CreateRepost(ctx, toCreate)
	if err != nil {
		return nil, err
	}

	return repost, nil
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
	toDelete := models.Repost{
		DeletedAt: &tNow,
	}

	err := s.repo.DeleteRepost(ctx, repostID, toDelete)
	if err != nil {
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return nil
}
