// Package - GCASダッシュボードAPIを実行
package gcasdashboardapi

//go:generate ../../../../../bin/mockgen -source=$GOFILE -destination=../../apiimpl/mocks/gcas_dashboard_api_mock.go -package=mockapi

// GCASダッシュボードAPIを実行
type GCASDashboardAPI interface {
	// アカウント一覧を取得する。
	GetAccounts() (*GetAccountsResponse, error)
	// 指定したアカウントのコスト情報を取得する。
	GetCost(csp string, accountID string) (*GetCostResponse, error)
}

/*
GetAccountsResponse - アカウント一覧

	{
		"aws": ["1111", "2222"],
		"azure": ["3333", "4444"],
		...
	}
*/
type GetAccountsResponse map[string][]string

// GetCostResponse - コスト情報
type GetCostResponse struct {
	AccountID  string         // アカウントID
	TotalCost  int            // 合計コスト
	Identifier map[string]int // 詳細コスト
	Other      int            // その他のコスト
}
