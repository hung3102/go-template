package invoiceunitlist

// Input - Input of usecase function
type Input struct {
	EventID string
}

// Output - Output of usecase
type Output struct {
	IsPaymentAgent bool   // 支払代行
	Subject        string // 団体か事業者
	CSP            string // csp
}
