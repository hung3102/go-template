// Package - GCAS APIを実行
package gcasapi

//go:generate ../../../../../bin/mockgen -source=$GOFILE -destination=../../apiimpl/mocks/gcas_api_mock.go -package=mockapi

// GCASAPI - GCAS APIを実行
type GCASAPI interface {
	// GetAccounts - アカウント一覧を取得する
	GetAccounts() (*GetAccountsResponse, error)
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
