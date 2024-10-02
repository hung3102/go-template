// Code generated by MockGen. DO NOT EDIT.
// Source: org_csp_account_cost_repository.go
//
// Generated by this command:
//
//	mockgen -source=org_csp_account_cost_repository.go -destination=../repositoryimpl/mocks/org_csp_account_cost_repository_mock.go -package=mockrepositories
//

// Package mockrepositories is a generated GoMock package.
package mockrepositories

import (
	context "context"
	reflect "reflect"

	entities "github.com/topgate/gcim-temporary/back/app/internal/entities"
	repositories "github.com/topgate/gcim-temporary/back/app/internal/repositories"
	gomock "go.uber.org/mock/gomock"
)

// MockORGCSPAccountCostRepository is a mock of ORGCSPAccountCostRepository interface.
type MockORGCSPAccountCostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockORGCSPAccountCostRepositoryMockRecorder
}

// MockORGCSPAccountCostRepositoryMockRecorder is the mock recorder for MockORGCSPAccountCostRepository.
type MockORGCSPAccountCostRepositoryMockRecorder struct {
	mock *MockORGCSPAccountCostRepository
}

// NewMockORGCSPAccountCostRepository creates a new mock instance.
func NewMockORGCSPAccountCostRepository(ctrl *gomock.Controller) *MockORGCSPAccountCostRepository {
	mock := &MockORGCSPAccountCostRepository{ctrl: ctrl}
	mock.recorder = &MockORGCSPAccountCostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockORGCSPAccountCostRepository) EXPECT() *MockORGCSPAccountCostRepositoryMockRecorder {
	return m.recorder
}

// SearchByParam mocks base method.
func (m *MockORGCSPAccountCostRepository) SearchByParam(ctx context.Context, param *repositories.OrgCSPAccountCostSearchParam) ([]*entities.OrgCSPAccountCost, *repositories.OrgCSPAccountCostPagingResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByParam", ctx, param)
	ret0, _ := ret[0].([]*entities.OrgCSPAccountCost)
	ret1, _ := ret[1].(*repositories.OrgCSPAccountCostPagingResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchByParam indicates an expected call of SearchByParam.
func (mr *MockORGCSPAccountCostRepositoryMockRecorder) SearchByParam(ctx, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByParam", reflect.TypeOf((*MockORGCSPAccountCostRepository)(nil).SearchByParam), ctx, param)
}
