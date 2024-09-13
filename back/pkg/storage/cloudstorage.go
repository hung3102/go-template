package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

var _ Storage = (*CloudStorage)(nil)

// CloudStorage - CloudStorageアクセス用
type CloudStorage struct {
	client     *storage.Client
	bucketName string
	isLocal    bool
}

// CloudStorageParam - CloudStorageを作成するためのパラメーター
type CloudStorageParam struct {
	Client     *storage.Client
	BucketName string
	IsLocal    bool
}

// NewStorage - Storageを作成する(Cloud Storage用)
func NewCloudStorage(params *CloudStorageParam) Storage {
	return &CloudStorage{
		client:     params.Client,
		bucketName: params.BucketName,
		isLocal:    params.IsLocal,
	}
}

// Upload - ファイルをアップロードする
func (this *CloudStorage) Upload(ctx context.Context, file []byte, path string, contentType string) error {
	writer := this.newWriter(ctx, path, contentType)
	reader := bytes.NewReader(file)
	if _, err := io.Copy(writer, reader); err != nil {
		return fmt.Errorf("CloudStorage.Upload: io.Copy: %v", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("CloudStorage.Upload: writer.Close: %v", err)
	}
	return nil
}

// newWriter - バケット書き込み用のWriterを取得する
func (this *CloudStorage) newWriter(ctx context.Context, path string, contentType string) *storage.Writer {
	writer := this.bucket().Object(path).NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}
	return writer
}

// bucket - バケットを取得する
func (this *CloudStorage) bucket() *storage.BucketHandle {
	return this.client.Bucket(this.bucketName)
}

// DownloadURL - ファイルのダウンロードURLを取得する(本番環境は署名付きURL)
func (this *CloudStorage) DownloadURL(path string) (string, error) {
	if this.isLocal {
		return this.localDownloadURL(path), nil
	}
	url, err := this.bucket().SignedURL(path, this.signedURLOptions())
	if err != nil {
		return "", fmt.Errorf("CloudStorage.DownloadURL: bucket.SignedURL: %v", err)
	}

	return url, nil
}

// localDownloadURL - ローカル環境の場合のダウンロードURLを取得する http://${FIREBASE_STORAGE_EMULATOR_HOST}/${バケット名}/${ファイルパス}
func (this *CloudStorage) localDownloadURL(path string) string {
	return fmt.Sprintf("http://localhost:9199/%s/%s", this.bucketName, path)
}

// signedURLOptions - 署名付きURL取得設定を取得する
func (this *CloudStorage) signedURLOptions() *storage.SignedURLOptions {
	return &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}
}
