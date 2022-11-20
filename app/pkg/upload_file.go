package pkg

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"tpk-backend/app/config"
	"tpk-backend/app/models/model"

	"github.com/gofrs/uuid"
)

func UploadReportFile(file *multipart.FileHeader) (*model.ReportMedia, error) {
	now := GetDatetime()
	path := config.LoadPathMedia()
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileName := CutnameFile(file.Filename)

	// Destination
	dst, err := os.Create(path.Path + "report/" + fileName)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

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

func UploadProfileFile(file *multipart.FileHeader, email string) (*model.ProfileMedia, error) {
	now := GetDatetime()
	path := config.LoadPathMedia()
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileName := CutnameFile(file.Filename)
	// Destination
	dst, err := os.Create(path.Path + "profile/" + fileName)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	uid, _ := uuid.NewV4()
	image := model.ProfileMedia{
		Id:          uid.String(),
		FileName:    fileName,
		Size:        file.Size,
		ContentType: file.Header.Get("Content-Type"),
		CreateAt:    now,
		Email:       email,
	}
	return &image, nil
}

func CutnameFile(filename string) string {
	splitName := strings.Split(filename, ".")
	id, _ := uuid.NewV4()
	name := fmt.Sprintf(`%v.%v`, id.String(), splitName[1])
	return name
}
