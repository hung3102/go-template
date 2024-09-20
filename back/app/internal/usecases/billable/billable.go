package billable

import (
	"context"
	"slices"

	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasapi"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasdashboardapi"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"github.com/topgate/gcim-temporary/back/app/internal/usecaseerrors"
	"golang.org/x/xerrors"
)

// Billable - 請求書作成の開始判定をする
func (u *Usecase) Billable(ctx context.Context, input *Input) (*Output, error) {
	shouldcreateInvoice, err := u.deps.EventStatusService.ShouldcreateInvoice(ctx, input.EventDocID)
	if err != nil {
		return nil, xerrors.Errorf("error in billable.Billable: %w", err)
	}
	if !shouldcreateInvoice {
		return u.emptyOutput(), nil
	}

	result, err := u.billableMain(ctx, input)
	if err != nil {
		return nil, xerrors.Errorf("error in billable.Billable: %w", err)
	}

	if err := u.deps.EventStatusService.SetInvoiceCreationChecked(ctx, input.EventDocID); err != nil {
		return nil, xerrors.Errorf("error in billable.Billable: %w", err)
	}
	return result, nil
}

// emptyOutput - 空のOutputを作成する
func (u *Usecase) emptyOutput() *Output {
	return &Output{
		GCASAccounts: []*OutputAccount{},
	}
}

// billableMain - 請求書作成の開始判定のメイン処理
func (u *Usecase) billableMain(ctx context.Context, input *Input) (*Output, error) {
	gcasDashboardAPIGetAccountsResponse, err := u.fetchAccountInfo()
	if err != nil {
		return nil, xerrors.Errorf("error in billable.createAccountAndCost: %w", err)
	}

	err = u.createGCASCSPCost(ctx, input.EventDocID, gcasDashboardAPIGetAccountsResponse)
	if err != nil {
		return nil, xerrors.Errorf("error in billable.createAccountAndCost: %w", err)
	}

	return u.toOutputFromGCASAccount(gcasDashboardAPIGetAccountsResponse), nil
}

// fetchAccountInfo - APIからアカウント情報を取得する
func (u *Usecase) fetchAccountInfo() (*gcasdashboardapi.GetAccountsResponse, error) {
	gcasDashboardAPIGetAccountsResponse, err := u.deps.GCASDashboardAPI.GetAccounts()
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeGCASDashboardAPI, "error in billable.fetchAccountInfo", err)
	}

	gcasAPIGetAccountsResponse, err := u.deps.GCASAPI.GetAccounts()
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeGCASAPI, "error in billable.fetchAccountInfo", err)
	}

	err = u.compareAccountInfo(gcasDashboardAPIGetAccountsResponse, gcasAPIGetAccountsResponse)
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeGCASAPI, "error in billable.fetchAccountInfo", err)
	}

	return gcasDashboardAPIGetAccountsResponse, nil
}

// compareAccountInfo - GCASダッシュボードAPIとGCASAPIのアカウント情報が一致するか確認する
func (u *Usecase) compareAccountInfo(
	gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse,
	gcasAPIGetAccountsResponse *gcasapi.GetAccountsResponse,
) error {
	gcaDashboardCspAccountMap := make(map[string][]string)
	for csp, gcasDashboardAccounts := range *gcasDashboardAPIGetAccountsResponse {
		slices.Sort(gcasDashboardAccounts)
		gcaDashboardCspAccountMap[csp] = gcasDashboardAccounts
	}
	gcapCspAccountMap := make(map[string][]string)
	for csp, gcasDashboardAccounts := range *gcasAPIGetAccountsResponse {
		slices.Sort(gcasDashboardAccounts)
		gcapCspAccountMap[csp] = gcasDashboardAccounts
	}

	if len(gcaDashboardCspAccountMap) != len(gcapCspAccountMap) {
		return usecaseerrors.NewUnknownError(errorcode.ErrorCodeAccountInfoIsMissing, "error in billable.compareAccountInfo", nil)
	}

	for csp, gcapDashboardAccountIDs := range gcaDashboardCspAccountMap {
		gcapAccountIDs, ok := gcapCspAccountMap[csp]
		if !ok {
			return usecaseerrors.NewUnknownError(errorcode.ErrorCodeAccountInfoIsMissing, "error in billable.compareAccountInfo", nil)
		}
		if len(gcapDashboardAccountIDs) != len(gcapAccountIDs) {
			return usecaseerrors.NewUnknownError(errorcode.ErrorCodeAccountInfoIsMissing, "error in billable.compareAccountInfo", nil)
		}
		for i, gcasDashboardAccountID := range gcapDashboardAccountIDs {
			if gcapAccountIDs[i] != gcasDashboardAccountID {
				return usecaseerrors.NewUnknownError(errorcode.ErrorCodeAccountInfoIsMissing, "error in billable.compareAccountInfo", nil)
			}
		}
	}
	return nil
}

// createGCASCSPCost - GCASCSPCostを登録する
func (u *Usecase) createGCASCSPCost(ctx context.Context, eventDocID string, gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse) error {
	gcapCSPCostExists, err := u.deps.GCASCSPCostRepository.Exists(ctx, eventDocID)
	if err != nil {
		return usecaseerrors.NewUnknownError(errorcode.ErrorCodeDBAccess, "error in billable.createGCASCSPCost", err)
	}
	if gcapCSPCostExists {
		return nil
	}

	gcasCSPCosts, err := u.toGCASCSPCost(eventDocID, gcasDashboardAPIGetAccountsResponse)
	if err != nil {
		return xerrors.Errorf("error in billable.createGCASCSPCost: %w", err)
	}

	err = u.deps.GCASCSPCostRepository.CreateMany(ctx, gcasCSPCosts)
	if err != nil {
		return usecaseerrors.NewUnknownError(errorcode.ErrorCodeDBAccess, "error in billable.createGCASCSPCost", err)
	}
	return nil
}

func (u *Usecase) toGCASCSPCost(eventDocID string, gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse) ([]*entities.GCASCSPCost, error) {
	cspCostInfo, err := u.fetchCostInfo(gcasDashboardAPIGetAccountsResponse)
	if err != nil {
		return nil, xerrors.Errorf("error in billable.createGCASCSPCost: %w", err)
	}

	cspTotalCostMap := u.toCSPTotalCostMapFromCspAccountIDCostInfoMap(cspCostInfo)

	gcasCSPCosts, err := u.toGCAPCSPCostsFromCostTotalCostMap(eventDocID, cspTotalCostMap)
	if err != nil {
		return nil, xerrors.Errorf("error in billable.createGCASCSPCost: %w", err)
	}

	return gcasCSPCosts, nil
}

// fetchCostInfo - GCASダッシュボードAPIを実行しコスト情報を取得する。
func (u *Usecase) fetchCostInfo(gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse) (map[string][]*gcasdashboardapi.GetCostResponse, error) {
	result := make(map[string][]*gcasdashboardapi.GetCostResponse)
	for csp, accountIDs := range *gcasDashboardAPIGetAccountsResponse {
		costInfo := make([]*gcasdashboardapi.GetCostResponse, len(accountIDs))
		for i, accountID := range accountIDs {
			gcasDashboardAPIGetCostResponse, err := u.deps.GCASDashboardAPI.GetCost(accountID)
			if err != nil {
				return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeGCASDashboardAPI, "error in billable.fetchCostInfo", err)
			}
			costInfo[i] = gcasDashboardAPIGetCostResponse
		}
		result[csp] = costInfo
	}
	return result, nil
}

// toCSPTotalCostMapFromCspAccountIDCostInfoMap - GCASダッシュボードAPIを実行しコスト情報を取得する。
func (u *Usecase) toCSPTotalCostMapFromCspAccountIDCostInfoMap(cspCostInfo map[string][]*gcasdashboardapi.GetCostResponse) map[string]int {
	result := make(map[string]int)
	for csp, costInfo := range cspCostInfo {
		result[csp] = 0
		for _, ci := range costInfo {
			result[csp] = result[csp] + ci.TotalCost
		}
	}
	return result
}

// toGCAPCSPCostsFromCostTotalCostMap - コスト情報をGCASCSPCostに変換する
func (u *Usecase) toGCAPCSPCostsFromCostTotalCostMap(eventDocID string, cspTotalCostMap map[string]int) ([]*entities.GCASCSPCost, error) {
	result := []*entities.GCASCSPCost{}
	for csp, totalCost := range cspTotalCostMap {
		uuid, err := u.deps.UUID.GetUUID()
		if err != nil {
			return nil, xerrors.Errorf("error in billable.toGCAPCSPCostsFromCostTotalCostMap: %w", err)
		}
		result = append(result, entities.NewGCASCSPCost(
			&entities.NewGCASCSPCostParam{
				ID:         uuid,
				EventDocID: eventDocID,
				CSP:        csp,
				TotalCost:  totalCost,
			},
		))
	}
	return result, nil
}

// toOutputFromGCASAccount - gcasDashboardAPIGetAccountsResponseをOutputに変換する
func (u *Usecase) toOutputFromGCASAccount(gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse) *Output {
	outputAccounts := []*OutputAccount{}
	for csp, accontIDs := range *gcasDashboardAPIGetAccountsResponse {
		for _, accountID := range accontIDs {
			outputAccounts = append(outputAccounts, &OutputAccount{
				CSP:       csp,
				AccountID: accountID,
			})
		}
	}
	return &Output{
		GCASAccounts: outputAccounts,
	}
}
