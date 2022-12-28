package comment

import (
	"context"
	"github.com/uptrace/bun"
	"postMaker/internal/entity"
	"postMaker/internal/service/comment"
	"time"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) GetAll(ctx context.Context, filter comment.Filter) ([]entity.Comment, int, error) {
	var list []entity.Comment
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

func (r Repository) GetAllByPostId(ctx context.Context, filter comment.Filter, postId int) ([]entity.Comment, int, error) {
	var list []entity.Comment
	q := r.NewSelect().Model(&list).Where("post_id = ", postId)

	if filter.Limit != nil {
		q.Limit(*filter.Limit)
	}
	if filter.Offset != nil {
		q.Limit(*filter.Offset)
	}

	count, err := q.ScanAndCount(ctx)

	return list, count, err
}

func (r Repository) GetById(ctx context.Context, id int) (entity.Comment, error) {
	var detail entity.Comment

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)

	return detail, err
}

func (r Repository) GetByPostId(ctx context.Context, postId int) (entity.Comment, error) {
	var detail entity.Comment

	err := r.NewSelect().Model(&detail).Where("post_id = ?", postId).Scan(ctx)

	return detail, err
}

func (r Repository) Create(ctx context.Context, data comment.Create) (entity.Comment, error) {
	var detail entity.Comment

	createdDate := time.Now()

	detail.PostId = data.PostId
	detail.UserId = data.UserId
	detail.Text = data.Text
	detail.CreatedAt = createdDate

	_, err := r.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r Repository) Update(ctx context.Context, data comment.Update) (entity.Comment, error) {
	var detail entity.Comment

	err := r.NewSelect().Model(&detail).Where("id = ?", data.Id).Scan(ctx)
	if err != nil {
		return entity.Comment{}, err
	}

	detail.PostId = data.PostId
	detail.UserId = data.UserId
	detail.Text = data.Text

	_, err = r.NewUpdate().Model(&detail).Where("id = ?", detail.Id).Exec(ctx)

	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int) error {

	_, err := r.NewDelete().Table("comments").Where("id = ?", id).Exec(ctx)

	return err
}
