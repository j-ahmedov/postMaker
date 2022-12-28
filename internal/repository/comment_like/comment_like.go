package comment_like

import (
	"context"
	"github.com/uptrace/bun"
	"postMaker/internal/entity"
	"postMaker/internal/service/comment_like"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) GetAll(ctx context.Context, filter comment_like.Filter) ([]entity.CommentLike, int, error) {
	var list []entity.CommentLike
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

func (r Repository) GetById(ctx context.Context, id int) (entity.CommentLike, error) {
	var detail entity.CommentLike

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)

	return detail, err
}

func (r Repository) Create(ctx context.Context, data comment_like.Create) (entity.CommentLike, error) {
	var detail entity.CommentLike

	detail.UserId = data.UserId
	detail.CommentId = data.CommentId

	_, err := r.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r Repository) Update(ctx context.Context, data comment_like.Update) (entity.CommentLike, error) {
	var detail entity.CommentLike

	err := r.NewSelect().Model(&detail).Where("id = ?", data.Id).Scan(ctx)
	if err != nil {
		return entity.CommentLike{}, err
	}

	detail.UserId = data.UserId
	detail.CommentId = data.CommentId

	_, err = r.NewUpdate().Model(&detail).Where("id = ?", detail.Id).Exec(ctx)

	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int) error {

	_, err := r.NewDelete().Table("comment_likes").Where("id = ?", id).Exec(ctx)

	return err
}
