package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/event_status_gen.go EventStatus

// EventStatus - イベントステータス
type EventStatus struct {
	ID      string `firestore:"-" firestore_key:""` // {event_id}_{status}
	EventID string ``                               // event_id
	Status  int    ``                               // ステータス
	Meta
}
