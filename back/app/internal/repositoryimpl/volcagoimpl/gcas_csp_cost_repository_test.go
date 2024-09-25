//go:build !standalone

package volcagoimpl_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/volcagoimpl"
	"github.com/topgate/gcim-temporary/back/app/internal/testhelper"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago"
	"google.golang.org/api/iterator"
)

const (
	collectionName = "GCASCSPCost"
	eventID        = "202409251049"
	eventIDOther   = "202409251049other"
)

// sut.CreateMany(ctx, []*entities.GCASCSPCost{
// 	entities.NewGCASCSPCost(&entities.NewGCASCSPCostParam{
// 		ID:        "a",
// 		EventID:   "b",
// 		CSP:       "c",
// 		TotalCost: 1,
// 		Meta:      &entities.Meta{},
// 	}),
// })

// // CreateMany - 複数レコードを一括登録する
// func (g *gcasCSPCostImpl) CreateMany(ctx context.Context, gcasCSPCosts []*entities.GCASCSPCost) error {
// 	volcagoGCASCSPCost := make([]*volcago.GCASCSPCost, len(gcasCSPCosts))
// 	for i, gcasCSPCost := range gcasCSPCosts {
// 		volcagoGCASCSPCost[i] = &volcago.GCASCSPCost{
// 			ID:        gcasCSPCost.ID(),
// 			EventID:   gcasCSPCost.EventID(),
// 			CSP:       gcasCSPCost.CSP(),
// 			TotalCost: gcasCSPCost.TotalCost(),
// 			Meta: volcago.Meta{
// 				CreatedAt: gcasCSPCost.Meta().CreatedAt(),
// 				CreatedBy: gcasCSPCost.Meta().CreatedBy(),
// 				UpdatedAt: gcasCSPCost.Meta().UpdatedAt(),
// 				UpdatedBy: gcasCSPCost.Meta().UpdatedBy(),
// 				DeletedAt: gcasCSPCost.Meta().DeletedAt(),
// 				DeletedBy: gcasCSPCost.Meta().DeletedBy(),
// 			},
// 		}
// 	}
// 	_, err := g.infra.InsertMulti(ctx, volcagoGCASCSPCost)
// 	if err != nil {
// 		return repositoryerrors.NewUnknownError("failed to create gcas_csp_cost", err)
// 	}
// 	return nil
// }

// // Exists - event_idに紐付くコレクションの存在フラグを取得する
// func (g *gcasCSPCostImpl) Exists(ctx context.Context, eventID string) (bool, error) {
// 	chainer := infrastructures.NewQueryChainer
// 	param := &infrastructures.GCASCSPCostSearchParam{
// 		EventID:     chainer().Filters(eventID, infrastructures.FilterTypeAdd),
// 		CursorLimit: 0,
// 	}

// 	gcasCSPCosts, err := g.infra.Search(ctx, param, nil)
// 	if err != nil {
// 		return false, repositoryerrors.NewUnknownError("error in gcasCSPCostImpl.Exists", err)
// 	}

// 	return 0 < len(gcasCSPCosts), nil
// }

func TestGCASCSPCostImplExists(t *testing.T) {
	testhelper.LoadEnv(t)
	firestoreClient := testhelper.FirestoreClient(t)
	ctx := context.Background()

	tests := []struct {
		name      string
		prepareFn func(t *testing.T, firestoreClient *firestore.Client)
		want      bool
		wantErr   bool
	}{
		{
			name:      "正常系：データがない場合",
			prepareFn: func(t *testing.T, firestoreClient *firestore.Client) {},
			want:      false,
			wantErr:   false,
		},
		{
			name: "正常系：データが1件ある場合",
			prepareFn: func(t *testing.T, firestoreClient *firestore.Client) {
				addGCASPost(t, firestoreClient, &volcago.GCASCSPCost{
					ID:        "id" + eventID,
					EventID:   eventID,
					CSP:       "aws",
					TotalCost: 1000,
					Meta:      volcago.Meta{},
				})
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "正常系：データが2件ある場合",
			prepareFn: func(t *testing.T, firestoreClient *firestore.Client) {
				addGCASPost(t, firestoreClient, &volcago.GCASCSPCost{
					ID:        "id1" + eventID,
					EventID:   eventID,
					CSP:       "aws",
					TotalCost: 1000,
					Meta:      volcago.Meta{},
				})
				addGCASPost(t, firestoreClient, &volcago.GCASCSPCost{
					ID:        "id2" + eventID,
					EventID:   eventID,
					CSP:       "gcp",
					TotalCost: 2000,
					Meta:      volcago.Meta{},
				})
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "別のIDのデータがある場合",
			prepareFn: func(t *testing.T, firestoreClient *firestore.Client) {
				addGCASPost(t, firestoreClient, &volcago.GCASCSPCost{
					ID:        "id" + eventIDOther,
					EventID:   eventIDOther,
					CSP:       "aws",
					TotalCost: 1000,
					Meta:      volcago.Meta{},
				})
			},
			want:    false,
			wantErr: false,
		},
		// TODO 異常系：エラーが発生した場合
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteGCASPost(t, firestoreClient, eventID)
			deleteGCASPost(t, firestoreClient, eventIDOther)

			tt.prepareFn(t, firestoreClient)

			sut := volcagoimpl.NewGCASCSPCost(firestoreClient)
			got, err := sut.Exists(ctx, eventID)

			if (err != nil) != tt.wantErr {
				t.Fatalf("error :%v", err)
			}

			if got != tt.want {
				t.Fatalf("error : got = %t, want = %t", got, tt.want)
			}
		})
	}
}

func deleteGCASPost(t *testing.T, firestoreClient *firestore.Client, eventID string) {
	t.Helper()
	findGCASPost(t, firestoreClient, eventID, func(doc *firestore.DocumentSnapshot) {
		ctx := context.Background()
		doc.Ref.Delete(ctx)
	})
}

func findGCASPost(t *testing.T, firestoreClient *firestore.Client, eventID string, callback func(doc *firestore.DocumentSnapshot)) {
	t.Helper()

	ctx := context.Background()
	iter := firestoreClient.
		Collection(collectionName).
		Where("EventID", "==", eventID).
		Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		callback(doc)
	}
}

func addGCASPost(t *testing.T, firestoreClient *firestore.Client, gcasCSPCost *volcago.GCASCSPCost) {
	t.Helper()
	ctx := context.Background()
	_, err := firestoreClient.
		Collection(collectionName).
		Doc(gcasCSPCost.ID).
		Set(ctx, gcasCSPCost)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
