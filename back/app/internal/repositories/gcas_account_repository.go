package repositories

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../repositoryimpl/mocks/gcas_account_repository_mock.go -package=mockrepositories

// GCASAccountRepository - gcas_accountリポジトリ
type GCASAccountRepository interface {
	// CreateMany - 複数レコードを一括登録する
	CreateMany(ctx context.Context, gcasAccounts []*entities.GCASAccount) error
	// GetAccounts - 指定したevent_doc_idに紐付くコレクションを取得する。
	GetAccounts(ctx context.Context, eventDocID string) ([]*entities.GCASAccount, error)
}
