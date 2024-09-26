package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"mini_blog/internal/repository"
	"time"

	"github.com/rs/zerolog"
)

type Service struct {
	repo   *repository.Repository
	logger zerolog.Logger
}

func NewService(repo *repository.Repository, logger zerolog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) CreatePost(ctx context.Context, p models.Post) (*models.Post, error) {
	if err := p.Validate(); err != nil {
		s.logger.Error().Err(err).Msg("failed to validate post")
		return nil, errors.Join(errs.ErrValidationFailed, err)
	}

	post, err := s.repo.CreatePost(ctx, p)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to create post")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return post, nil
}

func (s *Service) GetPostByID(ctx context.Context, postID int, userID int) (*models.Post, error) {
	post, err := s.repo.GetPostByID(ctx, postID)
	if err != nil {
		s.logger.Error().Err(err).Int("postID", postID).Msg("error getting post by id")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	comments, err := s.repo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		s.logger.Error().Err(err).Int("postID", postID).Msg("error getting comments by post id")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	likes, err := s.repo.GetLikesByPostID(ctx, postID)
	if err != nil {
		s.logger.Error().Err(err).Int("postID", postID).Msg("err getting likes by post id")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	post.Comments = append(post.Comments, comments...)
	post.Likes = append(post.Likes, likes...)

	return post, nil
}

func (s *Service) GetPosts(ctx context.Context, userID int) ([]*models.PostList, error) {
	posts, err := s.repo.GetPosts(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to get posts form database")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	postList := []*models.PostList{}
	for _, post := range posts {
		countComments, err := s.repo.CountComments(ctx, post.ID)
		if err != nil {
			s.logger.Error().Err(err).Int("postID", post.ID).Msg("failed to count comments for post")
			return nil, errors.Join(errs.ErrInternalDatabaseError, err)
		}

		countLikes, err := s.repo.CountLikes(ctx, post.ID)
		if err != nil {
			s.logger.Error().Err(err).Int("postID", post.ID).Msg("failed to count likes for post")
			return nil, errors.Join(errs.ErrInternalDatabaseError, err)
		}

		pl := &models.PostList{
			PostID:   post.ID,
			Title:    post.Title,
			Comments: countComments,
			Likes:    countLikes,
		}
		postList = append(postList, pl)
	}

	return postList, nil
}

func (s *Service) UpdatePost(ctx context.Context, postID int, p models.Post) (*models.Post, error) {
	post, err := s.repo.UpdatePost(ctx, postID, p)
	if err != nil {
		s.logger.Error().Err(err).Int("postID", postID).Msg("failed to update post in database")
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return post, nil
}

func (s *Service) DeletePost(ctx context.Context, postID int, userID int) error {
	tNow := time.Now()
	toUpdate := models.Post{
		ID:        postID,
		DeletedAt: &tNow,
	}

	err := s.repo.DeletePost(ctx, toUpdate)
	if err != nil {
		s.logger.Error().Err(err).Int("postID", postID).Msg("failed to delete post")
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return nil
}
