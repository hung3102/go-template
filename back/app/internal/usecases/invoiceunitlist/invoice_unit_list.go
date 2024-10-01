package invoiceunitlist

import (
	"context"
	"fmt"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago/infrastructures"
	"golang.org/x/xerrors"
)

// List - 団体ごとのCSPごと、と事業者ごとのCSPごとのリストを取得する
func (u *Usecase) List(ctx context.Context, input *Input) ([]*Output, error) {
	queryLimit := 1000
	uniqueNamesMap := make(map[string]*Output)
	var pagingResult *infrastructures.PagingResult
	var accountCosts []*entities.OrgCSPAccountCost
	var err error

	// Loop until pagingResult is nil
	for {
		searchParam := &infrastructures.OrgCSPAccountCostSearchParam{
			EventID: &infrastructures.QueryChainer{
				QueryGroup: []*infrastructures.Query{{
					Operator: infrastructures.OpTypeEqual,
					Value:    input.EventID,
				},
				},
			},
			CursorLimit: queryLimit,
		}

		if pagingResult != nil {
			searchParam.CursorKey = pagingResult.NextCursorKey
		}

		accountCosts, pagingResult, err = u.deps.ORGCSPAccountRepository.SearchByParam(ctx, searchParam)
		if err != nil {
			return nil, xerrors.Errorf("error in List: %w", err)
		}

		for _, v := range accountCosts {
			uniqKey := fmt.Sprintf("%s_%s", v.Organization(), v.CSP())

			// TODOH: test PaymentAgency
			paymentAgency := v.PaymentAgency()

			if paymentAgency != nil {
				// TODOH: test AgencyName()
				uniqKey = fmt.Sprintf("%s_%s", paymentAgency.AgencyName(), v.CSP())
			}

			if _, exists := uniqueNamesMap[uniqKey]; !exists {
				output := &Output{
					IsPaymentAgent: false,            // temp
					Subject:        v.Organization(), // temp
					CSP:            v.CSP(),
				}

				if paymentAgency != nil {
					output.Subject = paymentAgency.AgencyName() // TODOH: test AgencyName()
					output.IsPaymentAgent = true
				}

				uniqueNamesMap[uniqKey] = output
			}
		}

		if pagingResult == nil {
			break
		}
	}

	output := make([]*Output, 0, queryLimit)
	for _, v := range uniqueNamesMap {
		output = append(output, v)
	}

	return output, nil
}
