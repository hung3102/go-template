package volcago

// PaymentAgency - 団体ごとのCSPごとのアカウントごと支払い代行者情報
type PaymentAgency struct {
	AgencyName      string `firestore:"agency_name"`      // 事業者名
	CorporateNumber string `firestore:"corporate_number"` // 法人番号
}
