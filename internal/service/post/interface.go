package post

import (
	"context"
	"postMaker/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter Filter) ([]entity.Post, int, error)
	GetById(ctx context.Context, id int) (entity.Post, error)
	Create(ctx context.Context, data Create) (entity.Post, error)
	Update(ctx context.Context, data Update) (entity.Post, error)
	Delete(ctx context.Context, id int) error
}
