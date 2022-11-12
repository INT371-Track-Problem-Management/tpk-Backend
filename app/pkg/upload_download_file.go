package pkg

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(file *multipart.FileHeader) (*string, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	fileName := file.Filename

	// Destination
	dst, err := os.Create("../../images/" + fileName)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}
	return &fileName, nil
}
