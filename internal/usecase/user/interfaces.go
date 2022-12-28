package user

import (
	"context"
	"mime/multipart"
	"postMaker/internal/entity"
	"postMaker/internal/service/user"
)

type User interface {
	GetAll(ctx context.Context, filter user.Filter) ([]entity.User, int, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	Create(ctx context.Context, data user.Create) (entity.User, error)
	Update(ctx context.Context, data user.Update) (entity.User, error)
	Delete(ctx context.Context, id int) error
}

type File interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error)
	Delete(ctx context.Context, dst string) error
}
