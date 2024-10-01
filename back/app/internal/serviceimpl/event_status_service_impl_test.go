//go:build !standalone

package serviceimpl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/volcagoimpl"
	"github.com/topgate/gcim-temporary/back/app/internal/serviceimpl"
	"github.com/topgate/gcim-temporary/back/app/internal/testhelper"
	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago"
)

func TestEventStatusImpl(t *testing.T) {
	testhelper.SetEnv(t)
	firestoreClient := testhelper.FirestoreClient(t)

	ctx := context.Background()
	collectionName := "event_status"
	eventID := valueobjects.NewEventID()

	sut := serviceimpl.NewEventStatusService(&serviceimpl.NewEventStatusServiceParam{
		EventStatusRepository: volcagoimpl.NewEventStatus(firestoreClient),
	})

	// いったんイベントのデータ削除
	testhelper.DeleteDocsByEventID(t, firestoreClient, collectionName, eventID.String())

	// イベントがない場合は請求書の作成ができない
	got, err := sut.IsInvoiceCreatable(ctx, eventID)
	if err != nil {
		t.Fatalf("error sut.IsInvoiceCreatable: %+v", err)
	}
	if got != false {
		t.Fatalf("error sut.IsInvoiceCreatable: got = %t", got)
	}

	// 請求書作成イベントを登録
	eventStatusStart := volcago.EventStatus{
		ID:      fmt.Sprintf("%s_%d", eventID, entities.EventStatusStart),
		EventID: eventID.String(),
		Status:  entities.EventStatusStart,
		Meta:    volcago.Meta{},
	}
	testhelper.AddDoc(t, firestoreClient, collectionName, eventStatusStart.ID, eventStatusStart)

	// イベントがあるので請求書の作成ができる
	got, err = sut.IsInvoiceCreatable(ctx, eventID)
	if err != nil {
		t.Fatalf("error sut.IsInvoiceCreatable: %+v", err)
	}
	if got != true {
		t.Fatalf("error sut.IsInvoiceCreatable: got = %t", got)
	}

	// 請求書の作成が完了状態にする
	err = sut.SetBillable(ctx, eventID)
	if err != nil {
		t.Fatalf("error sut.SetBillable: %+v", err)
	}

	snapshot, err := firestoreClient.
		Collection(collectionName).
		Doc(fmt.Sprintf("%s_%d", eventID, entities.EventStatusInvoiceCreationChecked)).
		Get(ctx)
	if err != nil {
		t.Fatalf("error Get(ctx): %+v", err)
	}

	data := snapshot.Data()
	if data["event_id"] != eventID.String() || data["status"].(int64) != int64(entities.EventStatusInvoiceCreationChecked) {
		t.Fatalf("error data = %v", data)
	}

	// 請求書作成済なので、請求書の開始できない
	got, err = sut.IsInvoiceCreatable(ctx, eventID)
	if err != nil {
		t.Fatalf("error sut.IsInvoiceCreatable: %+v", err)
	}
	if got != false {
		t.Fatalf("error sut.IsInvoiceCreatable: got = %t", got)
	}

	// テストに使用したデータを削除
	testhelper.DeleteDocsByEventID(t, firestoreClient, collectionName, eventID.String())
}
