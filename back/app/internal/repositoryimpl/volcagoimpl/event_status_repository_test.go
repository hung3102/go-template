//go:build !standalone

package volcagoimpl_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryerrors"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/volcagoimpl"
	"github.com/topgate/gcim-temporary/back/app/internal/testhelper"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago"
)

func TestEventStatusImplGetByEventIDAndStatus(t *testing.T) {
	eventID := "202409251049"
	eventIDOther := "202409251049other"
	createdBy := "event_status_tester"

	type args struct {
		param *repositories.GetByEventIDAndStatusParam
	}
	tests := []struct {
		name          string
		prepareMockFn func(t *testing.T, firestoreClient *firestore.Client)
		args          args
		want          *entities.EventStatus
		wantErr       bool
		checkError    func(t *testing.T, err error)
	}{
		{
			name: "正常系：データが取れる場合",
			prepareMockFn: func(t *testing.T, firestoreClient *firestore.Client) {
				addEventStatus(t, firestoreClient, &volcago.EventStatus{
					ID:      fmt.Sprintf("%s_%d", eventID, 9),
					EventID: eventID,
					Status:  9,
					Meta: volcago.Meta{
						CreatedBy: createdBy,
						UpdatedBy: createdBy,
					},
				})
			},
			args: args{
				param: &repositories.GetByEventIDAndStatusParam{
					EventID: eventID,
					Status:  9,
				},
			},
			want: entities.NewEventStatus(&entities.NewEventStatusParam{
				EventID: eventID,
				Status:  9,
				Meta: entities.NewMeta(&entities.NewMetaParam{
					CreatedBy: createdBy,
					UpdatedBy: createdBy,
				}),
			}),
			wantErr:    false,
			checkError: func(t *testing.T, err error) {},
		},
		{
			name: "正常系：データが取れない場合",
			prepareMockFn: func(t *testing.T, firestoreClient *firestore.Client) {
				addEventStatus(t, firestoreClient, &volcago.EventStatus{
					ID:      fmt.Sprintf("%s_%d", eventIDOther, 9),
					EventID: eventIDOther, // event_idが異なる
					Status:  9,
					Meta: volcago.Meta{
						CreatedBy: createdBy,
						UpdatedBy: createdBy,
					},
				})
				addEventStatus(t, firestoreClient, &volcago.EventStatus{
					ID:      fmt.Sprintf("%s_%d", eventID, 8),
					EventID: eventID,
					Status:  8, // statusが異なる
					Meta: volcago.Meta{
						CreatedBy: createdBy,
						UpdatedBy: createdBy,
					},
				})
			},
			args: args{
				param: &repositories.GetByEventIDAndStatusParam{
					EventID: eventID,
					Status:  9,
				},
			},
			want:    nil,
			wantErr: true,
			checkError: func(t *testing.T, err error) {
				t.Helper()
				var e repositoryerrors.RepositoryError[repositoryerrors.NotFoundError]
				if !errors.As(err, &e) {
					t.Fatalf("err is not repositoryerrors.NotFoundError: %+v", err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetEnv(t)
			firestoreClient := testhelper.FirestoreClient(t)
			ctx := context.Background()

			deleteEventStatus(t, firestoreClient, eventID)
			deleteEventStatus(t, firestoreClient, eventIDOther)

			tt.prepareMockFn(t, firestoreClient)

			sut := volcagoimpl.NewEventStatus(firestoreClient)
			got, err := sut.GetByEventIDAndStatus(ctx, tt.args.param)

			if err != nil {
				if !tt.wantErr {
					t.Fatalf("error :%+v", err)
				}
				tt.checkError(t, err)
			}

			if tt.want == nil {
				if got != nil {
					t.Fatalf("error : got = %+v, want = nil", got)
				}
			} else {
				if got.ID() != tt.want.ID() ||
					got.EventID() != tt.want.EventID() ||
					got.Status() != tt.want.Status() ||
					got.Meta().CreatedBy() != tt.want.Meta().CreatedBy() ||
					got.Meta().UpdatedBy() != tt.want.Meta().UpdatedBy() ||
					got.Meta().DeletedBy() != "" ||
					got.Meta().DeletedAt() != nil {
					t.Fatalf("error: got = %v, want = %v", got, *tt.want)
				}
			}
		})
	}
}

func TestEventStatusImplCreate(t *testing.T) {
	eventID := "202409251049"
	createdBy := "event_status_tester"

	type args struct {
		eventStatus *entities.EventStatus
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "異常系:nilの場合",
			args: args{
				eventStatus: nil,
			},
		},
		{
			name: "正常系：1件登録",
			args: args{
				eventStatus: entities.NewEventStatus(&entities.NewEventStatusParam{
					EventID: eventID,
					Status:  2,
					Meta: entities.NewMeta(&entities.NewMetaParam{
						CreatedBy: createdBy,
						UpdatedBy: createdBy,
					}),
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testhelper.SetEnv(t)
			firestoreClient := testhelper.FirestoreClient(t)
			ctx := context.Background()

			deleteEventStatus(t, firestoreClient, eventID)

			before := findEventStatus(t, firestoreClient, eventID)
			if len(before) != 0 {
				t.Fatalf("error: len(before) = %d", len(before))
			}

			sut := volcagoimpl.NewEventStatus(firestoreClient)
			err := sut.Create(ctx, tt.args.eventStatus)
			if err != nil {
				t.Fatalf("error :%+v", err)
			}

			got := findEventStatus(t, firestoreClient, eventID)

			if tt.args.eventStatus == nil {
				if len(got) > 0 {
					t.Fatalf("error: len(got) = %d, want = nil", len(got))
				}
				return
			}

			if len(got) != 1 {
				t.Fatalf("error: len(got) = %d, want = nil", len(got))
			}

			data := got[0].Data()
			want := tt.args.eventStatus
			if got[0].Ref.ID != fmt.Sprintf("%s_%d", want.EventID(), want.Status()) ||
				data["event_id"] != want.EventID() ||
				data["status"].(int64) != int64(want.Status()) ||
				data["created_by"] != want.Meta().CreatedBy() ||
				data["updated_by"] != want.Meta().UpdatedBy() ||
				data["deleted_by"] != "" ||
				data["deleted_at"] != nil {
				t.Fatalf("error: got[0] = %v, want = %v", data, *want)
			}
		})
	}
}

func deleteEventStatus(t *testing.T, firestoreClient *firestore.Client, eventID string) {
	t.Helper()
	collectionName := "event_status"
	testhelper.DeleteDocsByEventID(t, firestoreClient, collectionName, eventID)
}

func findEventStatus(t *testing.T, firestoreClient *firestore.Client, eventID string) []*firestore.DocumentSnapshot {
	t.Helper()
	collectionName := "event_status"
	return testhelper.FindDocsByEventID(t, firestoreClient, collectionName, eventID)
}

func addEventStatus(t *testing.T, firestoreClient *firestore.Client, data *volcago.EventStatus) {
	t.Helper()
	collectionName := "event_status"
	testhelper.Add(t, firestoreClient, collectionName, data.ID, data)
}
