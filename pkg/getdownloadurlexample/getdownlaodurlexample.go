package getdownloadurlexample

import (
	"context"
	"fmt"
	"gcim/example/internal/api"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"

	cs "cloud.google.com/go/storage"
)

// PDFファイルダウンロード調査用のコード
// /get-download-url?path=aaaaaa/bbbbbb/PDFファイル2.pdf
type GetDownloadURLExample struct{}

func NewGetDownloadURLExample() *GetDownloadURLExample {
	return &GetDownloadURLExample{}
}

func (ge *GetDownloadURLExample) Run(c echo.Context, params api.GetDownloadUrlExampleParams) error {
	ctx := c.Request().Context()
	object := params.Path
	url, err := ge.run(ctx, object)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("get download url error. %s", err))
	}
	return c.JSON(http.StatusOK, &api.GetDownloadUrlExampleResponse{
		Url: &url,
	})
}

func (ge *GetDownloadURLExample) run(ctx context.Context, object string) (string, error) {
	bucketName := os.Getenv("PROJECT_ID") + ".appspot.com"

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "development" {
		//ローカル環境の場合は、http://${FIREBASE_STORAGE_EMULATOR_HOST}/${バケット名}/${ファイルパス}
		storageEmulatorHost := "localhost:9199"
		return fmt.Sprintf("http://%s/%s/%s", storageEmulatorHost, bucketName, object), nil
	}

	config := &firebase.Config{StorageBucket: bucketName}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		return "", fmt.Errorf("downloadExampleMain: firebase.NewApp: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return "", fmt.Errorf("downloadExampleMain: app.Storage: %v", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return "", fmt.Errorf("downloadExampleMain: client.DefaultBucket: %v", err)
	}

	opts := &cs.SignedURLOptions{
		Scheme:  cs.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}

	u, err := bucket.SignedURL(object, opts)
	if err != nil {
		return "", fmt.Errorf("downloadExampleMain: bucket.SignedURL: %v", err)
	}

	return u, nil
}
