//go:build !standalone

package volcagoimpl_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/topgate/gcim-temporary/back/app/internal/repositories"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/volcagoimpl"
	"github.com/topgate/gcim-temporary/back/app/internal/testhelper"
	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
	"github.com/topgate/gcim-temporary/back/app/internal/volcago"
)

func TestOrgCSPAccountCostSearchByParam_HasPagingResult(t *testing.T) {
	testhelper.SetEnv(t)
	firestoreClient := testhelper.FirestoreClient(t)
	eventID := valueobjects.NewEventID()

	t.Run("Has PagingResult", func(t *testing.T) {
		ctx := context.Background()
		// init all data
		addOrgCSPAccountCost(t, firestoreClient, &volcago.OrgCSPAccountCost{
			ID:                   valueobjects.NewOrgCSPAccountCostID().String(),
			CSP:                  "AWS",
			EventID:              eventID.String(),
			GCASProportionCostID: valueobjects.NewGCASProportionCostID().String(),
			GCASAccountCostID:    valueobjects.NewGCASAccountCostID().String(),
			PaymentAgency: &volcago.PaymentAgency{
				AgencyName:      "agencyName1",
				CorporateNumber: "corporateName1",
			},
		})
		addOrgCSPAccountCost(t, firestoreClient, &volcago.OrgCSPAccountCost{
			ID:                   valueobjects.NewOrgCSPAccountCostID().String(),
			CSP:                  "AWS",
			EventID:              eventID.String(),
			GCASProportionCostID: valueobjects.NewGCASProportionCostID().String(),
			GCASAccountCostID:    valueobjects.NewGCASAccountCostID().String(),
			PaymentAgency: &volcago.PaymentAgency{
				AgencyName:      "agencyName2",
				CorporateNumber: "corporateName2",
			},
		})
		addOrgCSPAccountCost(t, firestoreClient, &volcago.OrgCSPAccountCost{
			ID:                   valueobjects.NewOrgCSPAccountCostID().String(),
			CSP:                  "AWS",
			EventID:              eventID.String(),
			GCASProportionCostID: valueobjects.NewGCASProportionCostID().String(),
			GCASAccountCostID:    valueobjects.NewGCASAccountCostID().String(),
			PaymentAgency: &volcago.PaymentAgency{
				AgencyName:      "agencyName3",
				CorporateNumber: "corporateName3",
			},
		})

		impl := volcagoimpl.NewOrgCSPAccountCost(firestoreClient)
		searchParam := repositories.OrgCSPAccountCostSearchParam{
			EventID: eventID,
			Limit:   2,
		}
		result, paging, err := impl.SearchByParam(ctx, &searchParam)

		if err != nil {
			t.Fatalf("error :%+v", err)
		}

		if len(result) != 2 {
			t.Fatalf("expected result has length of %d, got %d", 2, len(result))
		}

		if paging == nil {
			t.Fatal("expected paging is not nil, got nil")
		}

		firstAgencyName := result[0].PaymentAgency().AgencyName()
		if firstAgencyName != "agencyName1" {
			t.Fatalf("expected first agencyName is %v, got %v", "agencyName1", firstAgencyName)
		}

		t.Log(len(result), paging)
	})

	// deleteOrgCSPAccountCost(t, firestoreClient, eventID)
}

func addOrgCSPAccountCost(t *testing.T, firestoreClient *firestore.Client, doc *volcago.OrgCSPAccountCost) {
	t.Helper()
	collectionName := "org_csp_account_cost"
	testhelper.AddDoc(t, firestoreClient, collectionName, doc.ID, doc)
}

func deleteOrgCSPAccountCost(t *testing.T, firestoreClinet *firestore.Client, eventID valueobjects.EventID) {
	t.Helper()
	collectionName := "org_csp_account_cost"
	testhelper.DeleteDocsByEventID(t, firestoreClinet, collectionName, eventID.String())
}
