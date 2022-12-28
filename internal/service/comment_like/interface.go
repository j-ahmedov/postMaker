package comment_like

import (
	"context"
	"postMaker/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter Filter) ([]entity.CommentLike, int, error)
	GetById(ctx context.Context, id int) (entity.CommentLike, error)
	Create(ctx context.Context, data Create) (entity.CommentLike, error)
	Update(ctx context.Context, data Update) (entity.CommentLike, error)
	Delete(ctx context.Context, id int) error
}
