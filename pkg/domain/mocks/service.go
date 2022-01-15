// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/opencars/operations/pkg/domain (interfaces: CustomerService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/opencars/operations/pkg/domain/model"
	query "github.com/opencars/operations/pkg/domain/query"
)

// MockCustomerService is a mock of CustomerService interface.
type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

// MockCustomerServiceMockRecorder is the mock recorder for MockCustomerService.
type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}

// NewMockCustomerService creates a new mock instance.
func NewMockCustomerService(ctrl *gomock.Controller) *MockCustomerService {
	mock := &MockCustomerService{ctrl: ctrl}
	mock.recorder = &MockCustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerService) EXPECT() *MockCustomerServiceMockRecorder {
	return m.recorder
}

// FindByNumber mocks base method.
func (m *MockCustomerService) FindByNumber(arg0 context.Context, arg1 *query.ListByNumber) ([]model.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByNumber", arg0, arg1)
	ret0, _ := ret[0].([]model.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByNumber indicates an expected call of FindByNumber.
func (mr *MockCustomerServiceMockRecorder) FindByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByNumber", reflect.TypeOf((*MockCustomerService)(nil).FindByNumber), arg0, arg1)
}

// FindByVIN mocks base method.
func (m *MockCustomerService) FindByVIN(arg0 context.Context, arg1 *query.ListByVIN) ([]model.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByVIN", arg0, arg1)
	ret0, _ := ret[0].([]model.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByVIN indicates an expected call of FindByVIN.
func (mr *MockCustomerServiceMockRecorder) FindByVIN(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByVIN", reflect.TypeOf((*MockCustomerService)(nil).FindByVIN), arg0, arg1)
}
