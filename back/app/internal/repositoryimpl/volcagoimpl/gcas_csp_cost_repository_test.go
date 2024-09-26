//go:build !standalone

package volcagoimpl_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/volcagoimpl"
	"github.com/topgate/gcim-temporary/back/app/internal/testhelper"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago"
)

func TestGCASCSPCostImplCreateMany(t *testing.T) {
	eventID := "202409251049"
	createdBy := "gcas_csp_cost_tester"

	type args struct {
		gcasCspCosts []*entities.GCASCSPCost
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "正常系：データがない場合",
			args: args{
				gcasCspCosts: []*entities.GCASCSPCost{},
			},
		},
		{
			name: "正常系：データが1件ある場合",
			args: args{
				gcasCspCosts: []*entities.GCASCSPCost{
					entities.NewGCASCSPCost(&entities.NewGCASCSPCostParam{
						ID:        "id1" + eventID,
						EventID:   eventID,
						CSP:       "aws",
						TotalCost: 123,
						Meta: entities.NewMeta(&entities.NewMetaParam{
							CreatedBy: createdBy,
							UpdatedBy: createdBy,
						}),
					}),
				},
			},
		},
		{
			name: "正常系：データが2件ある場合",
			args: args{
				gcasCspCosts: []*entities.GCASCSPCost{
					entities.NewGCASCSPCost(&entities.NewGCASCSPCostParam{
						ID:        "id1" + eventID,
						EventID:   eventID,
						CSP:       "aws",
						TotalCost: 123,
						Meta: entities.NewMeta(&entities.NewMetaParam{
							CreatedBy: createdBy,
							UpdatedBy: createdBy,
						}),
					}),
					entities.NewGCASCSPCost(&entities.NewGCASCSPCostParam{
						ID:        "id2" + eventID,
						EventID:   eventID,
						CSP:       "gcp",
						TotalCost: 123,
						Meta: entities.NewMeta(&entities.NewMetaParam{
							CreatedBy: createdBy,
							UpdatedBy: createdBy,
						}),
					}),
				},
			},
		},
		{
			name: "正常系：データがnilの場合",
			args: args{
				gcasCspCosts: []*entities.GCASCSPCost{
					nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetEnv(t)
			firestoreClient := testhelper.FirestoreClient(t)
			ctx := context.Background()

			deleteGCASCSPCost(t, firestoreClient, eventID)

			before := findGCASCSPCost(t, firestoreClient, eventID)
			if len(before) != 0 {
				t.Fatalf("error: len(before) = %d", len(before))
			}

			sut := volcagoimpl.NewGCASCSPCost(firestoreClient)
			err := sut.CreateMany(ctx, tt.args.gcasCspCosts)
			if err != nil {
				t.Fatalf("error :%+v", err)
			}

			idGCASCSPCostMap := make(map[string]*entities.GCASCSPCost)
			for _, gcasCspCost := range tt.args.gcasCspCosts {
				if gcasCspCost == nil {
					continue
				}
				idGCASCSPCostMap[gcasCspCost.ID()] = gcasCspCost
			}

			got := findGCASCSPCost(t, firestoreClient, eventID)
			if len(got) != len(idGCASCSPCostMap) {
				t.Fatalf("error: len(got) = %d, len(idGCASCSPCostMap) = %d", len(got), len(idGCASCSPCostMap))
			}

			for i, g := range got {
				data := g.Data()

				gcasCSPCost, ok := idGCASCSPCostMap[g.Ref.ID]
				if !ok {
					t.Fatalf("error: got[%d] = %v, gcasCSPCost = nil", i, data)
				}
				if data["event_id"] != gcasCSPCost.EventID() ||
					data["csp"] != gcasCSPCost.CSP() ||
					data["total_cost"].(int64) != int64(gcasCSPCost.TotalCost()) ||
					data["created_by"] != gcasCSPCost.Meta().CreatedBy() ||
					data["updated_by"] != gcasCSPCost.Meta().UpdatedBy() ||
					data["deleted_by"] != "" ||
					data["deleted_at"] != nil {
					t.Fatalf("error: got[%d] = %v, gcasCSPCost = %v", i, data, *gcasCSPCost)
				}
			}
		})
	}
}

func TestGCASCSPCostImplExists(t *testing.T) {
	eventID := "202409251049"
	eventIDOther := "202409251049other"

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
				addGCASCSPCost(t, firestoreClient, &volcago.GCASCSPCost{
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
				addGCASCSPCost(t, firestoreClient, &volcago.GCASCSPCost{
					ID:        "id1" + eventID,
					EventID:   eventID,
					CSP:       "aws",
					TotalCost: 1000,
					Meta:      volcago.Meta{},
				})
				addGCASCSPCost(t, firestoreClient, &volcago.GCASCSPCost{
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
				addGCASCSPCost(t, firestoreClient, &volcago.GCASCSPCost{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetEnv(t)
			firestoreClient := testhelper.FirestoreClient(t)
			ctx := context.Background()

			deleteGCASCSPCost(t, firestoreClient, eventID)
			deleteGCASCSPCost(t, firestoreClient, eventIDOther)

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

func deleteGCASCSPCost(t *testing.T, firestoreClient *firestore.Client, eventID string) {
	t.Helper()
	collectionName := "gcas_csp_cost"
	testhelper.DeleteDocsByEventID(t, firestoreClient, collectionName, eventID)
}

func findGCASCSPCost(t *testing.T, firestoreClient *firestore.Client, eventID string) []*firestore.DocumentSnapshot {
	t.Helper()
	collectionName := "gcas_csp_cost"
	return testhelper.FindDocsByEventID(t, firestoreClient, collectionName, eventID)
}

func addGCASCSPCost(t *testing.T, firestoreClient *firestore.Client, gcasCost *volcago.GCASCSPCost) {
	t.Helper()
	collectionName := "gcas_csp_cost"
	testhelper.Add(t, firestoreClient, collectionName, gcasCost.ID, gcasCost)
}
