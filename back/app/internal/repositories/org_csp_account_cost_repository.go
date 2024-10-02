package repositories

import (
	"context"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
)

//go:generate ../../../../bin/mockgen -source=$GOFILE -destination=../repositoryimpl/mocks/org_csp_account_cost_repository_mock.go -package=mockrepositories

type OrgCSPAccountCostSearchParam struct {
	EventID   valueobjects.EventID // eventID
	Limit     int                  // limit
	StartAtID *string              // start at id
}

// OrgCSPAccountCostPagingResult - paging result
type OrgCSPAccountCostPagingResult struct {
	NextID string
	Length int
}

// ORGCSPAccountCostRepository - org_csp_account_cost repository
type ORGCSPAccountCostRepository interface {
	// SearchByParam - Search documents by param
	SearchByParam(ctx context.Context, param *OrgCSPAccountCostSearchParam) ([]*entities.OrgCSPAccountCost, *OrgCSPAccountCostPagingResult, error)
}
