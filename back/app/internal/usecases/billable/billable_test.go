package billable_test

import (
	"context"
	"testing"

	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasapi"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasdashboardapi"
	mockapi "github.com/topgate/gcim-temporary/back/app/internal/apiimpl/mocks"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	mockrepositories "github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/mocks"
	mockservices "github.com/topgate/gcim-temporary/back/app/internal/serviceimpl/mocks"
	"github.com/topgate/gcim-temporary/back/app/internal/usecases/billable"
	"github.com/topgate/gcim-temporary/back/pkg/uuid"
	"go.uber.org/mock/gomock"
)

func Test_Usecase_Billable_正常系(t *testing.T) {
	sut, mock, deferFunc := NewSUT(t)
	defer deferFunc()

	ctx := context.Background()
	eventDocID := "eventDocID"
	csp := "aws"
	accountID := "11111"
	totalCost := 2222

	mock.MockEventStatusService.EXPECT().ShouldcreateInvoice(ctx, eventDocID).Return(true, nil)
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
	mock.MockEventStatusService.EXPECT().SetInvoiceCreationChecked(ctx, eventDocID).Return(nil)

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

	mock.MockEventStatusService.EXPECT().ShouldcreateInvoice(ctx, eventDocID).Return(false, nil)

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
	MockEventStatusService    *mockservices.MockEventStatusService
	MockGCASAccountRepository *mockrepositories.MockGCASAccountRepository
	MockGCASCSPCostRepository *mockrepositories.MockGCASCSPCostRepository
	deferFunc                 func()
}

func NewSUT(t *testing.T) (*billable.Usecase, *Mock, func()) {
	mockCtrlGCASDashboardAPI := gomock.NewController(t)
	mockCtrlGCASAPI := gomock.NewController(t)
	mockCtrlEventStatusService := gomock.NewController(t)
	mockCtrlGCASAccountRepository := gomock.NewController(t)
	mockCtrlGCASCSPCostRepository := gomock.NewController(t)
	deferFunc := func() {
		mockCtrlGCASDashboardAPI.Finish()
		mockCtrlGCASAPI.Finish()
		mockCtrlEventStatusService.Finish()
		mockCtrlGCASAccountRepository.Finish()
		mockCtrlGCASCSPCostRepository.Finish()
	}

	mockGCASDashboardAPI := mockapi.NewMockGCASDashboardAPI(mockCtrlGCASDashboardAPI)
	mockGCASAPI := mockapi.NewMockGCASAPI(mockCtrlGCASAPI)
	mockEventStatusService := mockservices.NewMockEventStatusService(mockCtrlEventStatusService)
	mockGCASAccountRepository := mockrepositories.NewMockGCASAccountRepository(mockCtrlGCASAccountRepository)
	mockGCASCSPCostRepository := mockrepositories.NewMockGCASCSPCostRepository(mockCtrlGCASCSPCostRepository)

	sut := billable.NewUsecase(billable.Dependencies{
		GCASDashboardAPI:      mockGCASDashboardAPI,
		GCASAPI:               mockGCASAPI,
		EventStatusService:    mockEventStatusService,
		GCASAccountRepository: mockGCASAccountRepository,
		GCASCSPCostRepository: mockGCASCSPCostRepository,
		UUID:                  uuid.UUID{},
	})

	return sut, &Mock{
		MockGCASDashboardAPI:      mockGCASDashboardAPI,
		MockGCASAPI:               mockGCASAPI,
		MockEventStatusService:    mockEventStatusService,
		MockGCASAccountRepository: mockGCASAccountRepository,
		MockGCASCSPCostRepository: mockGCASCSPCostRepository,
	}, deferFunc
}
