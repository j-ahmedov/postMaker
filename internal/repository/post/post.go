package post

import (
	"context"
	"github.com/uptrace/bun"
	"postMaker/internal/entity"
	"postMaker/internal/service/post"
	"time"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) GetAll(ctx context.Context, filter post.Filter) ([]entity.Post, int, error) {
	var list []entity.Post
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

func (r Repository) GetById(ctx context.Context, id int) (entity.Post, error) {
	var detail entity.Post

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)

	return detail, err
}

func (r Repository) Create(ctx context.Context, data post.Create) (entity.Post, error) {
	var detail entity.Post

	currentTime := time.Now()

	detail.Description = data.Description
	detail.UserId = data.UserId
	detail.CreatedAt = currentTime

	_, err := r.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r Repository) Update(ctx context.Context, data post.Update) (entity.Post, error) {
	var detail entity.Post

	err := r.NewSelect().Model(&detail).Where("id = ?", data.Id).Scan(ctx)
	if err != nil {
		return entity.Post{}, err
	}

	detail.Description = data.Description

	_, err = r.NewUpdate().Model(&detail).Where("id = ?", detail.Id).Exec(ctx)

	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int) error {

	_, err := r.NewDelete().Table("posts").Where("id = ?", id).Exec(ctx)

	return err
}
