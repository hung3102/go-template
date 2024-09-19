package storage

import (
	"context"
	"io"

	"golang.org/x/xerrors"
)

// UploadParam - Upload のパラメータ
type UploadParam struct {
	ObjectName  string    // アップロード先のオブジェクト名
	Reader      io.Reader // アップロードするデータ
	ContentType string    // Content-Type
}

// Provider - storage provider
type Provider interface {
	Get(ctx context.Context, objectName string) (io.Reader, error)
	GetContentType(ctx context.Context, objectName string) (string, error)
	Upload(ctx context.Context, param UploadParam) (string, error)
}

var (
	// ErrBucketDoesNotExist - bucket doesn't exist
	ErrBucketDoesNotExist = xerrors.New("bucket doesn't exist")
	// ErrObjectDoesNotExist - object doesn't exist
	ErrObjectDoesNotExist = xerrors.New("object doesn't exist")
)
