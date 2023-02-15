package post_like

import (
	"context"
	"github.com/uptrace/bun"
	"postMaker/internal/entity"
	"postMaker/internal/service/post_like"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) GetAll(ctx context.Context, filter post_like.Filter) ([]entity.PostLike, int, error) {
	var list []entity.PostLike
	q := r.NewSelect().Model(&list)

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}
	if filter.Offset != nil {
		q.Limit(*filter.Offset)
	}

	count, err := q.ScanAndCount(ctx)

	return list, count, err
}

func (r Repository) GetById(ctx context.Context, id int) (entity.PostLike, error) {
	var detail entity.PostLike

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)

	return detail, err
}

func (r Repository) GetAllByPostId(ctx context.Context, filter post_like.Filter, postId int) ([]entity.PostLike, int, error) {
	var list []entity.PostLike
	q := r.NewSelect().Model(&list).Where("post_id = ?", postId)

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}
	if filter.Offset != nil {
		q.Limit(*filter.Offset)
	}

	count, err := q.ScanAndCount(ctx)

	return list, count, err
}

func (r Repository) GetByUserAndPost(ctx context.Context, userId, postId int) (entity.PostLike, error) {
	var detail entity.PostLike

	err := r.NewSelect().Model(&detail).Where("user_id = ? AND post_id = ?", userId, postId).Scan(ctx)
	return detail, err
}

func (r Repository) Create(ctx context.Context, data post_like.Create) (entity.PostLike, error) {
	var detail entity.PostLike

	detail.UserId = data.UserId
	detail.PostId = data.PostId

	_, err := r.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int) error {

	_, err := r.NewDelete().Table("post_likes").Where("id = ?", id).Exec(ctx)

	return err
}
