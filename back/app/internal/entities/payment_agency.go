package entities

// PaymentAgency - 団体ごとのCSPごとのアカウントごと支払い代行者情報
type PaymentAgency struct {
	agencyName      string // 事業者名
	corporateNumber string // 法人番号
}

// NewPaymentAgencyParam - 団体ごとのCSPごとのアカウントごと支払い代行者情報作成パラメータ
type NewPaymentAgencyParam struct {
	AgencyName      string // 事業者名
	CorporateNumber string // 法人番号
}

// NewPaymentAgency - 団体ごとのCSPごとのアカウントごと支払い代行者情報作成
func NewPaymentAgency(param *NewPaymentAgencyParam) *PaymentAgency {
	return &PaymentAgency{
		agencyName:      param.AgencyName,
		corporateNumber: param.CorporateNumber,
	}
}

// AgencyName - AgencyName Getter
func (e *PaymentAgency) AgencyName() string {
	return e.agencyName
}

// CorporateNumber - CorporateNumber Getter
func (e *PaymentAgency) CorporateNumber() string {
	return e.corporateNumber
}
