package post_file

import (
	"context"
	"postMaker/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter Filter) ([]entity.PostFile, int, error)
	GetById(ctx context.Context, id int) (entity.PostFile, error)
	Create(ctx context.Context, data Create) (entity.PostFile, error)
	Update(ctx context.Context, data Update) (entity.PostFile, error)
	Delete(ctx context.Context, id int) error
}
