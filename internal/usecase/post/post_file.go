package post

import (
	"context"
	"postMaker/internal/entity"
	"postMaker/internal/service/post_file"
	"strings"
)

func (cu UseCase) GetPostFileList(ctx context.Context, filter post_file.Filter) ([]post_file.List, int, error) {
	data, count, err := cu.postFile.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []post_file.List

	for _, r := range data {
		var detail post_file.List

		detail.Id = r.Id
		detail.PostId = r.PostId
		detail.Link = r.Link

		list = append(list, detail)
	}

	return list, count, err
}

func (cu UseCase) GetPostFileById(ctx context.Context, id int) (entity.PostFile, error) {
	data, err := cu.postFile.GetById(ctx, id)
	if err != nil {
		return entity.PostFile{}, err
	}

	return data, err
}

func (cu UseCase) CreatePostFile(ctx context.Context, create post_file.FileCreate) (entity.PostFile, error) {

	postLink, err := cu.file.Upload(ctx, create.File, "post-file")
	if err != nil {
		return entity.PostFile{}, err
	}

	var detail post_file.Create

	detail.PostId = create.PostId
	detail.Link = postLink

	data, err := cu.postFile.Create(ctx, detail)

	return data, err
}

func (cu UseCase) CreateMultipleFiles(ctx context.Context, multiCreate post_file.MultipleCreate) (entity.PostFile, error) {
	if (multiCreate.Files != nil) && (len(multiCreate.Files) > 0) {
		postLinks, err := cu.file.MultipleUpload(ctx, multiCreate.Files, "file-folder")
		if err != nil {
			return entity.PostFile{}, err
		}

		postLink := strings.Join(postLinks, "")
		var detail post_file.Create

		detail.PostId = multiCreate.PostId
		detail.Link = postLink

		data, err := cu.postFile.Create(ctx, detail)

		return data, err

	}

	return entity.PostFile{}, nil
}

func (cu UseCase) UpdatePostFile(ctx context.Context, update post_file.FileUpdate) (entity.PostFile, error) {

	err := cu.DeletePostFile(ctx, update.Id)
	if err != nil {
		return entity.PostFile{}, err
	}

	postLink, err := cu.file.Upload(ctx, update.File, "file-folder")

	var detail post_file.Update

	detail.Id = update.Id
	detail.PostId = update.PostId
	detail.Link = postLink

	data, err := cu.postFile.Update(ctx, detail)

	return data, err
}

func (cu UseCase) DeletePostFile(ctx context.Context, id int) error {
	data, err := cu.postFile.GetById(ctx, id)
	if err != nil {
		return err
	}

	err = cu.file.Delete(ctx, data.Link)
	if err != nil {
		return err
	}

	return cu.postFile.Delete(ctx, id)
}
