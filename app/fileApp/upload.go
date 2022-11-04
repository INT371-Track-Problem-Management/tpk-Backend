package fileApp

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/pkg/config"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
)

var fileConfig = config.LoadFileConfig()
var bucket = fileConfig.Bucket

// var opts *cloud.SignedURLOptions

func UploadFileToStorage(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func UploadFile(context context.Context, storage *cloud.Client, file multipart.File, handler *multipart.FileHeader, client *firestore.Client, floder string) (*string, error) {
	defer file.Close()

	imagePath := handler.Filename

	obj := storage.Bucket(bucket).Object(floder + "/" + imagePath)

	wc := obj.NewWriter(context)
	_, err := io.Copy(wc, file)
	if err != nil {
		return nil, err
	}

	if err := wc.Close(); err != nil {
		return nil, err
	}

	r, err := obj.NewReader(context)
	if err != nil {
		return nil, err
	}

	defer r.Close()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return nil, err
	}

	objAttrs, err := obj.Attrs(context)
	if err != nil {
		return nil, err
	}
	url := objAttrs.MediaLink

	// writer := object.NewWriter(context)
	// if err := writer.Close(); err != nil {
	// }
	// fmt.Println("file: ", url)

	// err = CreateImageUrl(imagePath, context, client, floder)
	// if err != nil {
	// 	return nil, err
	// }

	return &url, nil
}

func CreateImageUrl(imagePath string, ctx context.Context, client *firestore.Client, floder string) error {
	imageStructure := entity.Media{
		Name: imagePath,
		URL:  fileConfig.URL + "/" + fileConfig.Bucket + "/" + floder + "/" + imagePath,
	}

	_, _, err := client.Collection("image").Add(ctx, imageStructure)
	if err != nil {
		return err
	}

	return nil
}
