package service

import (
	"context"
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

func (s *Service) CreatePostService(ctx context.Context, p models.Post) (post *models.Post,err error) {
	return s.repo.CreatePostRepo(ctx, p)
}

func (s *Service) GetPostByIDService(ctx context.Context, postID int) (post *models.Post,err error) {
	return s.repo.GetPostByIDRepo(ctx, postID)
}
 
func (s *Service) GetPostsService(ctx context.Context) (posts []*models.Post, err error) {
	return s.repo.GetPostsRepo(ctx)
}

func (s *Service) UpdatePostService(ctx context.Context, postID int, p models.Post) (post *models.Post, err error) {
	return s.repo.UpdatePostRepo(ctx, postID, p)
}

func (s *Service) DeletePostService(ctx context.Context, postID int) (err error) {
	return s.repo.DeletePostRepo(ctx, postID)
}
