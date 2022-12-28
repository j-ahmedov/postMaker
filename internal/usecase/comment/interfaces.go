package comment

import (
	"context"
	"postMaker/internal/entity"
	"postMaker/internal/service/comment"
	"postMaker/internal/service/comment_like"
)

type Comment interface {
	GetAll(ctx context.Context, filter comment.Filter) ([]entity.Comment, int, error)
	GetAllByPostId(ctx context.Context, filter comment.Filter, postId int) ([]entity.Comment, int, error)
	GetById(ctx context.Context, id int) (entity.Comment, error)
	GetByPostId(ctx context.Context, postId int) (entity.Comment, error)
	Create(ctx context.Context, data comment.Create) (entity.Comment, error)
	Update(ctx context.Context, data comment.Update) (entity.Comment, error)
	Delete(ctx context.Context, id int) error
}

type CommentLike interface {
	GetAll(ctx context.Context, filter comment_like.Filter) ([]entity.CommentLike, int, error)
	GetById(ctx context.Context, id int) (entity.CommentLike, error)
	Create(ctx context.Context, data comment_like.Create) (entity.CommentLike, error)
	Update(ctx context.Context, data comment_like.Update) (entity.CommentLike, error)
	Delete(ctx context.Context, id int) error
}

type User interface {
	GetById(ctx context.Context, id int) (entity.User, error)
}
