package post_like

import (
	"context"
	"postMaker/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter Filter) ([]entity.PostLike, int, error)
	GetById(ctx context.Context, id int) (entity.PostLike, error)
	Create(ctx context.Context, data Create) (entity.PostLike, error)
	Update(ctx context.Context, data Update) (entity.PostLike, error)
	Delete(ctx context.Context, id int) error
}
