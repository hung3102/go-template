package entities

// InvoiceUnit - 請求書単位
type InvoiceUnit struct {
	isPaymentAgent bool   // 支払代行
	subject        string // 団体か事業者
	csp            string // csp
}

// NewInvoiceUnitParam - NewInvoiceUnit Param
type NewInvoiceUnitParam struct {
	IsPaymentAgent bool   // 支払代行
	Subject        string // 団体か事業者
	CSP            string // csp
}

// NewInvoiceUnit - Constructor of InvoiceUnit
func NewInvoiceUnit(param *NewInvoiceUnitParam) *InvoiceUnit {
	return &InvoiceUnit{
		isPaymentAgent: param.IsPaymentAgent,
		subject:        param.Subject,
		csp:            param.CSP,
	}
}

// IsPaymentAgent - IsPaymentAgent Getter
func (e *InvoiceUnit) IsPaymentAgent() bool {
	return e.isPaymentAgent
}

// Subject - Subject Getter
func (e *InvoiceUnit) Subject() string {
	return e.subject
}

// CSP - CSP Getter
func (e *InvoiceUnit) CSP() string {
	return e.csp
}
