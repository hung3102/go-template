package storage

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func Test_CloudStorage_Upload_DownloadURL(t *testing.T) {
	ctx := context.Background()

	client, err := newClient(ctx)
	if err != nil {
		t.Errorf("Test_CloudStorage_Upload_DownloadURL: newClient: %v", err)
	}

	sut := NewCloudStorage(&CloudStorageParam{
		Client:     client,
		BucketName: "test-project.appspot.com",
		IsLocal:    true,
	})
	fmt.Println(sut)

	// file := []byte("aaaaaaaa")

	// err = sut.Upload(ctx, file, "file/path/filename.txt", "text/plain")
	// if err != nil {
	// 	t.Errorf("Test_CloudStorage_Upload_DownloadURL: sut.Upload: %v", err)
	// }

	// url, err := sut.DownloadURL("file/path/filename.txt")
	// if err != nil {
	// 	t.Errorf("Test_CloudStorage_Upload_DownloadURL: sut.DownloadURL: %v", err)
	// }
	// fmt.Println(url)
}

func newClient(ctx context.Context) (*storage.Client, error) {
	options := make([]option.ClientOption, 0)
	options = append(options, option.WithoutAuthentication())
	options = append(options, option.WithEndpoint("http://localhost:9199"))
	client, err := storage.NewClient(ctx, options...)
	return client, err
}
