package accountlist_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasapi"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasdashboardapi"
	mockapi "github.com/topgate/gcim-temporary/back/app/internal/apiimpl/mocks"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	mockrepositories "github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/mocks"
	"github.com/topgate/gcim-temporary/back/app/internal/usecases/accountlist"
	"github.com/topgate/gcim-temporary/back/pkg/uuid"
	"go.uber.org/mock/gomock"
)

type Hoge struct {
	ID  string `json:"id"`
	Map map[string][]string
}

func Test_Usecase_AccountList_正常系(t *testing.T) {
	sut, mock := NewSUT(t)
	ctx := context.Background()
	eventDocID := "eventDocID"
	csp := "aws"
	accountID := "11111"
	totalCost := 2222

	mock.MockEventRepository.EXPECT().GetByID(ctx, gomock.Any()).Return(
		entities.NewEvent(&entities.NewEventParam{
			ID:           eventDocID,
			BillingMonth: time.Now(),
			Status:       strconv.Itoa(entities.EventStatusInvoiceCreationPossible),
			Meta:         &entities.Meta{},
		}),
		nil,
	)
	mock.MockGCASCSPCostRepository.EXPECT().Exists(ctx, eventDocID).Return(false, nil)
	mock.MockGCASDashboardAPI.EXPECT().GetAccounts().Return(&gcasdashboardapi.GetAccountsResponse{
		csp: []string{accountID},
	}, nil)
	mock.MockGCASAccountRepository.EXPECT().GetAccounts(ctx, eventDocID).Return([]*entities.GCASAccount{}, nil)
	mock.MockGCASAPI.EXPECT().GetAccounts().Return(&gcasapi.GetAccountsResponse{
		csp: []string{accountID},
	}, nil)
	mock.MockGCASCSPCostRepository.EXPECT().CreateMulti(ctx, gomock.Any()).Return(nil)

	input := accountlist.AccountListInput{
		EventDocID: eventDocID,
	}
	mock.MockGCASAccountRepository.EXPECT().CreateMulti(ctx, gomock.Any()).Return(nil)
	mock.MockGCASDashboardAPI.EXPECT().GetCost(accountID).Return(&gcasdashboardapi.GetCostResponse{
		AccountID:  accountID,
		TotalCost:  totalCost,
		Identifier: make(map[string]int),
		Other:      0,
	}, nil)
	output, err := sut.AccountList(ctx, &input)
	if err != nil {
		t.Fatalf("error in AccountList: %v", err)
	}
	t.Logf("%v", output)

}

func Test_Usecase_AccountList(t *testing.T) {
	sut, mock := NewSUT(t)
	eventDocID := "eventDocID"

	mock.MockEventRepository.EXPECT().GetByID(gomock.Any(), eventDocID).Return(
		entities.NewEvent(&entities.NewEventParam{
			ID:           eventDocID,
			BillingMonth: time.Now(),
			Status:       strconv.Itoa(entities.EventStatusStored),
			Meta:         &entities.Meta{},
		}),
		nil,
	)

	ctx := context.Background()
	input := accountlist.AccountListInput{
		EventDocID: eventDocID,
	}
	output, err := sut.AccountList(ctx, &input)
	if err != nil {
		t.Fatalf("error in AccountList: %v", err)
	}
	t.Logf("%v", output)

}

type Mock struct {
	MockGCASDashboardAPI      *mockapi.MockGCASDashboardAPI
	MockGCASAPI               *mockapi.MockGCASAPI
	MockEventRepository       *mockrepositories.MockBaseRepository[entities.Event]
	MockGCASAccountRepository *mockrepositories.MockGCASAccountRepository
	MockGCASCSPCostRepository *mockrepositories.MockGCASCSPCostRepository
}

func NewSUT(t *testing.T) (*accountlist.Usecase, *Mock) {
	mockCtrlGCASDashboardAPI := gomock.NewController(t)
	defer mockCtrlGCASDashboardAPI.Finish()
	mockGCASDashboardAPI := mockapi.NewMockGCASDashboardAPI(mockCtrlGCASDashboardAPI)

	mockCtrlGCASAPI := gomock.NewController(t)
	defer mockCtrlGCASAPI.Finish()
	mockGCASAPI := mockapi.NewMockGCASAPI(mockCtrlGCASAPI)

	mockCtrlEventRepository := gomock.NewController(t)
	defer mockCtrlEventRepository.Finish()
	mockEventRepository := mockrepositories.NewMockBaseRepository[entities.Event](mockCtrlEventRepository)

	mockCtrlGCASAccountRepository := gomock.NewController(t)
	defer mockCtrlGCASAccountRepository.Finish()
	mockGCASAccountRepository := mockrepositories.NewMockGCASAccountRepository(mockCtrlGCASAccountRepository)

	mockCtrlGCASCSPCostRepository := gomock.NewController(t)
	defer mockCtrlGCASCSPCostRepository.Finish()
	mockGCASCSPCostRepository := mockrepositories.NewMockGCASCSPCostRepository(mockCtrlGCASCSPCostRepository)

	sut := accountlist.NewUsecase(accountlist.Dependencies{
		GCASDashboardAPI:      mockGCASDashboardAPI,
		GCASAPI:               mockGCASAPI,
		EventsRepository:      mockEventRepository,
		GCASAccountRepository: mockGCASAccountRepository,
		GCASCSPCostRepository: mockGCASCSPCostRepository,
		UUID:                  uuid.UUID{},
	})

	return sut, &Mock{
		MockGCASDashboardAPI:      mockGCASDashboardAPI,
		MockGCASAPI:               mockGCASAPI,
		MockEventRepository:       mockEventRepository,
		MockGCASAccountRepository: mockGCASAccountRepository,
		MockGCASCSPCostRepository: mockGCASCSPCostRepository,
	}
}
