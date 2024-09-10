package repositories

import "context"

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../repositoryimpl/mocks/base_repository_mock.go -package=mockrepositories

// BaseRepository - ベースリポジトリ
type BaseRepository[T any] interface {
	// Create - エンティティを作成
	Create(ctx context.Context, entity *T) error
	// GetByID - IDからエンティティを取得
	GetByID(ctx context.Context, id string) (*T, error)
}
