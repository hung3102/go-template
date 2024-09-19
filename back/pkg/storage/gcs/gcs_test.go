package gcs_test

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"io"
// 	"testing"

// 	cloudstorage "cloud.google.com/go/storage"
// 	"github.com/topgate/gcim-temporary/back/pkg/storage"
// 	"github.com/topgate/gcim-temporary/back/pkg/storage/gcs"
// 	"google.golang.org/api/option"
// )

// func Test_Provider(t *testing.T) {
// 	ctx := context.Background()
// 	client, err := newClient(ctx)
// 	if err != nil {
// 		t.Fatalf("error sut.Get: %+v", err)
// 	}
// 	sut := gcs.NewProvider(&gcs.NewProviderParams{
// 		Client:     client,
// 		BucketName: "test-project.appspot.com",
// 		IsLocal:    true,
// 	})

// 	objectName := "hoge.txt"

// 	// Upload確認
// 	uploadPath, err := sut.Upload(ctx, storage.UploadParam{
// 		ObjectName:  objectName,
// 		Reader:      bytes.NewReader([]byte("本文")),
// 		ContentType: "plain/text",
// 	})
// 	if err != nil {
// 		t.Fatalf("error sut.Upload: %+v", err)
// 	}
// 	fmt.Println(uploadPath)

// 	// Get確認
// 	reader, err := sut.Get(ctx, objectName)
// 	if err != nil {
// 		t.Fatalf("error sut.Get: %+v", err)
// 	}
// 	data, err := io.ReadAll(reader)
// 	if err != nil {
// 		t.Fatalf("error io.ReadAll: %+v", err)
// 	}
// 	str := string(data)
// 	fmt.Println(str)

// 	// GetContentType確認
// 	contentType, err := sut.GetContentType(ctx, objectName)
// 	if err != nil {
// 		t.Fatalf("error sut.GetContentType: %+v", err)
// 	}
// 	fmt.Println(contentType)

// 	// DownloadURL確認
// 	downlaodURL, err := sut.DownloadURL(objectName)
// 	if err != nil {
// 		t.Fatalf("error sut.DownloadURL: %+v", err)
// 	}
// 	fmt.Println(downlaodURL)
// }

// func newClient(ctx context.Context) (*cloudstorage.Client, error) {
// 	options := make([]option.ClientOption, 0)
// 	options = append(options, option.WithoutAuthentication())
// 	options = append(options, option.WithEndpoint("http://localhost:9199"))
// 	client, err := cloudstorage.NewClient(ctx, options...)
// 	return client, err
// }
