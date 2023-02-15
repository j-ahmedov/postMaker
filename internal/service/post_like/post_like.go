package post_like

import (
	"context"
	"postMaker/internal/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) GetAll(ctx context.Context, filter Filter) ([]entity.PostLike, int, error) {
	return s.repo.GetAll(ctx, filter)
}

func (s Service) GetById(ctx context.Context, id int) (entity.PostLike, error) {
	return s.repo.GetById(ctx, id)
}

func (s Service) GetAllByPostId(ctx context.Context, filter Filter, postId int) ([]entity.PostLike, int, error) {
	return s.repo.GetAllByPostId(ctx, filter, postId)
}

func (s Service) GetByUserAndPost(ctx context.Context, userId, postId int) (entity.PostLike, error) {
	return s.repo.GetByUserAndPost(ctx, userId, postId)
}

func (s Service) Create(ctx context.Context, data Create) (entity.PostLike, error) {
	return s.repo.Create(ctx, data)
}

func (s Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
