package volcagoimpl

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryerrors"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago/infrastructures"
)

type gcasCSPCostImpl struct {
	infra infrastructures.GCASCSPCostRepository
}

// NewGCASCSPCost - GCASCSPCost リポジトリを生成
func NewGCASCSPCost(client *firestore.Client) repositories.GCASCSPCostRepository {
	return &gcasCSPCostImpl{
		infra: infrastructures.NewGCASCSPCostRepository(client),
	}
}

// CreateMany - 複数レコードを一括登録する
func (g *gcasCSPCostImpl) CreateMany(ctx context.Context, gcasCSPCosts []*entities.GCASCSPCost) error {
	if len(gcasCSPCosts) == 0 {
		return nil
	}

	volcagoGCASCSPCost := make([]*volcago.GCASCSPCost, 0, len(gcasCSPCosts))
	for _, gcasCSPCost := range gcasCSPCosts {
		if gcasCSPCost == nil {
			continue
		}

		volcagoGCASCSPCost = append(volcagoGCASCSPCost, &volcago.GCASCSPCost{
			ID:        gcasCSPCost.ID(),
			EventID:   gcasCSPCost.EventID(),
			CSP:       gcasCSPCost.CSP(),
			TotalCost: gcasCSPCost.TotalCost(),
			Meta: volcago.Meta{
				CreatedAt: gcasCSPCost.Meta().CreatedAt(),
				CreatedBy: gcasCSPCost.Meta().CreatedBy(),
				UpdatedAt: gcasCSPCost.Meta().UpdatedAt(),
				UpdatedBy: gcasCSPCost.Meta().UpdatedBy(),
				DeletedAt: gcasCSPCost.Meta().DeletedAt(),
				DeletedBy: gcasCSPCost.Meta().DeletedBy(),
			},
		})
	}

	if len(volcagoGCASCSPCost) == 0 {
		return nil
	}

	_, err := g.infra.InsertMulti(ctx, volcagoGCASCSPCost)
	if err != nil {
		return repositoryerrors.NewUnknownError("failed to create gcas_csp_cost", err)
	}

	return nil
}

// Exists - event_idに紐付くコレクションの存在フラグを取得する
func (g *gcasCSPCostImpl) Exists(ctx context.Context, eventID string) (bool, error) {
	qb := infrastructures.NewQueryBuilder(g.infra.GetCollection()).
		Equal("event_id", eventID).
		Limit(1)

	gcasCSPCosts, err := g.infra.Search(ctx, nil, qb.Query())
	if err != nil {
		return false, repositoryerrors.NewUnknownError("error in gcasCSPCostImpl.Exists", err)
	}

	return 0 < len(gcasCSPCosts), nil
}
