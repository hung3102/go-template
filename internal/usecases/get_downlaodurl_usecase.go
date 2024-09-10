package usecases

import (
	"context"
	"fmt"
	"os"

	"gcim/example/internal/domain/storage"
	"gcim/example/internal/usecases/dto/input"
	"gcim/example/internal/usecases/dto/output"
)

// PDFファイルダウンロード調査用のコード
// /get-download-url?path=aaaaaa/bbbbbb/PDFファイル2.pdf
type IGetDownloadUrlUsecase interface {
	Execute(ctx context.Context, params *input.GetDownloadUrlInput) (*output.GetDownloadUrlOutput, error)
}

type GetDownloadUrlUsecase struct {
	storage storage.Storage
}

func NewGetDownloadUrlUsecase(storage storage.Storage) IGetDownloadUrlUsecase {
	return &GetDownloadUrlUsecase{
		storage: storage,
	}
}

func (u *GetDownloadUrlUsecase) Execute(ctx context.Context, params *input.GetDownloadUrlInput) (*output.GetDownloadUrlOutput, error) {
	object := params.Path

	var url string = ""
	var err error = nil
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "development" {
		//ローカル環境の場合は、http://${FIREBASE_STORAGE_EMULATOR_HOST}/${バケット名}/${ファイルパス}
		storageEmulatorHost := "localhost:9199"
		bucketName := u.storage.BucketName()
		url = fmt.Sprintf("http://%s/%s/%s", storageEmulatorHost, bucketName, object)
	} else {
		url, err = u.storage.SignedUrl(object)
		if err != nil {
			return nil, fmt.Errorf("GetDownlaodUrlUsecase.Execute u.storage.SignedUrl: %v", err)
		}
	}

	output := &output.GetDownloadUrlOutput{
		Url: url,
	}
	return output, nil
}
