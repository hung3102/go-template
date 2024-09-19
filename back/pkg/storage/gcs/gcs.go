package gcs

import (
	"context"
	"errors"
	"io"
	"path/filepath"
	"time"

	gcs "cloud.google.com/go/storage"
	"github.com/topgate/gcim-temporary/back/pkg/storage"
	"golang.org/x/xerrors"
)

const storageBaseURL = "https://storage.googleapis.com/"

// 署名付きURLの有効期間(分)
const expiresMin = 15

type impl struct {
	client     *gcs.Client
	bucketName string
	expiresMin int
}

// NewProvider - constructor
func NewProvider(client *gcs.Client, bucketName string) storage.Provider {
	return &impl{
		client:     client,
		bucketName: bucketName,
		expiresMin: expiresMin,
	}
}

// Get - get content
func (i *impl) Get(ctx context.Context, objectName string) (io.Reader, error) {
	reader, err := i.client.Bucket(i.bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		if errors.Is(err, gcs.ErrObjectNotExist) {
			return nil, xerrors.Errorf("failed to get bucket(object=%s, bucket=%s): %w", objectName, i.bucketName, storage.ErrBucketDoesNotExist)
		} else if errors.Is(err, gcs.ErrBucketNotExist) {
			return nil, xerrors.Errorf("failed to get object(object=%s, bucket=%s): %w", objectName, i.bucketName, storage.ErrObjectDoesNotExist)
		}
		return nil, xerrors.Errorf("failed to get reader(object=%s, bucket=%s): %w", objectName, i.bucketName, err)
	}

	return reader, nil
}

// GetContentType - get Content-Type
func (i *impl) GetContentType(ctx context.Context, objectName string) (string, error) {
	attrs, err := i.client.Bucket(i.bucketName).Object(objectName).Attrs(ctx)
	if err != nil {
		if errors.Is(err, gcs.ErrObjectNotExist) {
			return "", xerrors.Errorf("failed to get bucket(object=%s, bucket=%s): %w", objectName, i.bucketName, storage.ErrBucketDoesNotExist)
		} else if errors.Is(err, gcs.ErrBucketNotExist) {
			return "", xerrors.Errorf("failed to get object(object=%s, bucket=%s): %w", objectName, i.bucketName, storage.ErrObjectDoesNotExist)
		}
		return "", xerrors.Errorf("failed to get attrs(object=%s, bucket=%s): %w", objectName, i.bucketName, err)
	}

	return attrs.ContentType, nil
}

// Upload - upload content
func (i *impl) Upload(ctx context.Context, param storage.UploadParam) (string, error) {
	bucket := i.client.Bucket(i.bucketName)
	if _, err := bucket.Attrs(ctx); err != nil {
		return "", xerrors.Errorf("error in bucket.Attrs method: %w", err)
	}
	object := bucket.Object(param.ObjectName)
	writer := object.NewWriter(ctx)

	writer.ObjectAttrs.ContentType = param.ContentType
	writer.ObjectAttrs.CacheControl = "no-cache"

	if _, err := io.Copy(writer, param.Reader); err != nil {
		return "", xerrors.Errorf("error in io.Copy method: %w", err)
	}

	if err := writer.Close(); err != nil {
		return "", xerrors.Errorf("error in writer.Close method: %w", err)
	}

	return storageBaseURL + filepath.Join(i.bucketName, param.ObjectName), nil
}

// DownloadURL - 署名付きURLを取得する
func (i *impl) DownloadURL(path string) (string, error) {
	url, err := i.client.Bucket(i.bucketName).SignedURL(path, i.signedURLOptions())
	if err != nil {
		return "", xerrors.Errorf("error in CloudStorage.DownloadURL: %w", err)
	}

	return url, nil
}

// signedURLOptions - 署名付きURL取得設定を取得する
func (i *impl) signedURLOptions() *gcs.SignedURLOptions {
	return &gcs.SignedURLOptions{
		Scheme:  gcs.SigningSchemeV4,
		Method:  "GET",
		Expires: i.expires(),
	}
}

// expires - 有効期限を取得する
func (i *impl) expires() time.Time {
	return time.Now().Add(time.Duration(i.expiresMin) * time.Minute)
}
