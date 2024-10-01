package entities

// PaymentAgency - 団体ごとのCSPごとのアカウントごと支払い代行者情報
type PaymentAgency struct {
	agencyName string // 事業者名
}

// NewPaymentAgencyParam - 団体ごとのCSPごとのアカウントごと支払い代行者情報作成パラメータ
type NewPaymentAgencyParam struct {
	AgencyName string // 事業者名
}

// NewPaymentAgency - 団体ごとのCSPごとのアカウントごと支払い代行者情報作成
func NewPaymentAgency(param *NewPaymentAgencyParam) *PaymentAgency {
	return &PaymentAgency{
		agencyName: param.AgencyName,
	}
}

// AgencyName - AgencyName Getter
func (e *PaymentAgency) AgencyName() string {
	return e.agencyName
}
