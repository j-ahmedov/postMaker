package post

import (
	"context"
	"postMaker/internal/entity"
	"postMaker/internal/service/post"
)

type UseCase struct {
	post     Post
	postLike PostLike
	postFile PostFile
	file     File
	user     User
}

func NewUseCase(post Post, postLike PostLike, postFile PostFile, file File, user User) *UseCase {
	return &UseCase{
		post:     post,
		postLike: postLike,
		postFile: postFile,
		file:     file,
		user:     user,
	}
}

func (cu UseCase) GetPostList(ctx context.Context, filter post.Filter) ([]post.List, int, error) {
	data, count, err := cu.post.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []post.List

	for _, r := range data {
		var detail post.List

		detail.Id = r.Id
		detail.Description = r.Description
		detail.UserId = r.UserId
		detail.CreatedAt = r.CreatedAt.Format("02-01-2006")

		list = append(list, detail)
	}

	return list, count, err
}

func (cu UseCase) GetPostById(ctx context.Context, id int, username string) (post.Detail, error) {
	data, err := cu.post.GetById(ctx, id)
	if err != nil {
		return post.Detail{}, err
	}

	userDetail, err := cu.user.GetByUsername(ctx, username)
	if err != nil {
		return post.Detail{}, err
	}

	var detail post.Detail

	detail.Id = data.Id
	detail.Description = data.Description
	detail.User.Username = userDetail.Username
	detail.User.Avatar = *userDetail.Avatar
	detail.CreatedAt = data.CreatedAt.Format("02-01-2006")

	return detail, err
}

func (cu UseCase) CreatePost(ctx context.Context, create post.Create) (entity.Post, error) {

	return cu.post.Create(ctx, create)
}

func (cu UseCase) UpdatePost(ctx context.Context, update post.Update) (entity.Post, error) {
	return cu.post.Update(ctx, update)
}

func (cu UseCase) DeletePost(ctx context.Context, id int) error {
	return cu.post.Delete(ctx, id)
}
