package post_file

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

func (s Service) GetAll(ctx context.Context, filter Filter) ([]entity.PostFile, int, error) {
	return s.repo.GetAll(ctx, filter)
}

func (s Service) GetById(ctx context.Context, id int) (entity.PostFile, error) {
	return s.repo.GetById(ctx, id)
}

func (s Service) Create(ctx context.Context, data Create) (entity.PostFile, error) {
	return s.repo.Create(ctx, data)
}

func (s Service) Update(ctx context.Context, data Update) (entity.PostFile, error) {
	return s.repo.Update(ctx, data)
}

func (s Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
