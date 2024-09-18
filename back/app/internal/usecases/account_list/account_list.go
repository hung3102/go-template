package accountlist

import (
	"context"
	"slices"

	gcasapi "github.com/topgate/gcim-temporary/back/app/internal/api/gcas_api"
	gcasdashboardapi "github.com/topgate/gcim-temporary/back/app/internal/api/gcas_dashboard_api"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/errorcode"
	"github.com/topgate/gcim-temporary/back/app/internal/usecaseerrors"
	"golang.org/x/xerrors"
)

// AccountList - アカウントリストを登録する
func (u *Usecase) AccountList(ctx context.Context, input *AccoutnListInput) (*AccountListOutput, error) {
	event, err := u.deps.EventsRepository.GetByID(ctx, input.EventDocID)
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeDBAccess, "error in AccountList", err)
	}
	if !event.IsInvoiceCreateionPossible() {
		return u.emptyAccountListOutput(), nil
	}
	return u.accountListMain(ctx, input)
}

// emptyAccountListOutput - 空のAccountListOutputを作成する
func (u *Usecase) emptyAccountListOutput() *AccountListOutput {
	return &AccountListOutput{
		GCASAccountDocIDs: []string{},
	}
}

// accountListMain - アカウントリスト作成のメイン処理
func (u *Usecase) accountListMain(ctx context.Context, input *AccoutnListInput) (*AccountListOutput, error) {
	gcapCSPCostExists, err := u.deps.GCASCSPCostRepository.Exists(ctx, input.EventDocID)
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeDBAccess, "error in accountListMain", err)
	}
	if gcapCSPCostExists {
		return u.getAccountListOutputFromGCASAccountRepository(ctx, input)
	}
	return u.createAccountAndCost(ctx, input)
}

// getAccountListOutputFromGCASAccountRepository - GCASAccountRepositoryからアカウント情報を取得しAccountListOutputを作成する
func (u *Usecase) getAccountListOutputFromGCASAccountRepository(ctx context.Context, input *AccoutnListInput) (*AccountListOutput, error) {
	gcasAccounts, err := u.deps.GCASAccountRepository.GetAccounts(ctx, input.EventDocID)
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeDBAccess, "error in getAccountListOutputFromGCASAccountRepository", err)
	}
	return u.toAccountListOutputFromGCASAccount(gcasAccounts), nil
}

// toAccountListOutputFromGCASAccount - []*entities.GCASAccountをAccountListOutputに変換する
func (u *Usecase) toAccountListOutputFromGCASAccount(gcasAccounts []*entities.GCASAccount) *AccountListOutput {
	return &AccountListOutput{
		GCASAccountDocIDs: u.toAccountListFromGCASAccount(gcasAccounts),
	}
}

// toAccountListFromGCASAccount - []*entities.GCASAccountをアカウントIDの配列に変換する
func (u *Usecase) toAccountListFromGCASAccount(gcasAccounts []*entities.GCASAccount) []string {
	result := make([]string, len(gcasAccounts))
	for i, gcasAccount := range gcasAccounts {
		result[i] = gcasAccount.ID()
	}
	return result
}

// createAccountAndCost - アカウント情報を取得し、アカウント情報とコスト情報をDBに登録する
func (u *Usecase) createAccountAndCost(ctx context.Context, input *AccoutnListInput) (*AccountListOutput, error) {
	gcasDashboardAPIGetAccountsResponse, err := u.fetchAccountInfo()
	if err != nil {
		return nil, xerrors.Errorf("error in createAccountAndCost: %w", err)
	}

	gcasAccounts, err := u.deps.GCASAccountRepository.GetAccounts(ctx, input.EventDocID)
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeDBAccess, "error in createAccountAndCost", err)
	}

	if len(gcasAccounts) == 0 {
		gcasAccounts, err = u.createGCASAccount(ctx, input.EventDocID, gcasDashboardAPIGetAccountsResponse)
		if err != nil {
			return nil, xerrors.Errorf("error in createAccountAndCost: %w", err)
		}
	}

	err = u.createGCASCSPCost(ctx, input.EventDocID, gcasDashboardAPIGetAccountsResponse)
	if err != nil {
		return nil, xerrors.Errorf("error in createAccountAndCost: %w", err)
	}

	return u.toAccountListOutputFromGCASAccount(gcasAccounts), nil
}

// fetchAccountInfo - APIからアカウント情報を取得する
func (u *Usecase) fetchAccountInfo() (*gcasdashboardapi.GetAccountsResponse, error) {
	gcasDashboardAPIGetAccountsResponse, err := u.deps.GCASDashboardAPI.GetAccounts()
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeGCASDashboardAPI, "error in fetchAccountInfo", err)
	}

	gcasAPIGetAccountsResponse, err := u.deps.GCASAPI.GetAccounts()
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeGCASAPI, "error in fetchAccountInfo", err)
	}

	err = u.compareAccountInfo(gcasDashboardAPIGetAccountsResponse, gcasAPIGetAccountsResponse)
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeGCASAPI, "error in fetchAccountInfo", err)
	}

	return gcasDashboardAPIGetAccountsResponse, nil
}

// compareAccountInfo - GCASダッシュボードAPIとGCASAPIのアカウント情報が一致するか確認する
func (u *Usecase) compareAccountInfo(
	gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse,
	gcasAPIGetAccountsResponse *gcasapi.GetAccountsResponse,
) error {
	gcaDashboardCspAccountMap := map[string][]string{}
	for csp, gcasDashboardAccounts := range *gcasDashboardAPIGetAccountsResponse {
		slices.Sort(gcasDashboardAccounts)
		gcaDashboardCspAccountMap[csp] = gcasDashboardAccounts
	}
	gcapCspAccountMap := map[string][]string{}
	for csp, gcasDashboardAccounts := range *gcasAPIGetAccountsResponse {
		slices.Sort(gcasDashboardAccounts)
		gcapCspAccountMap[csp] = gcasDashboardAccounts
	}

	if len(gcaDashboardCspAccountMap) != len(gcapCspAccountMap) {
		return usecaseerrors.NewUnknownError(errorcode.ErrorCodeAccountInfoIsMissing, "error in compareAccountInfo: accountID does not match", nil)
	}

	for csp, gcapDashboardAccountIDs := range gcaDashboardCspAccountMap {
		gcapAccountIDs, ok := gcapCspAccountMap[csp]
		if !ok {
			return usecaseerrors.NewUnknownError(errorcode.ErrorCodeAccountInfoIsMissing, "error in compareAccountInfo: accountID does not match", nil)
		}
		for len(gcapDashboardAccountIDs) != len(gcapAccountIDs) {
			return usecaseerrors.NewUnknownError(errorcode.ErrorCodeAccountInfoIsMissing, "error in compareAccountInfo: accountID does not match", nil)
		}
		for i, gcasDashboardAccountID := range gcapDashboardAccountIDs {
			if gcapAccountIDs[i] != gcasDashboardAccountID {
				return usecaseerrors.NewUnknownError(errorcode.ErrorCodeAccountInfoIsMissing, "error in compareAccountInfo: accountID does not match", nil)
			}
		}
	}
	return nil
}

// createGCASAccount - アカウント情報をDBに登録する
func (u *Usecase) createGCASAccount(ctx context.Context, eventDocID string, gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse) ([]*entities.GCASAccount, error) {
	gcasAccounts, err := u.toGCASAccountsFromGetAccountsResponse(eventDocID, gcasDashboardAPIGetAccountsResponse)
	if err != nil {
		return nil, xerrors.Errorf("error in createGCASAccount: %w", err)
	}

	err = u.deps.GCASAccountRepository.CreateMulti(ctx, gcasAccounts)
	if err != nil {
		return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeDBAccess, "error in createGCASAccount", err)
	}
	return gcasAccounts, nil
}

// toGCASAccountsFromGetAccountsResponse - GetAccountsResponseをGCASAccountに変換する
func (u *Usecase) toGCASAccountsFromGetAccountsResponse(eventDocID string, gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse) ([]*entities.GCASAccount, error) {
	result := []*entities.GCASAccount{}
	for _, accountIDs := range *gcasDashboardAPIGetAccountsResponse {
		for _, accountID := range accountIDs {
			uuid, err := u.deps.UUID.GetUUID()
			if err != nil {
				return nil, xerrors.Errorf("error in toGCASAccountsFromGetAccountsResponse: %w", err)
			}
			result = append(result, entities.NewGCASAccount(
				&entities.NewGCASAccountParam{
					ID:         uuid,
					EventDocID: eventDocID,
					AccountID:  accountID,
				},
			))
		}
	}
	return result, nil
}

// createGCASCSPCost - GCASCSPCostを登録する
func (u *Usecase) createGCASCSPCost(ctx context.Context, eventDocID string, gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse) error {
	cspAccountIDCostInfoMap, err := u.fetchCostInfo(gcasDashboardAPIGetAccountsResponse)
	if err != nil {
		return xerrors.Errorf("error in createGCASCSPCost: %w", err)
	}

	cspTotalCostMap := u.toCSPTotalCostMapFromCspAccountIDCostInfoMap(cspAccountIDCostInfoMap)

	gcasCSPCosts, err := u.toGCAPCSPCostsFromCostTotalCostMap(eventDocID, cspTotalCostMap)
	if err != nil {
		return xerrors.Errorf("error in createGCASCSPCost: %w", err)
	}

	err = u.deps.GCASCSPCostRepository.CreateMulti(ctx, gcasCSPCosts)
	if err != nil {
		return usecaseerrors.NewUnknownError(errorcode.ErrorCodeDBAccess, "error in createGCASCSPCost", err)
	}
	return nil
}

// fetchCostInfo - GCASダッシュボードAPIを実行しコスト情報を取得する。
func (u *Usecase) fetchCostInfo(gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse) (map[string]map[string]*gcasdashboardapi.GetCostResponse, error) {
	result := map[string]map[string]*gcasdashboardapi.GetCostResponse{}
	for csp, accountIDs := range *gcasDashboardAPIGetAccountsResponse {
		accountIDCostMap := map[string]*gcasdashboardapi.GetCostResponse{}
		for _, accountID := range accountIDs {
			gcasDashboardAPIGetCostResponse, err := u.deps.GCASDashboardAPI.GetCost(accountID)
			if err != nil {
				return nil, usecaseerrors.NewUnknownError(errorcode.ErrorCodeGCASDashboardAPI, "error in fetchCostInfo", err)
			}
			accountIDCostMap[accountID] = gcasDashboardAPIGetCostResponse
		}
		result[csp] = accountIDCostMap
	}
	return result, nil
}

// toCSPTotalCostMapFromCspAccountIDCostInfoMap - GCASダッシュボードAPIを実行しコスト情報を取得する。
func (u *Usecase) toCSPTotalCostMapFromCspAccountIDCostInfoMap(cspAccountIDCostInfoMap map[string]map[string]*gcasdashboardapi.GetCostResponse) map[string]int {
	result := map[string]int{}
	for csp, accountIDCostInfoMap := range cspAccountIDCostInfoMap {
		result[csp] = 0
		for _, constInfo := range accountIDCostInfoMap {
			result[csp] = result[csp] + constInfo.TotalCost
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
			return nil, xerrors.Errorf("error in toGCAPCSPCostsFromCostTotalCostMap: %w", err)
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
