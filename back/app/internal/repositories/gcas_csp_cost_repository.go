package repositories

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../repositoryimpl/mocks/gcas_csp_cost_repository_mock.go -package=mockrepositories

// GCASCSPCostRepository - gcas_csp_costリポジトリ
type GCASCSPCostRepository interface {
	// CreateMany - 複数レコードを一括登録する
	CreateMany(ctx context.Context, gcasCSPCosts []*entities.GCASCSPCost) error
	// Exists - event_doc_idに紐付くコレクションの存在フラグを取得する
	Exists(ctx context.Context, eventDocID string) (bool, error)
}
