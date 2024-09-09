package uploadexample

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
)

type UploadExample struct{}

// PDFファイルアップロード調査用のコード
// /upload-sample/{eventId}/{orgCspDocId}
func NewUploadExample() *UploadExample {
	return &UploadExample{}
}

func (ue *UploadExample) Run(c echo.Context, eventId string, orgCspDocId string) error {
	ctx := c.Request().Context()
	if err := ue.run(ctx, eventId, orgCspDocId); err != nil {
		return c.String(http.StatusInternalServerError, "upload error")
	}
	return c.String(http.StatusOK, "upload success")
}

func (ue *UploadExample) run(ctx context.Context, eventId string, orgCspDecId string) error {
	config := &firebase.Config{
		StorageBucket: os.Getenv("PROJECT_ID") + ".appspot.com",
	}

	// Cloud Runに割り当てたサービスアカウントの権限で動作するため、コードでcredentialの指定は不要
	// 権限の設定については以下を参照
	// https://www.notion.so/Go-PDF-Cloud-Storage-3afaf9f4cecb4fce8f286cecc7f1a243
	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		return fmt.Errorf("uploadExampleMain: firebase.NewApp: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return fmt.Errorf("uploadExampleMain: app.Storage: %v", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return fmt.Errorf("uploadExampleMain: client.DefaultBucket: %v", err)
	}

	localFilename := "pkg/uploadExample/PDFファイル.pdf"                          // ローカルのファイル名
	remoteFilename := fmt.Sprintf("%s/%s/PDFファイル2.pdf", eventId, orgCspDecId) // Bucketに保存されるファイル名
	contentType := "application/pdf"

	writer := bucket.Object(remoteFilename).NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}

	f, err := os.Open(localFilename)
	if err != nil {
		return fmt.Errorf("uploadExampleMain: os.Open: %v", err)
	}
	defer f.Close()

	if _, err = io.Copy(writer, f); err != nil {
		return fmt.Errorf("uploadExampleMain: io.Copy: %v", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("uploadExampleMain: writer.Close: %v", err)
	}

	return nil
}
