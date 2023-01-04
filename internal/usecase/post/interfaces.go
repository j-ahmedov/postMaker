package post

import (
	"context"
	"mime/multipart"
	"postMaker/internal/entity"
	"postMaker/internal/service/post"
	"postMaker/internal/service/post_file"
	"postMaker/internal/service/post_like"
)

type Post interface {
	GetAll(ctx context.Context, filter post.Filter) ([]entity.Post, int, error)
	GetById(ctx context.Context, id int) (entity.Post, error)
	Create(ctx context.Context, data post.Create) (entity.Post, error)
	Update(ctx context.Context, data post.Update) (entity.Post, error)
	Delete(ctx context.Context, id int) error
}

type PostLike interface {
	GetAll(ctx context.Context, filter post_like.Filter) ([]entity.PostLike, int, error)
	GetById(ctx context.Context, id int) (entity.PostLike, error)
	Create(ctx context.Context, data post_like.Create) (entity.PostLike, error)
	Update(ctx context.Context, data post_like.Update) (entity.PostLike, error)
	Delete(ctx context.Context, id int) error
}

type PostFile interface {
	GetAll(ctx context.Context, filter post_file.Filter) ([]entity.PostFile, int, error)
	GetById(ctx context.Context, id int) (entity.PostFile, error)
	GetByPostId(ctx context.Context, postId int) (entity.PostFile, error)
	Create(ctx context.Context, data post_file.MCreate) (entity.PostFile, error)
	Update(ctx context.Context, data post_file.MUpdate) (entity.PostFile, error)
	Delete(ctx context.Context, id int) error
}

type File interface {
	Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error)
	Delete(ctx context.Context, dst string) error
	MultipleUpload(ctx context.Context, files []*multipart.FileHeader, folder string) ([]string, error)
}

type User interface {
	GetById(ctx context.Context, id int) (entity.User, error)
	//GetByUsername(ctx context.Context, username string) (entity.User, error)
}
