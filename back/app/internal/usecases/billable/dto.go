package billable

// Input - 請求書作成の開始判定のinput
type Input struct {
	EventID string
}

// Output - 請求書作成の開始判定のoutput
type Output struct {
	GCASAccounts []*OutputAccount
}

// OutputAccount - 請求書作成の開始判定のoutputのアカウント情報
type OutputAccount struct {
	CSP       string // CSP
	AccountID string // アカウントID
}
