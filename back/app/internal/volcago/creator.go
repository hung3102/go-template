package volcago

//go:generate ../../../../bin/volcago -p infrastructures -o ./infrastructures -c creator -mockgen ../../../../../bin/mockgen -mock-output mocks/creator_gen.go Creator

// Creator - 請求書の発行元情報
type Creator struct {
	ID           string `firestore:"-" firestore_key:""` // id
	EventID      string `firestore:"event_id"`           // event_id
	Organization string `firestore:"organization"`       // 団体名
	Address      string `firestore:"address"`            // 住所
	PersonName   string `firestore:"person_name"`        // person_name
	Meta
}
