package comment

import (
	"context"
	"postMaker/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter Filter) ([]entity.Comment, int, error)
	GetAllByPostId(ctx context.Context, filter Filter, postId int) ([]entity.Comment, int, error)
	GetById(ctx context.Context, id int) (entity.Comment, error)
	GetByPostId(ctx context.Context, postId int) (entity.Comment, error)
	Create(ctx context.Context, data Create) (entity.Comment, error)
	Update(ctx context.Context, data Update) (entity.Comment, error)
	Delete(ctx context.Context, id int) error
}
