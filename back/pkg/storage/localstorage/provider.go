package localstorage

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/topgate/gcim-temporary/back/pkg/storage"
	"golang.org/x/xerrors"
)

type impl struct {
	directoryPath string
}

// NewProvider - storage.Provider のコンストラクタ
func NewProvider(directoryPath string) storage.Provider {
	return &impl{directoryPath: directoryPath}
}

// Get - オブジェクトを取得する
func (i *impl) Get(_ context.Context, objectName string) (io.Reader, error) {
	file, err := os.Open(filepath.Join(i.directoryPath, objectName))
	if err != nil {
		return nil, xerrors.Errorf("failed to open file: %w", err)
	}
	return file, nil
}

// GetContentType - オブジェクトの Content-Type を取得する
func (i *impl) GetContentType(_ context.Context, objectName string) (string, error) {
	file, err := os.Open(filepath.Join(i.directoryPath, objectName))
	if err != nil {
		return "", xerrors.Errorf("failed to open file: %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", xerrors.Errorf("failed to read file: %w", err)
	}

	contentType := func() string {
		// javascriptはtext/javascriptとして扱う
		if filepath.Ext(objectName) == ".js" {
			return "text/javascript"
		}
		return http.DetectContentType(data)
	}()

	return contentType, nil
}

// Upload - オブジェクトをアップロードする
func (i *impl) Upload(_ context.Context, param storage.UploadParam) (string, error) {
	err := os.MkdirAll(filepath.Join(i.directoryPath, filepath.Dir(param.ObjectName)), os.ModePerm)
	if err != nil {
		return "", xerrors.Errorf("failed to create directory: %w", err)
	}

	uploadPath := filepath.Join(i.directoryPath, param.ObjectName)
	file, err := os.Create(uploadPath)
	if err != nil {
		return "", xerrors.Errorf("failed to create file: %w", err)
	}

	if _, err := io.Copy(file, param.Reader); err != nil {
		return "", xerrors.Errorf("failed to copy file: %w", err)
	}

	return uploadPath, nil
}
