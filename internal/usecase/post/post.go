package post

import (
	"context"
	"os"
	"postMaker/internal/service/post"
	"postMaker/internal/service/post_like"
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

func (cu UseCase) GetPostList(ctx context.Context, filter post.Filter, likeFilter post_like.Filter, userId int) ([]post.Detail, int, error) {
	data, count, err := cu.post.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []post.Detail

	for _, r := range data {
		var detail post.Detail

		user, err1 := cu.user.GetById(ctx, r.UserId)
		if err1 != nil {
			return nil, 0, err1
		}

		postFile, err2 := cu.postFile.GetByPostId(ctx, r.Id)
		if err2 != nil {
			return nil, 0, err2
		}

		_, likeCount, err3 := cu.postLike.GetAllByPostId(ctx, likeFilter, r.Id)
		if err3 != nil {
			return nil, 0, err3
		}

		_, err4 := cu.postLike.GetByUserAndPost(ctx, userId, detail.Id)
		if err4 == nil {
			detail.Liked = true
		} else {
			detail.Liked = false
		}

		detail.Id = r.Id
		detail.Description = r.Description
		detail.CreatedAt = r.CreatedAt.Format("02-01-2006")
		detail.Likes = likeCount
		detail.User.Username = user.Username
		if user.Avatar != nil {
			avatarLink := os.Getenv("BASE_URL") + *user.Avatar
			detail.User.Avatar = &avatarLink
		}
		if postFile.Link != nil {
			detail.Files = postFile.Link
		}

		list = append(list, detail)
	}

	return list, count, err
}

func (cu UseCase) GetPostById(ctx context.Context, id int) (post.Detail, error) {
	data, err := cu.post.GetById(ctx, id)
	if err != nil {
		return post.Detail{}, err
	}

	user, err1 := cu.user.GetById(ctx, data.UserId)
	if err != nil {
		return post.Detail{}, err1
	}

	postFile, err2 := cu.postFile.GetByPostId(ctx, data.Id)
	if err2 != nil {
		return post.Detail{}, err2
	}

	var detail post.Detail

	detail.Id = data.Id
	detail.Description = data.Description
	detail.CreatedAt = data.CreatedAt.Format("02-01-2006")
	detail.User.Username = user.Username
	if user.Avatar != nil {
		avatarLink := os.Getenv("BASE_URL") + *user.Avatar
		detail.User.Avatar = &avatarLink
	}
	if postFile.Link != nil {
		detail.Files = postFile.Link
	}

	return detail, err
}

func (cu UseCase) CreatePost(ctx context.Context, create post.Create) (post.Detail, error) {

	data, err := cu.post.Create(ctx, create)
	if err != nil {
		return post.Detail{}, err
	}

	user, err := cu.user.GetById(ctx, data.UserId)
	if err != nil {
		return post.Detail{}, err
	}

	var detail post.Detail

	detail.Id = data.Id
	detail.Description = data.Description
	detail.CreatedAt = data.CreatedAt.Format("02-01-2006")
	detail.User.Username = user.Username
	if user.Avatar != nil {
		avatarLink := os.Getenv("BASE_URL") + *user.Avatar
		detail.User.Avatar = &avatarLink
	}

	return detail, err

}

func (cu UseCase) UpdatePost(ctx context.Context, update post.Update) (post.Detail, error) {
	data, err := cu.post.Update(ctx, update)
	if err != nil {
		return post.Detail{}, err
	}

	user, err := cu.user.GetById(ctx, data.UserId)
	if err != nil {
		return post.Detail{}, err
	}

	var detail post.Detail

	detail.Id = data.Id
	detail.Description = data.Description
	detail.CreatedAt = data.CreatedAt.Format("02-01-2006")
	detail.User.Username = user.Username
	if user.Avatar != nil {
		avatarLink := os.Getenv("BASE_URL") + *user.Avatar
		detail.User.Avatar = &avatarLink
	}

	return detail, err

}

func (cu UseCase) DeletePost(ctx context.Context, id int) error {
	return cu.post.Delete(ctx, id)
}
