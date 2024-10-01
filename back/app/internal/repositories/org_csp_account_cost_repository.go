package repositories

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago/infrastructures"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../repositoryimpl/mocks/org_csp_account_cost_repository_mock.go -package=mockrepositories

// ORGCSPAccountCostRepository - org_csp_account_cost repository
type ORGCSPAccountCostRepository interface {
	// SearchByParam - Search documents by param
	SearchByParam(ctx context.Context, param *infrastructures.OrgCSPAccountCostSearchParam) ([]*entities.OrgCSPAccountCost, *infrastructures.PagingResult, error)
}
