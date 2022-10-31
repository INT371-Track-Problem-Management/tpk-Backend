package fileApp

import (
	"context"
	"io"
	"mime/multipart"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/pkg/config"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
)

var fileConfig = config.LoadFileConfig()

func UploadFile(context context.Context, storage *cloud.Client, file multipart.File, handler *multipart.FileHeader, client *firestore.Client) (*string, error) {
	defer file.Close()

	imagePath := handler.Filename

	wc := storage.Bucket(fileConfig.Bucket).Object(imagePath).NewWriter(context)
	_, err := io.Copy(wc, file)
	if err != nil {
		return nil, err
	}

	if err := wc.Close(); err != nil {
		return nil, err
	}

	err = CreateImageUrl(imagePath, context, client)
	if err != nil {
		return nil, err
	}

	return &imagePath, nil
}

func CreateImageUrl(imagePath string, ctx context.Context, client *firestore.Client) error {
	imageStructure := entity.Media{
		Name: imagePath,
		URL:  fileConfig.URL + "/" + fileConfig.Bucket + "/" + imagePath,
	}

	_, _, err := client.Collection("image").Add(ctx, imageStructure)
	if err != nil {
		return err
	}

	return nil
}
