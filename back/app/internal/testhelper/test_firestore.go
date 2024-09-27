package testhelper

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// FirestoreClient - firestoreClientを取得する
func FirestoreClient(t *testing.T) *firestore.Client {
	t.Helper()

	ctx := context.Background()
	projectID := os.Getenv("FIRESTORE_PROJECT_ON_EMULATOR")
	options := make([]option.ClientOption, 0)
	options = append(options, option.WithoutAuthentication())
	result, err := firestore.NewClient(ctx, projectID, options...)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	return result
}

// DeleteDocsByEventID - 指定したEventIDのコレクションを削除する
func DeleteDocsByEventID(t *testing.T, firestoreClient *firestore.Client, collectionName string, eventID string) {
	t.Helper()
	ctx := context.Background()

	for _, doc := range FindDocsByEventID(t, firestoreClient, collectionName, eventID) {
		doc.Ref.Delete(ctx)
	}
}

// FindDocsByEventID - 指定したEventIDのコレクションのリストを取得する
func FindDocsByEventID(t *testing.T, firestoreClient *firestore.Client, collectionName string, eventID string) []*firestore.DocumentSnapshot {
	t.Helper()
	result := make([]*firestore.DocumentSnapshot, 0)
	ctx := context.Background()
	iter := firestoreClient.Collection(collectionName).Where("event_id", "==", eventID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		result = append(result, doc)
	}
	return result
}

// AddDoc - DBにデータを登録する
func AddDoc[T any](t *testing.T, firestoreClient *firestore.Client, collectionName string, id string, data T) {
	t.Helper()
	ctx := context.Background()
	_, err := firestoreClient.
		Collection(collectionName).
		Doc(id).
		Set(ctx, data)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
