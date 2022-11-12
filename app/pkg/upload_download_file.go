package pkg

import (
	"io"
	"mime/multipart"
	"os"
	"tpk-backend/app/config"
	"tpk-backend/app/models/model"

	"github.com/gofrs/uuid"
)

func UploadFile(file *multipart.FileHeader, dest string) (*model.ReportMedia, error) {

	path := config.LoadPathMedia()

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileName := file.Filename

	// Destination
	dst, err := os.Create(path.Path + dest + "/" + fileName)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}
	now := GetDatetime()
	uid, _ := uuid.NewV4()
	image := model.ReportMedia{
		Id:          uid.String(),
		FileName:    fileName,
		Size:        file.Size,
		ContentType: file.Header.Get("Content-Type"),
		CreateAt:    now,
	}
	return &image, nil
}
