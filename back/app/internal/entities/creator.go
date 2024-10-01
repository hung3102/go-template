package entities

import "github.com/topgate/gcim-temporary/back/app/internal/valueobjects"

// Creator - 請求書の発行元情報
type Creator struct {
	id           valueobjects.CreatorID // ID
	eventID      valueobjects.EventID   // イベントID
	organization string                 // 団体名
	address      string                 // 住所
	personName   string                 // 氏名
	meta         *Meta                  // メタ
}

// NewCreatorParam - 請求書の発行元情報作成パラメータ
type NewCreatorParam struct {
	ID           valueobjects.CreatorID // ID
	EventID      valueobjects.EventID   // イベントID
	Organization string                 // 団体名
	Address      string                 // 住所
	PersonName   string                 // 氏名
	Meta         *Meta                  // メタ
}

// NewCreator - 請求書の発行元情報作成
func NewCreator(param *NewCreatorParam) *Creator {
	id := param.ID
	if id.IsNil() {
		id = valueobjects.NewCreatorID()
	}
	return &Creator{
		id:           id,
		eventID:      param.EventID,
		organization: param.Organization,
		address:      param.Address,
		personName:   param.PersonName,
		meta:         param.Meta,
	}
}

// ID - ID のゲッター
func (e *Creator) ID() valueobjects.CreatorID {
	return e.id
}

// EventID - EventID のゲッター
func (e *Creator) EventID() valueobjects.EventID {
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
