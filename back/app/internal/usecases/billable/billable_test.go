package billable_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasapi"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasdashboardapi"
	mockapi "github.com/topgate/gcim-temporary/back/app/internal/apiimpl/mocks"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	"github.com/topgate/gcim-temporary/back/app/internal/repositoryerrors"
	mockrepositories "github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/mocks"
	"github.com/topgate/gcim-temporary/back/app/internal/usecases/billable"
	"github.com/topgate/gcim-temporary/back/pkg/uuid"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
)

func Test_Usecase_Billable_正常系(t *testing.T) {
	sut, mock, deferFunc := NewSUT(t)
	defer deferFunc()

	ctx := context.Background()
	eventDocID := "eventDocID"
	csp := "aws"
	accountID := "11111"
	totalCost := 2222

	mock.MockEventStatusRepository.EXPECT().GetByEventDocIDAndStatus(ctx, eventDocID, entities.EventStatusStart).Return(
		entities.NewEventStatus(&entities.NewEventStatusParam{
			ID:         fmt.Sprintf("%s_%d", eventDocID, entities.EventStatusStart),
			EventDocID: eventDocID,
			Status:     entities.EventStatusStart,
			Meta:       &entities.Meta{},
		}),
		nil,
	)
	mock.MockEventStatusRepository.EXPECT().GetByEventDocIDAndStatus(ctx, eventDocID, entities.EventStatusInvoiceCreationChecked).Return(
		nil, repositoryerrors.NewNotFoundError(xerrors.New("error in GetByEventDocIDAndStatus")),
	)
	mock.MockGCASCSPCostRepository.EXPECT().Exists(ctx, eventDocID).Return(false, nil)
	mock.MockGCASDashboardAPI.EXPECT().GetAccounts().Return(&gcasdashboardapi.GetAccountsResponse{
		csp: []string{accountID},
	}, nil)
	mock.MockGCASAccountRepository.EXPECT().GetAccounts(ctx, eventDocID).Return([]*entities.GCASAccount{}, nil)
	mock.MockGCASAPI.EXPECT().GetAccounts().Return(&gcasapi.GetAccountsResponse{
		csp: []string{accountID},
	}, nil)
	mock.MockGCASCSPCostRepository.EXPECT().CreateMany(ctx, gomock.Any()).Return(nil)
	mock.MockGCASAccountRepository.EXPECT().CreateMany(ctx, gomock.Any()).Return(nil)
	mock.MockGCASDashboardAPI.EXPECT().GetCost(accountID).Return(&gcasdashboardapi.GetCostResponse{
		AccountID:  accountID,
		TotalCost:  totalCost,
		Identifier: make(map[string]int),
		Other:      0,
	}, nil)
	mock.MockEventStatusRepository.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	input := billable.Input{
		EventDocID: eventDocID,
	}
	output, err := sut.Billable(ctx, &input)
	if err != nil {
		t.Fatalf("error in Billable: %+v", err)
	}
	t.Logf("%v", output)

}

func Test_Usecase_Billable(t *testing.T) {
	sut, mock, deferFunc := NewSUT(t)
	defer deferFunc()

	ctx := context.Background()
	eventDocID := "eventDocID"

	mock.MockEventStatusRepository.EXPECT().GetByEventDocIDAndStatus(ctx, eventDocID, entities.EventStatusStart).Return(
		nil, repositoryerrors.NewNotFoundError(xerrors.New("error in GetByEventDocIDAndStatus")),
	)

	input := billable.Input{
		EventDocID: eventDocID,
	}
	output, err := sut.Billable(ctx, &input)
	if err != nil {
		t.Fatalf("error in Billable: %+v", err)
	}
	t.Logf("%v", output)
}

type Mock struct {
	MockGCASDashboardAPI      *mockapi.MockGCASDashboardAPI
	MockGCASAPI               *mockapi.MockGCASAPI
	MockEventStatusRepository *mockrepositories.MockEventStatusRepository
	MockGCASAccountRepository *mockrepositories.MockGCASAccountRepository
	MockGCASCSPCostRepository *mockrepositories.MockGCASCSPCostRepository
	deferFunc                 func()
}

func NewSUT(t *testing.T) (*billable.Usecase, *Mock, func()) {
	mockCtrlGCASDashboardAPI := gomock.NewController(t)
	mockCtrlGCASAPI := gomock.NewController(t)
	mockCtrlEventStatusRepository := gomock.NewController(t)
	mockCtrlGCASAccountRepository := gomock.NewController(t)
	mockCtrlGCASCSPCostRepository := gomock.NewController(t)
	deferFunc := func() {
		mockCtrlGCASDashboardAPI.Finish()
		mockCtrlGCASAPI.Finish()
		mockCtrlEventStatusRepository.Finish()
		mockCtrlGCASAccountRepository.Finish()
		mockCtrlGCASCSPCostRepository.Finish()
	}

	mockGCASDashboardAPI := mockapi.NewMockGCASDashboardAPI(mockCtrlGCASDashboardAPI)
	mockGCASAPI := mockapi.NewMockGCASAPI(mockCtrlGCASAPI)
	mockEventStatusRepository := mockrepositories.NewMockEventStatusRepository(mockCtrlEventStatusRepository)
	mockGCASAccountRepository := mockrepositories.NewMockGCASAccountRepository(mockCtrlGCASAccountRepository)
	mockGCASCSPCostRepository := mockrepositories.NewMockGCASCSPCostRepository(mockCtrlGCASCSPCostRepository)

	sut := billable.NewUsecase(billable.Dependencies{
		GCASDashboardAPI:      mockGCASDashboardAPI,
		GCASAPI:               mockGCASAPI,
		EventStatusRepository: mockEventStatusRepository,
		GCASAccountRepository: mockGCASAccountRepository,
		GCASCSPCostRepository: mockGCASCSPCostRepository,
		UUID:                  uuid.UUID{},
	})

	return sut, &Mock{
		MockGCASDashboardAPI:      mockGCASDashboardAPI,
		MockGCASAPI:               mockGCASAPI,
		MockEventStatusRepository: mockEventStatusRepository,
		MockGCASAccountRepository: mockGCASAccountRepository,
		MockGCASCSPCostRepository: mockGCASCSPCostRepository,
	}, deferFunc
}
