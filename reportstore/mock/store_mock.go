// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	reportstore "github.com/diegojunges/reporting/reportstore"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateReport mocks base method.
func (m *MockStore) CreateReport(r reportstore.CreateReportRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReport", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReport indicates an expected call of CreateReport.
func (mr *MockStoreMockRecorder) CreateReport(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReport", reflect.TypeOf((*MockStore)(nil).CreateReport), r)
}