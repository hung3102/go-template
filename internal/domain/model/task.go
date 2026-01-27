package model

import "time"

//go:generate volcago -o ./../repositories -p repositories Task
type Task struct {
	ID      string          `firestore:"-"           firestore_key:""`
	Desc    string          `firestore:"description" indexer:"suffix,like" unique:""`
	Done    bool            `firestore:"done"        indexer:"equal"`
	Count   int             `firestore:"count"`
	Created time.Time       `firestore:"created"`
	Indexes map[string]bool `firestore:"indexes"`
}
