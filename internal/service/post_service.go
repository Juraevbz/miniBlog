package service

import (
	"context"
	"errors"
	"mini_blog/internal/errs"
	"mini_blog/internal/models"
	"mini_blog/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreatePost(ctx context.Context, p models.Post) (post *models.Post, err error) {
	return s.repo.CreatePost(ctx, p)
}

func (s *Service) GetPostByID(ctx context.Context, postID int) (*models.Post, error) {
	post, err := s.repo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	comments, err := s.repo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	likes, err := s.repo.GetLikesByPostID(ctx, postID)
	if err != nil {
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	post.Comments = append(post.Comments, comments...)
	post.Likes = append(post.Likes, likes...)

	return post, nil
}

func (s *Service) GetPosts(ctx context.Context) ([]*models.PostList, error) {
	posts, err := s.repo.GetPosts(ctx)
	if err != nil {
		return nil, errors.Join(errs.ErrInternalDatabaseError, err)
	}

	postList := []*models.PostList{}
	for _, post := range posts {
		countComments, err := s.repo.CountComments(ctx, int(post.ID))
		if err != nil {
			return nil, errors.Join(errs.ErrInternalDatabaseError, err)
		}

		countLikes, err := s.repo.CountLikes(ctx, int(post.ID))
		if err != nil {
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
	return s.repo.UpdatePost(ctx, postID, p)
}

func (s *Service) DeletePost(ctx context.Context, postID int) error {
	err := s.repo.DeletePost(ctx, postID)
	if err != nil {
		return errors.Join(errs.ErrInternalDatabaseError, err)
	}

	return nil
}
