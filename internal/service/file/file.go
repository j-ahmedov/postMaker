package file

import (
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) Upload(ctx context.Context, file *multipart.FileHeader, folder string) (string, error) {
	if file == nil {
		return "", nil
	}

	filename := filepath.Base(file.Filename)

	if _, err := os.Stat("./media/" + folder); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll("./media/"+folder, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	files, err := os.ReadDir("./media/" + folder)

	if err != nil {
		return "", err
	}

	for _, f := range files {
		if !f.IsDir() && (f.Name() == filename) {
			splitString := strings.Split(filename, ".")
			extra := strconv.Itoa(int(time.Now().Unix()))
			splitString[len(splitString)-2] = splitString[len(splitString)-2] + "-" + extra
			filename = strings.Join(splitString, ".")
			break
		}
	}

	dst := "./media/" + folder + "/" + filename

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()
	defer log.Println("file upload src.Close() error: ", src.Close())

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, src)

	if err != nil {
		return "err", err
	}

	return "/media/" + folder + "/" + filename, nil
}

func (s Service) Delete(ctx context.Context, dst string) error {

	return os.Remove("." + dst)
}

func (s Service) MultipleUpload(ctx context.Context, files []*multipart.FileHeader, folder string) ([]string, error) {

	var links []string

	for _, f := range files {
		link, err := s.Upload(ctx, f, folder)

		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}
