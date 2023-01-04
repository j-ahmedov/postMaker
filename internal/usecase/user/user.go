package user

import (
	"context"
	"errors"
	"postMaker/internal/service/auth"
	"postMaker/internal/service/user"
)

type UseCase struct {
	user User
	file File
}

func NewUseCase(user User, file File) *UseCase {
	return &UseCase{user: user, file: file}
}

func (cu UseCase) GetUserList(ctx context.Context, filter user.Filter) ([]user.List, int, error) {
	data, count, err := cu.user.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []user.List

	for _, r := range data {
		var detail user.List

		detail.Id = r.Id
		detail.Username = r.Username
		if r.Avatar != nil {
			detail.Avatar = *r.Avatar
		}

		list = append(list, detail)
	}

	return list, count, err
}

func (cu UseCase) GetUserById(ctx context.Context, id int) (user.Detail, error) {
	data, err := cu.user.GetById(ctx, id)
	if err != nil {
		return user.Detail{}, err
	}

	var detail user.Detail

	detail.Id = data.Id
	detail.Username = data.Username
	if data.Avatar != nil {
		detail.Avatar = *data.Avatar
	}

	return detail, err
}

func (cu UseCase) CreateUser(ctx context.Context, createForm user.CreateForm) (user.Detail, error) {

	var avatarLink string
	var err error

	_, err1 := cu.user.GetByUsername(ctx, createForm.Username)
	if err1 == nil {
		err2 := errors.New("user with such username already exists")
		return user.Detail{}, err2
	}

	var create user.Create

	if createForm.Avatar != nil {
		avatarLink, err = cu.file.Upload(ctx, createForm.Avatar, "avatars")
		if err != nil {
			return user.Detail{}, err
		}
		create.Avatar = &avatarLink
	}

	hashPassword, err := auth.HashPassword(createForm.Password)
	if err != nil {
		return user.Detail{}, err
	}

	create.Username = createForm.Username
	create.Password = hashPassword

	data, err := cu.user.Create(ctx, create)
	if err != nil {
		return user.Detail{}, err
	}

	var detail user.Detail

	detail.Id = data.Id
	detail.Username = data.Username
	if data.Avatar != nil {
		detail.Avatar = *data.Avatar
	}

	return detail, err
}

func (cu UseCase) UpdateUser(ctx context.Context, updateForm user.UpdateForm) (user.Detail, error) {

	userData, err := cu.user.GetById(ctx, updateForm.Id)
	if err != nil {
		return user.Detail{}, err
	}

	if updateForm.Avatar != nil && userData.Avatar != nil {
		err = cu.file.Delete(ctx, *userData.Avatar)
		if err != nil {
			return user.Detail{}, err
		}

	}

	avatarLink, err := cu.file.Upload(ctx, updateForm.Avatar, "avatars")
	if err != nil {
		return user.Detail{}, err
	}

	hashPassword, err := auth.HashPassword(updateForm.Password)
	if err != nil {
		return user.Detail{}, err
	}

	var update user.Update

	update.Id = updateForm.Id
	update.Username = updateForm.Username
	update.Password = hashPassword
	update.Avatar = &avatarLink

	data, err := cu.user.Update(ctx, update)
	if err != nil {
		return user.Detail{}, err
	}

	var detail user.Detail

	detail.Id = data.Id
	detail.Username = data.Username
	detail.Avatar = *data.Avatar

	return detail, err
}

func (cu UseCase) DeleteUser(ctx context.Context, id int) error {
	return cu.user.Delete(ctx, id)
}
