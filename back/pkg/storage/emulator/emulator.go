package emulator

import (
	"context"
	"fmt"
	"io"

	cloudstorage "cloud.google.com/go/storage"
	"golang.org/x/xerrors"

	"github.com/topgate/gcim-temporary/back/pkg/storage"
	"github.com/topgate/gcim-temporary/back/pkg/storage/gcs"
)

var _ storage.Provider = (*impl)(nil)

// impl - CloudStorageエミュレーターアクセス用
type impl struct {
	client      *cloudstorage.Client
	bucketName  string
	gcsProvider storage.Provider
}

// NewProvider - CloudStorageエミュレーター用のProviderを作成する
func NewProvider(client *cloudstorage.Client, bucketName string) storage.Provider {
	return &impl{
		client:      client,
		bucketName:  bucketName,
		gcsProvider: gcs.NewProvider(client, bucketName),
	}
}

// Get - ファイルを取得する
func (i *impl) Get(ctx context.Context, objectName string) (io.Reader, error) {
	result, err := i.gcsProvider.Get(ctx, objectName)
	if err != nil {
		return nil, xerrors.Errorf("error in impl.Get: %w", err)
	}
	return result, nil
}

// GetContentType - 指定したobjectのcontent-typeを取得する
func (i *impl) GetContentType(ctx context.Context, objectName string) (string, error) {
	result, err := i.gcsProvider.GetContentType(ctx, objectName)
	if err != nil {
		return "", xerrors.Errorf("error in impl.GetContentType: %w", err)
	}
	return result, nil
}

// Upload - ファイルをアップロードする
func (i *impl) Upload(ctx context.Context, param storage.UploadParam) (string, error) {
	writer := i.newWriter(ctx, param.ObjectName, param.ContentType)
	if _, err := io.Copy(writer, param.Reader); err != nil {
		return "", xerrors.Errorf("error in CloudStorage.Upload: %w", err)
	}
	if err := writer.Close(); err != nil {
		return "", xerrors.Errorf("error in CloudStorage.Upload: %w", err)
	}
	return i.DownloadURL(param.ObjectName)
}

// newWriter - バケット書き込み用のWriterを取得する
func (this *impl) newWriter(ctx context.Context, objectName string, contentType string) *cloudstorage.Writer {
	writer := this.client.Bucket(this.bucketName).Object(objectName).NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []cloudstorage.ACLRule{
		{
			Entity: cloudstorage.AllUsers,
			Role:   cloudstorage.RoleReader,
		},
	}
	return writer
}

// DownloadURL - ダウンロードURLを取得する
func (i *impl) DownloadURL(objectName string) (string, error) {
	return fmt.Sprintf("http://localhost:9199/%s/%s", i.bucketName, objectName), nil
}
