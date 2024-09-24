package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -mockgen ../../../../../bin/mockgen -mock-output mocks/creator_gen.go Creator

// Creator - 請求書の発行元情報
type Creator struct {
	ID           string `firestore:"-" firestore_key:""` // id
	EventID      string ``                               // event_id
	Organization string ``                               // 団体名
	Address      string ``                               // 住所
	PersonName   string ``                               // person_name
	Meta
}
