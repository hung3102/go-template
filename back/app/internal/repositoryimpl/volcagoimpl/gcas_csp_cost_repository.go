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
	volcagoGCASCSPCost := make([]*volcago.GCASCSPCost, len(gcasCSPCosts))
	for i, gcasCSPCost := range gcasCSPCosts {
		volcagoGCASCSPCost[i] = &volcago.GCASCSPCost{
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
		}
	}
	_, err := g.infra.InsertMulti(ctx, volcagoGCASCSPCost)
	if err != nil {
		return repositoryerrors.NewUnknownError("failed to create gcas_csp_cost", err)
	}
	return nil
}

// Exists - event_idに紐付くコレクションの存在フラグを取得する
func (g *gcasCSPCostImpl) Exists(ctx context.Context, eventID string) (bool, error) {
	chainer := infrastructures.NewQueryChainer
	param := &infrastructures.GCASCSPCostSearchParam{
		EventID:     chainer().Filters(eventID, infrastructures.FilterTypeAdd),
		CursorLimit: 0,
	}

	gcasCSPCosts, err := g.infra.Search(ctx, param, nil)
	if err != nil {
		return false, repositoryerrors.NewUnknownError("error in gcasCSPCostImpl.Exists", err)
	}

	return 0 < len(gcasCSPCosts), nil
}
