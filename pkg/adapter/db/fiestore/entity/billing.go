package entity

type Billing struct {
	DocumentId string `firestore:"-"`
	Cost       int    `firestore:"cost"`
}
