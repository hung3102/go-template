package invoiceunitlist

// Input - Input of usecase function
type Input struct {
	EventID string
}

// Output - Output of usecase
type Output struct {
	OrganizationCode string  // 団体コード
	OrganizationName string  // 団体名
	CSP              string  // csp
	AgencyName       *string // 事業者名 - 支払代行がある場合
	CorporateNumber  *string // 法人番号 - 支払代行がある場合
}
