package post

import (
	"context"
	"errors"
	"postMaker/internal/entity"
	"postMaker/internal/service/post_file"
)

func (cu UseCase) GetPostFileList(ctx context.Context, filter post_file.Filter) ([]entity.PostFile, int, error) {
	data, count, err := cu.postFile.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	var list []entity.PostFile

	for _, r := range data {
		var detail entity.PostFile

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

//func (cu UseCase) CreatePostFile(ctx context.Context, create post_file.FileCreate) (entity.PostFile, error) {
//
//	postLink, err := cu.file.Upload(ctx, create.File, "post-file")
//	if err != nil {
//		return entity.PostFile{}, err
//	}
//
//	var detail post_file.Create
//
//	detail.PostId = create.PostId
//	detail.Link = postLink
//
//	data, err := cu.postFile.Create(ctx, detail)
//
//	return data, err
//}

func (cu UseCase) CreateMultipleFiles(ctx context.Context, multiCreate post_file.MultipleCreate) (entity.PostFile, error) {
	if (multiCreate.Files != nil) && (len(multiCreate.Files) > 0) {
		postLinks, err := cu.file.MultipleUpload(ctx, multiCreate.Files, "file-folder")
		if err != nil {
			return entity.PostFile{}, err
		}

		//postLink := strings.Join(postLinks, "")
		var detail post_file.MCreate

		_, err = cu.postFile.GetByPostId(ctx, multiCreate.PostId)
		if err == nil {
			err1 := errors.New("file with such post id already exists")
			return entity.PostFile{}, err1
		}

		detail.PostId = multiCreate.PostId
		detail.Link = &postLinks

		data, err2 := cu.postFile.Create(ctx, detail)

		return data, err2

	}

	return entity.PostFile{}, nil
}

//func (cu UseCase) UpdatePostFile(ctx context.Context, update post_file.FileUpdate) (entity.PostFile, error) {
//
//	err := cu.DeletePostFile(ctx, update.Id)
//	if err != nil {
//		return entity.PostFile{}, err
//	}
//
//	postLink, err := cu.file.Upload(ctx, update.File, "file-folder")
//
//	var detail post_file.Update
//
//	detail.Id = update.Id
//	detail.PostId = update.PostId
//	detail.Link = postLink
//
//	data, err := cu.postFile.Update(ctx, detail)
//
//	return data, err
//}

func (cu UseCase) UpdateMultiplePostFiles(ctx context.Context, multipleUpdate post_file.MultipleUpdate) (entity.PostFile, error) {

	err := cu.DeletePostFiles(ctx, multipleUpdate.Id)
	if err != nil {
		return entity.PostFile{}, err
	}

	if (multipleUpdate.Files != nil) && (len(multipleUpdate.Files) > 0) {
		postLinks, err := cu.file.MultipleUpload(ctx, multipleUpdate.Files, "file-folder")
		if err != nil {
			return entity.PostFile{}, err
		}

		var detail post_file.MUpdate

		detail.Id = multipleUpdate.Id
		detail.PostId = multipleUpdate.PostId
		detail.Link = &postLinks

		data, err := cu.postFile.Update(ctx, detail)

		return data, err
	}

	return entity.PostFile{}, err
}

//func (cu UseCase) DeletePostFile(ctx context.Context, id int) error {
//	data, err := cu.postFile.GetById(ctx, id)
//	if err != nil {
//		return err
//	}
//
//	err = cu.file.Delete(ctx, data.Link)
//	if err != nil {
//		return err
//	}
//
//	return cu.postFile.Delete(ctx, id)
//}

func (cu UseCase) DeletePostFiles(ctx context.Context, id int) error {
	data, err := cu.postFile.GetById(ctx, id)
	if err != nil {
		return err
	}

	for _, r := range data.Link {
		err = cu.file.Delete(ctx, r)
		if err != nil {
			return err
		}

	}

	return cu.postFile.Delete(ctx, id)
}
