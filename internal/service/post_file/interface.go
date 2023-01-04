package post_file

import (
	"context"
	"postMaker/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter Filter) ([]entity.PostFile, int, error)
	GetById(ctx context.Context, id int) (entity.PostFile, error)
	GetByPostId(ctx context.Context, postId int) (entity.PostFile, error)
	Create(ctx context.Context, data MCreate) (entity.PostFile, error)
	Update(ctx context.Context, data MUpdate) (entity.PostFile, error)
	Delete(ctx context.Context, id int) error
}
