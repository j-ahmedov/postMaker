package comment

import (
	"context"
	"postMaker/internal/entity"
	"postMaker/internal/service/comment_like"
)

// Comment Like

func (cu UseCase) GetCommentLikeList(ctx context.Context, filter comment_like.Filter) ([]comment_like.List, int, error) {
	data, count, err := cu.commentLike.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []comment_like.List

	for _, r := range data {
		var detail comment_like.List

		detail.Id = r.Id
		detail.UserId = r.UserId
		detail.CommentId = r.CommentId

		list = append(list, detail)
	}

	return list, count, err
}

func (cu UseCase) GetCommentLikeById(ctx context.Context, id int) (entity.CommentLike, error) {
	data, err := cu.commentLike.GetById(ctx, id)
	if err != nil {
		return entity.CommentLike{}, err
	}

	return data, err
}

func (cu UseCase) CreateCommentLike(ctx context.Context, create comment_like.Create) (entity.CommentLike, error) {
	data, err := cu.commentLike.Create(ctx, create)

	return data, err
}

func (cu UseCase) UpdateCommentLike(ctx context.Context, update comment_like.Update) (entity.CommentLike, error) {
	data, err := cu.commentLike.Update(ctx, update)

	return data, err
}

func (cu UseCase) DeleteCommentLike(ctx context.Context, id int) error {
	return cu.commentLike.Delete(ctx, id)
}
