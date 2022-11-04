package fileApp

import (
	"context"
	"io/ioutil"
	"time"

	cloud "cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func DownloadFile(floder string, name string) ([]byte, error) {
	ctx := context.Background()
	client, err := cloud.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := client.Bucket(bucket).Object(floder + "/" + name).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
