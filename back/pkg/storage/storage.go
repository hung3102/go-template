// Package storage - storage関係の処理
package storage

import (
	"context"
)

// Storage - storage関係の処理
type Storage interface {
	// Upload - ファイルをアップロードする
	Upload(ctx context.Context, file []byte, path string, contentType string) error
	// DownloadURL - ダウンロードURLを取得する
	DownloadURL(path string) (string, error)
}
