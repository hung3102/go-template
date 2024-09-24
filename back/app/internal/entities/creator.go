package entities

// Creator - 請求書の発行元情報
type Creator struct {
	id           string // id
	eventID      string // event_id
	organization string // 団体名
	address      string // 住所
	personName   string // person_name
	meta         *Meta  // メタ
}

// NewCreatorParam - 請求書の発行元情報作成パラメータ
type NewCreatorParam struct {
	ID           string // id
	EventID      string // event_id
	Organization string // 団体名
	Address      string // 住所
	PersonName   string // person_name
	Meta         *Meta  // メタ
}

// NewCreator - 請求書の発行元情報作成
func NewCreator(param *NewCreatorParam) *Creator {
	return &Creator{
		id:           param.ID,
		eventID:      param.EventID,
		organization: param.Organization,
		address:      param.Address,
		personName:   param.PersonName,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *Creator) ID() string {
	return e.id
}

// EventID - EventID のゲッター
func (e *Creator) EventID() string {
	return e.eventID
}

// Organization - Organization のゲッター
func (e *Creator) Organization() string {
	return e.organization
}

// Address - Address のゲッター
func (e *Creator) Address() string {
	return e.address
}

// PersonName - PersonName のゲッター
func (e *Creator) PersonName() string {
	return e.personName
}

// Meta - Meta のゲッター
func (e *Creator) Meta() *Meta {
	return e.meta
}
