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
func (i *InvoiceUnit) IsPaymentAgent() bool {
	return i.isPaymentAgent
}

// Subject - Subject Getter
func (i *InvoiceUnit) Subject() string {
	return i.subject
}

// CSP - CSP Getter
func (i *InvoiceUnit) CSP() string {
	return i.csp
}
