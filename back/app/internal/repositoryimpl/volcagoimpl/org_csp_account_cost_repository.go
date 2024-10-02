package volcagoimpl

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago/infrastructures"
	"golang.org/x/xerrors"
)

type impl struct {
	infra infrastructures.OrgCSPAccountCostRepository
}

// NewOrgCSPAccountCost - Constructor
func NewOrgCSPAccountCost(client *firestore.Client) repositories.ORGCSPAccountCostRepository {
	return &impl{
		infra: infrastructures.NewOrgCSPAccountCostRepository(client),
	}
}

func (i *impl) SearchByParam(ctx context.Context, param *repositories.OrgCSPAccountCostSearchParam) ([]*entities.OrgCSPAccountCost, *repositories.OrgCSPAccountCostPagingResult, error) {
	searchParam := &infrastructures.OrgCSPAccountCostSearchParam{
		EventID: &infrastructures.QueryChainer{
			QueryGroup: []*infrastructures.Query{{
				Operator: infrastructures.OpTypeEqual,
				Value:    param.EventID.String(),
			},
			},
		},
		CursorLimit: param.Limit,
	}

	if param.StartAtID != nil {
		searchParam.CursorKey = *param.StartAtID
	}

	accountCosts, pagingResult, err := i.infra.SearchByParam(ctx, searchParam)

	if err != nil {
		return nil, nil, xerrors.Errorf("error in SearchByParam: %w", err)
	}

	responseCosts := make([]*entities.OrgCSPAccountCost, 0, len(accountCosts))

	for _, v := range accountCosts {
		orgCspAccountCostID, err := valueobjects.NewOrgCSPAccountCostIDFromString(v.ID)
		if err != nil {
			return nil, nil, xerrors.Errorf("error in converting orgCspAccountCostId: %w", err)
		}

		eventID, err := valueobjects.NewEventIDFromString(v.EventID)
		if err != nil {
			return nil, nil, xerrors.Errorf("error in converting eventId: %w", err)
		}

		gcasProportionCostID, err := valueobjects.NewGCASProportionCostIDFromString(v.GCASProportionCostID)
		if err != nil {
			return nil, nil, xerrors.Errorf("error in converting gcasProportionCostID: %w", err)
		}

		gcasAccountCostID, err := valueobjects.NewGCASAccountCostIDFromString(v.GCASAccountCostID)
		if err != nil {
			return nil, nil, xerrors.Errorf("error in converting gcasAccountCostID: %w", err)
		}

		var paymentAgency *entities.PaymentAgency
		if v.PaymentAgency != nil {

			paymentAgency = entities.NewPaymentAgency(&entities.NewPaymentAgencyParam{
				AgencyName:      v.PaymentAgency.AgencyName,
				CorporateNumber: v.PaymentAgency.CorporateNumber,
			})
		}

		responseCosts = append(responseCosts, entities.NewOrgCSPAccountCost(&entities.NewOrgCSPAccountCostParam{
			ID:                   orgCspAccountCostID,
			EventID:              eventID,
			GCASProportionCostID: gcasProportionCostID,
			GCASAccountCostID:    gcasAccountCostID,
			OrganizationCode:     v.OrganizationCode,
			OrganizationName:     v.OrganizationName,
			CSP:                  v.CSP,
			AccountID:            v.AccountID,
			Cost:                 v.Cost,
			BillingUnitID:        v.BillingUnitID,
			PaymentAgency:        paymentAgency,
			Meta: entities.NewMeta(&entities.NewMetaParam{
				CreatedAt: v.CreatedAt,
				CreatedBy: v.CreatedBy,
				UpdatedAt: v.UpdatedAt,
				UpdatedBy: v.UpdatedBy,
				DeletedAt: v.DeletedAt,
				DeletedBy: v.DeletedBy,
			}),
		}))
	}

	var responsePagingResult *repositories.OrgCSPAccountCostPagingResult
	if pagingResult != nil {
		responsePagingResult = &repositories.OrgCSPAccountCostPagingResult{
			NextID: pagingResult.NextCursorKey,
			Length: pagingResult.Length,
		}
	}

	return responseCosts, responsePagingResult, nil
}
