package invoiceunitlist

import (
	"context"
	"fmt"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
	"golang.org/x/xerrors"
)

// List - 団体ごとのCSPごと、と事業者ごとのCSPごとのリストを取得する
func (u *Usecase) List(ctx context.Context, input *Input) ([]*Output, error) {
	queryLimit := 1000
	uniqueNamesMap := make(map[string]*Output)
	var pagingResult *repositories.OrgCSPAccountCostPagingResult
	var accountCosts []*entities.OrgCSPAccountCost
	var err error

	eventID, err := valueobjects.NewEventIDFromString(input.EventID)
	if err != nil {
		return nil, xerrors.Errorf("error in converting eventID: %w", err)
	}

	// Loop until pagingResult is nil
	for {
		searchParam := &repositories.OrgCSPAccountCostSearchParam{
			EventID: eventID,
			Limit:   queryLimit,
		}

		if pagingResult != nil {
			searchParam.StartAtID = &pagingResult.NextID
		}

		accountCosts, pagingResult, err = u.deps.ORGCSPAccountRepository.SearchByParam(ctx, searchParam)
		if err != nil {
			return nil, xerrors.Errorf("error in List: %w", err)
		}

		for _, v := range accountCosts {
			uniqKey := fmt.Sprintf("1_%s_%s", v.OrganizationCode(), v.CSP())

			paymentAgency := v.PaymentAgency()

			if paymentAgency != nil {
				uniqKey = fmt.Sprintf("2_%s_%s", paymentAgency.CorporateNumber(), v.CSP())
			}

			if _, exists := uniqueNamesMap[uniqKey]; !exists {
				output := &Output{
					OrganizationCode: v.OrganizationCode(),
					OrganizationName: v.OrganizationName(),
					CSP:              v.CSP(),
				}

				if paymentAgency != nil {
					agencyName := paymentAgency.AgencyName()
					corporateNumber := paymentAgency.CorporateNumber()
					output.AgencyName = &agencyName
					output.CorporateNumber = &corporateNumber
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
