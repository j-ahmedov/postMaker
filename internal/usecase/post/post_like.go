package post

import (
	"context"
	"postMaker/internal/entity"
	"postMaker/internal/service/post_like"
)

// Post Like

func (cu UseCase) GetPostLikeList(ctx context.Context, filter post_like.Filter) ([]post_like.List, int, error) {
	data, count, err := cu.postLike.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []post_like.List

	for _, r := range data {
		var detail post_like.List

		detail.Id = r.Id
		detail.UserId = r.UserId
		detail.PostId = r.PostId

		list = append(list, detail)
	}

	return list, count, err
}

func (cu UseCase) GetPostLikeById(ctx context.Context, id int) (entity.PostLike, error) {
	data, err := cu.postLike.GetById(ctx, id)
	if err != nil {
		return entity.PostLike{}, err
	}

	return data, err
}

func (cu UseCase) CreatePostLike(ctx context.Context, create post_like.Create) (entity.PostLike, error) {
	data, err := cu.postLike.Create(ctx, create)

	return data, err
}

func (cu UseCase) UpdatePostLike(ctx context.Context, update post_like.Update) (entity.PostLike, error) {
	data, err := cu.postLike.Update(ctx, update)

	return data, err
}

func (cu UseCase) DeletePostLike(ctx context.Context, id int) error {
	return cu.postLike.Delete(ctx, id)
}
