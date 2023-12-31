// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/ledger/test_service.go

// Package mock_ledger is a generated GoMock package.
package mock_ledger

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTestService is a mock of TestService interface.
type MockTestService struct {
	ctrl     *gomock.Controller
	recorder *MockTestServiceMockRecorder
}

// MockTestServiceMockRecorder is the mock recorder for MockTestService.
type MockTestServiceMockRecorder struct {
	mock *MockTestService
}

// NewMockTestService creates a new mock instance.
func NewMockTestService(ctrl *gomock.Controller) *MockTestService {
	mock := &MockTestService{ctrl: ctrl}
	mock.recorder = &MockTestServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestService) EXPECT() *MockTestServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTestService) Create(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTestServiceMockRecorder) Create(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTestService)(nil).Create), id)
}
