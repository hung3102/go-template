package gcs

import (
	"context"
	"errors"
	"io"
	"path/filepath"

	gcs "cloud.google.com/go/storage"
	"github.com/topgate/gcim-temporary/back/pkg/storage"
	"golang.org/x/xerrors"
)

const storageBaseURL = "https://storage.googleapis.com/"

type impl struct {
	client     *gcs.Client
	bucketName string
}

// NewProvider - constructor
func NewProvider(client *gcs.Client, bucketName string) storage.Provider {
	return &impl{
		client:     client,
		bucketName: bucketName,
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
