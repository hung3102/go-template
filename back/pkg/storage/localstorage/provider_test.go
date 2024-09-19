package localstorage_test

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"io"
// 	"testing"

// 	"github.com/topgate/gcim-temporary/back/pkg/storage"
// 	"github.com/topgate/gcim-temporary/back/pkg/storage/localstorage"
// )

// func Test_Provider(t *testing.T) {
// 	sut := localstorage.NewProvider("/tmp")

// 	ctx := context.Background()
// 	objectName := "hoge.txt"

// 	// Upload確認
// 	uploadPath, err := sut.Upload(ctx, storage.UploadParam{
// 		ObjectName:  objectName,
// 		Reader:      bytes.NewReader([]byte("text")),
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
