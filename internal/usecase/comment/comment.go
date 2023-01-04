package comment

import (
	"context"
	"postMaker/internal/entity"
	"postMaker/internal/service/comment"
)

type UseCase struct {
	comment     Comment
	commentLike CommentLike
	user        User
}

func NewUseCase(comment Comment, commentLike CommentLike, user User) *UseCase {
	return &UseCase{
		comment:     comment,
		commentLike: commentLike,
		user:        user,
	}
}

func (cu UseCase) GetCommentList(ctx context.Context, filter comment.Filter) ([]comment.List, int, error) {
	data, count, err := cu.comment.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []comment.List

	for _, r := range data {
		var detail comment.List

		detail.Id = r.Id
		detail.PostId = r.PostId
		detail.UserId = r.UserId
		detail.Text = r.Text
		detail.CreatedAt = r.CreatedAt.Format("02-01-2006")

		list = append(list, detail)
	}

	return list, count, err
}

func (cu UseCase) GetCommentListByPostId(ctx context.Context, filter comment.Filter, postId int) ([]comment.PostDetail, int, error) {
	data, count, err := cu.comment.GetAllByPostId(ctx, filter, postId)

	if err != nil {
		return nil, 0, err
	}

	var list []comment.PostDetail

	for _, r := range data {
		var detail comment.PostDetail

		user, err1 := cu.user.GetById(ctx, r.UserId)
		if err1 != nil {
			return nil, 0, err1
		}

		detail.Id = r.Id
		detail.Text = r.Text
		detail.User.Username = user.Username
		detail.User.Avatar = *user.Avatar
		detail.CreatedAt = r.CreatedAt.Format("02-01-2006")

		list = append(list, detail)

	}

	return list, count, err
}

func (cu UseCase) GetCommentById(ctx context.Context, id int) (comment.PostDetail, error) {
	data, err := cu.comment.GetById(ctx, id)
	if err != nil {
		return comment.PostDetail{}, err
	}

	user, err := cu.user.GetById(ctx, data.UserId)
	if err != nil {
		return comment.PostDetail{}, err
	}

	var detail comment.PostDetail

	detail.Id = data.Id
	detail.Text = data.Text
	detail.User.Username = user.Username
	detail.User.Avatar = *user.Avatar
	detail.CreatedAt = data.CreatedAt.Format("02-01-2006")

	return detail, err
}

func (cu UseCase) GetCommentByPostId(ctx context.Context, postId int) (comment.PostDetail, error) {
	data, err := cu.comment.GetByPostId(ctx, postId)
	if err != nil {
		return comment.PostDetail{}, err
	}

	user, err := cu.user.GetById(ctx, data.UserId)
	if err != nil {
		return comment.PostDetail{}, err
	}

	var detail comment.PostDetail

	detail.Id = data.Id
	detail.Text = data.Text
	detail.User.Username = user.Username
	detail.User.Avatar = *user.Avatar
	detail.CreatedAt = data.CreatedAt.Format("02-01-2006")

	return detail, err
}

func (cu UseCase) CreateComment(ctx context.Context, create comment.Create) (comment.PostDetail, error) {
	data, err := cu.comment.Create(ctx, create)
	if err != nil {
		return comment.PostDetail{}, err
	}

	user, err := cu.user.GetById(ctx, data.UserId)
	if err != nil {
		return comment.PostDetail{}, err
	}

	var detail comment.PostDetail

	detail.Id = data.Id
	detail.Text = data.Text
	detail.User.Username = user.Username
	detail.User.Avatar = *user.Avatar
	detail.CreatedAt = data.CreatedAt.Format("02-01-2006")

	return detail, err
}

func (cu UseCase) UpdateComment(ctx context.Context, update comment.Update) (entity.Comment, error) {
	data, err := cu.comment.Update(ctx, update)

	return data, err
}

func (cu UseCase) DeleteComment(ctx context.Context, id int) error {
	return cu.comment.Delete(ctx, id)
}
