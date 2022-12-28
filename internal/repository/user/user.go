package user

import (
	"context"
	"github.com/uptrace/bun"
	"postMaker/internal/entity"
	"postMaker/internal/service/user"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r Repository) GetAll(ctx context.Context, filter user.Filter) ([]entity.User, int, error) {
	var list []entity.User
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

func (r Repository) GetById(ctx context.Context, id int) (entity.User, error) {
	var detail entity.User

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)

	return detail, err
}

func (r Repository) Create(ctx context.Context, data user.Create) (entity.User, error) {
	var detail entity.User

	detail.Username = data.Username
	detail.Password = data.Password
	detail.Avatar = data.Avatar

	_, err := r.NewInsert().Model(&detail).Exec(ctx)

	return detail, err
}

func (r Repository) Update(ctx context.Context, data user.Update) (entity.User, error) {
	var detail entity.User

	err := r.NewSelect().Model(&detail).Where("id = ?", data.Id).Scan(ctx)
	if err != nil {
		return entity.User{}, err
	}

	detail.Id = data.Id
	detail.Username = data.Username
	detail.Password = data.Password
	detail.Avatar = &data.Avatar

	_, err = r.NewUpdate().Model(&detail).Where("id = ?", data.Id).Exec(ctx)

	return detail, err
}

func (r Repository) Delete(ctx context.Context, id int) error {

	_, err := r.NewDelete().Table("users").Where("id = ?", id).Exec(ctx)

	return err
}

func (r Repository) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	var detail entity.User

	err := r.NewSelect().Model(&detail).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return entity.User{}, err
	}

	return detail, err
}
