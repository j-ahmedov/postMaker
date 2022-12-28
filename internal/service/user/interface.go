package user

import (
	"context"
	"postMaker/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter Filter) ([]entity.User, int, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	Create(ctx context.Context, data Create) (entity.User, error)
	Update(ctx context.Context, data Update) (entity.User, error)
	Delete(ctx context.Context, id int) error
}
