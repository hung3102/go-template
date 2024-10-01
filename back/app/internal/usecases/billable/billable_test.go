package billable_test

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasapi"
	"github.com/topgate/gcim-temporary/back/app/internal/api/gcasdashboardapi"
	mockapi "github.com/topgate/gcim-temporary/back/app/internal/apiimpl/mocks"
	"github.com/topgate/gcim-temporary/back/app/internal/entities"
	mockrepositories "github.com/topgate/gcim-temporary/back/app/internal/repositoryimpl/mocks"
	mockservices "github.com/topgate/gcim-temporary/back/app/internal/serviceimpl/mocks"
	"github.com/topgate/gcim-temporary/back/app/internal/usecases/billable"
	"github.com/topgate/gcim-temporary/back/app/internal/valueobjects"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
)

func TestUsecaseBillable正常系(t *testing.T) {
	sut, mock, deferFunc := NewSUT(t)
	defer deferFunc()

	ctx := context.Background()
	eventID := valueobjects.NewEventID()
	csp := "aws"
	accountID := "11111"
	totalCost := 2222

	mock.MockEventStatusService.EXPECT().IsInvoiceCreatable(ctx, eventID).Return(true, nil)
	mock.MockGCASCSPCostRepository.EXPECT().Exists(ctx, eventID).Return(false, nil)
	mock.MockGCASDashboardAPI.EXPECT().GetAccounts().Return(&gcasdashboardapi.GetAccountsResponse{
		csp: []string{accountID},
	}, nil)
	mock.MockGCASAPI.EXPECT().GetAccounts().Return(&gcasapi.GetAccountsResponse{
		csp: []string{accountID},
	}, nil)
	mock.MockGCASCSPCostRepository.EXPECT().CreateMany(ctx, gomock.Any()).Return(nil)
	mock.MockGCASDashboardAPI.EXPECT().GetCost(csp, accountID).Return(&gcasdashboardapi.GetCostResponse{
		AccountID:  accountID,
		TotalCost:  totalCost,
		Identifier: make(map[string]int),
		Other:      0,
	}, nil)
	mock.MockEventStatusService.EXPECT().SetBillable(ctx, eventID).Return(nil)

	input := billable.Input{
		EventID: eventID.String(),
	}
	output, err := sut.Billable(ctx, &input)
	if err != nil {
		t.Fatalf("error in Billable: %+v", err)
	}

	t.Logf("%v", output)
}

func TestUsecaseBillable請求書開始判定済の場合(t *testing.T) {
	sut, mock, deferFunc := NewSUT(t)
	defer deferFunc()

	ctx := context.Background()
	eventID := valueobjects.NewEventID()

	mock.MockEventStatusService.EXPECT().IsInvoiceCreatable(ctx, eventID).Return(false, nil)

	input := billable.Input{
		EventID: eventID.String(),
	}
	output, err := sut.Billable(ctx, &input)
	if err != nil {
		t.Fatalf("error in Billable: %+v", err)
	}
	t.Logf("%v", output)
}

func TestUsecaseCompareAccountInfo(t *testing.T) {
	type args struct {
		gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse
		gcasAPIGetAccountsResponse          *gcasapi.GetAccountsResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系：0件の場合",
			args: args{
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{},
				gcasAPIGetAccountsResponse:          &gcasapi.GetAccountsResponse{},
			},
			wantErr: false,
		},
		{
			name: "正常系：CSP、アカウントが一致する場合",
			args: args{
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{
					"aws":   {"1111", "2222", "3333"},
					"gcp":   {"4444", "1111"},
					"azure": {"2222"},
					"oci":   {},
				},
				gcasAPIGetAccountsResponse: &gcasapi.GetAccountsResponse{
					"gcp":   {"1111", "4444"}, //awsとgcpの順を変えておく
					"aws":   {"1111", "3333", "2222"},
					"azure": {"2222"},
					"oci":   {},
				},
			},
			wantErr: false,
		},
		{
			name: "異常系：CSPの件数が違う場合",
			args: args{
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{
					"aws":   {"1111", "2222", "3333"},
					"gcp":   {"1111", "4444"},
					"azure": {"2222"},
					"oci":   {},
				},
				gcasAPIGetAccountsResponse: &gcasapi.GetAccountsResponse{
					"aws": {"1111", "2222", "3333"},
					// gcpがない
					"azure": {"2222"},
					"oci":   {},
				},
			},
			wantErr: true,
		},
		{
			name: "異常系：CSPの件数が一致するが内容が違う場合",
			args: args{
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{
					"aws":   {"1111", "2222", "3333"},
					"gcp":   {"1111", "4444"},
					"azure": {"2222"},
				},
				gcasAPIGetAccountsResponse: &gcasapi.GetAccountsResponse{
					"aws":   {"1111", "2222", "3333"},
					"oci":   {"1111", "4444"}, // gcpとociの違い
					"azure": {"2222"},
				},
			},
			wantErr: true,
		},
		{
			name: "異常系：CSPの件数が一致し、アカウント件数が違う場合",
			args: args{
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{
					"aws":   {"1111", "2222", "3333"},
					"gcp":   {"1111", "4444"},
					"azure": {"2222"},
					"oci":   {},
				},
				gcasAPIGetAccountsResponse: &gcasapi.GetAccountsResponse{
					"aws":   {"1111", "2222", "3333"},
					"gcp":   {"1111", "4444"},
					"azure": {"2222", "5555"}, // アカウント件数が違う
					"oci":   {},
				},
			},
			wantErr: true,
		},
		{
			name: "異常系：CSPの件数、アカウント件数が一致するが、アカウントの内容が違う場合",
			args: args{
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{
					"aws":   {"1111", "2222", "3333"},
					"gcp":   {"1111", "4444"},
					"azure": {"2222"},
					"oci":   {},
				},
				gcasAPIGetAccountsResponse: &gcasapi.GetAccountsResponse{
					"aws":   {"1111", "2222", "3333"},
					"gcp":   {"1111", "3333"}, // アカウントIDが違う
					"azure": {"2222"},
					"oci":   {},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut, _, deferFunc := NewSUT(t)
			defer deferFunc()

			got := sut.CompareAccountInfo(tt.args.gcasDashboardAPIGetAccountsResponse, tt.args.gcasAPIGetAccountsResponse)

			if (got != nil) != tt.wantErr {
				t.Fatalf("error in Test_Usecase_ToOutputFromGCASAccount: name = %s, got = %+v", tt.name, got)
			}
		})
	}
}

func TestUsecaseToGCASCSPCost(t *testing.T) {
	eventID := valueobjects.NewEventID()

	type args struct {
		eventID                             valueobjects.EventID
		gcasDashboardAPIGetAccountsResponse *gcasdashboardapi.GetAccountsResponse
	}
	tests := []struct {
		name          string
		prepareMockFn func(mock *Mock)
		args          args
		want          []*entities.GCASCSPCost
		wantErr       bool
	}{
		{
			name:          "正常系：0件の場合",
			prepareMockFn: func(mock *Mock) {},
			args: args{
				eventID:                             eventID,
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{},
			},
			want:    []*entities.GCASCSPCost{},
			wantErr: false,
		},
		{
			name: "正常系：CSP複数、アカウント複数",
			prepareMockFn: func(mock *Mock) {
				mock.MockGCASDashboardAPI.EXPECT().GetCost("aws", "1111").Return(&gcasdashboardapi.GetCostResponse{AccountID: "1111", TotalCost: 1111}, nil)
				mock.MockGCASDashboardAPI.EXPECT().GetCost("aws", "2222").Return(&gcasdashboardapi.GetCostResponse{AccountID: "2222", TotalCost: 111}, nil)
				mock.MockGCASDashboardAPI.EXPECT().GetCost("gcp", "1111").Return(&gcasdashboardapi.GetCostResponse{AccountID: "1111", TotalCost: 11}, nil)
			},
			args: args{
				eventID: eventID,
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{
					"aws":   {"1111", "2222"},
					"gcp":   {"1111"},
					"azure": {},
				},
			},
			want: []*entities.GCASCSPCost{
				entities.NewGCASCSPCost(&entities.NewGCASCSPCostParam{EventID: eventID, CSP: "aws", TotalCost: 1222}),
				entities.NewGCASCSPCost(&entities.NewGCASCSPCostParam{EventID: eventID, CSP: "azure", TotalCost: 0}),
				entities.NewGCASCSPCost(&entities.NewGCASCSPCostParam{EventID: eventID, CSP: "gcp", TotalCost: 11}),
			},
			wantErr: false,
		},
		{
			name: "異常系：アカウントに紐付くコスト情報が存在しない",
			prepareMockFn: func(mock *Mock) {
				mock.MockGCASDashboardAPI.EXPECT().GetCost("aws", "1111").Return(nil, xerrors.New("mock api error"))
			},
			args: args{
				eventID: eventID,
				gcasDashboardAPIGetAccountsResponse: &gcasdashboardapi.GetAccountsResponse{
					"aws": {"1111", "2222"},
				},
			},
			want:    []*entities.GCASCSPCost{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut, mock, deferFunc := NewSUT(t)
			defer deferFunc()

			tt.prepareMockFn(mock)

			got, err := sut.ToGCASCSPCost(tt.args.eventID, tt.args.gcasDashboardAPIGetAccountsResponse)

			if (err != nil) != tt.wantErr {
				t.Fatalf("error in Test_Usecase_ToGCASCSPCost: name = %s, err = %+v", tt.name, err)
			}

			if len(got) != len(tt.want) {
				t.Fatalf("error in Test_Usecase_ToGCASCSPCost: name = %s, len(got) = %d, len(want) = %d", tt.name, len(got), len(tt.want))
			}

			sort.Slice(got, func(i, j int) bool {
				return got[i].CSP() < got[j].CSP()
			})

			for i := range got {
				if got[i].EventID() != tt.want[i].EventID() || got[i].CSP() != tt.want[i].CSP() || got[i].TotalCost() != tt.want[i].TotalCost() {
					t.Fatalf("error in Test_Usecase_ToGCASCSPCost: name = %s, got[%d] = %v, want[%d] = %v", tt.name, i, got[i], i, tt.want[i])
				}
			}
		})
	}
}

func TestUsecaseToOutputFromGCASAccount(t *testing.T) {
	tests := []struct {
		name string
		args *gcasdashboardapi.GetAccountsResponse
		want *billable.Output
	}{
		{
			name: "正常系：0件の場合",
			args: &gcasdashboardapi.GetAccountsResponse{},
			want: &billable.Output{},
		},
		{
			name: "正常系：CSP1件、各アカウント1件の場合",
			args: &gcasdashboardapi.GetAccountsResponse{
				"aws": {"1111"},
			},
			want: &billable.Output{
				GCASAccounts: []*billable.OutputAccount{
					{CSP: "aws", AccountID: "1111"},
				},
			},
		},
		{
			name: "正常系：CSP1件、各アカウント2件の場合",
			args: &gcasdashboardapi.GetAccountsResponse{
				"aws": {"1111", "2222"},
			},
			want: &billable.Output{
				GCASAccounts: []*billable.OutputAccount{
					{CSP: "aws", AccountID: "1111"},
					{CSP: "aws", AccountID: "2222"},
				},
			},
		},
		{
			name: "正常系：CSP2件、各アカウント1件の場合",
			args: &gcasdashboardapi.GetAccountsResponse{
				"aws": {"1111"},
				"gcp": {"3333"},
			},
			want: &billable.Output{
				GCASAccounts: []*billable.OutputAccount{
					{CSP: "aws", AccountID: "1111"},
					{CSP: "gcp", AccountID: "3333"},
				},
			},
		},
		{
			name: "正常系：CSP2件、各アカウント2件の場合",
			args: &gcasdashboardapi.GetAccountsResponse{
				"aws": {"1111", "2222"},
				"gcp": {"3333", "4444"},
			},
			want: &billable.Output{
				GCASAccounts: []*billable.OutputAccount{
					{CSP: "aws", AccountID: "1111"},
					{CSP: "aws", AccountID: "2222"},
					{CSP: "gcp", AccountID: "3333"},
					{CSP: "gcp", AccountID: "4444"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut, _, deferFunc := NewSUT(t)
			defer deferFunc()

			got := sut.ToOutputFromGCASAccount(tt.args)

			if len(got.GCASAccounts) != len(tt.want.GCASAccounts) {
				t.Fatalf("error in Test_Usecase_ToOutputFromGCASAccount: name = %s, len(got) = %d, len(want) = %d", tt.name, len(got.GCASAccounts), len(tt.want.GCASAccounts))
			}

			sort.Slice(got.GCASAccounts, func(i, j int) bool {
				if got.GCASAccounts[i].CSP != got.GCASAccounts[j].CSP {
					return got.GCASAccounts[i].CSP < got.GCASAccounts[j].CSP
				}
				return got.GCASAccounts[i].AccountID < got.GCASAccounts[j].AccountID
			})
			for i := range got.GCASAccounts {
				if !reflect.DeepEqual(got.GCASAccounts[i], tt.want.GCASAccounts[i]) {
					t.Fatalf("error in Test_Usecase_ToOutputFromGCASAccount: name = %s, got[%d] = %v, want[%d] = %v", tt.name, i, got.GCASAccounts[i], i, tt.want.GCASAccounts[i])
				}
			}
		})
	}
}

type Mock struct {
	MockGCASDashboardAPI      *mockapi.MockGCASDashboardAPI
	MockGCASAPI               *mockapi.MockGCASAPI
	MockEventStatusService    *mockservices.MockEventStatusService
	MockGCASCSPCostRepository *mockrepositories.MockGCASCSPCostRepository
	deferFunc                 func()
}

func NewSUT(t *testing.T) (*billable.Usecase, *Mock, func()) {
	mockCtrlGCASDashboardAPI := gomock.NewController(t)
	mockCtrlGCASAPI := gomock.NewController(t)
	mockCtrlEventStatusService := gomock.NewController(t)
	mockCtrlGCASCSPCostRepository := gomock.NewController(t)
	deferFunc := func() {
		mockCtrlGCASDashboardAPI.Finish()
		mockCtrlGCASAPI.Finish()
		mockCtrlEventStatusService.Finish()
		mockCtrlGCASCSPCostRepository.Finish()
	}

	mockGCASDashboardAPI := mockapi.NewMockGCASDashboardAPI(mockCtrlGCASDashboardAPI)
	mockGCASAPI := mockapi.NewMockGCASAPI(mockCtrlGCASAPI)
	mockEventStatusService := mockservices.NewMockEventStatusService(mockCtrlEventStatusService)
	mockGCASCSPCostRepository := mockrepositories.NewMockGCASCSPCostRepository(mockCtrlGCASCSPCostRepository)

	sut := billable.NewUsecase(billable.Dependencies{
		GCASDashboardAPI:      mockGCASDashboardAPI,
		GCASAPI:               mockGCASAPI,
		EventStatusService:    mockEventStatusService,
		GCASCSPCostRepository: mockGCASCSPCostRepository,
	})

	return sut, &Mock{
		MockGCASDashboardAPI:      mockGCASDashboardAPI,
		MockGCASAPI:               mockGCASAPI,
		MockEventStatusService:    mockEventStatusService,
		MockGCASCSPCostRepository: mockGCASCSPCostRepository,
	}, deferFunc
}
