// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/opencars/operations/pkg/domain (interfaces: OperationRepository,ResourceRepository)

// Package mockstore is a generated GoMock package.
package mockstore

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/opencars/operations/pkg/domain"
)

// MockOperationRepository is a mock of OperationRepository interface.
type MockOperationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOperationRepositoryMockRecorder
}

// MockOperationRepositoryMockRecorder is the mock recorder for MockOperationRepository.
type MockOperationRepositoryMockRecorder struct {
	mock *MockOperationRepository
}

// NewMockOperationRepository creates a new mock instance.
func NewMockOperationRepository(ctrl *gomock.Controller) *MockOperationRepository {
	mock := &MockOperationRepository{ctrl: ctrl}
	mock.recorder = &MockOperationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationRepository) EXPECT() *MockOperationRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOperationRepository) Create(arg0 context.Context, arg1 ...*domain.Operation) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOperationRepositoryMockRecorder) Create(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOperationRepository)(nil).Create), varargs...)
}

// DeleteByResourceID mocks base method.
func (m *MockOperationRepository) DeleteByResourceID(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByResourceID", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByResourceID indicates an expected call of DeleteByResourceID.
func (mr *MockOperationRepositoryMockRecorder) DeleteByResourceID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByResourceID", reflect.TypeOf((*MockOperationRepository)(nil).DeleteByResourceID), arg0, arg1)
}

// FindByNumber mocks base method.
func (m *MockOperationRepository) FindByNumber(arg0 context.Context, arg1 string, arg2 uint64, arg3 string) ([]domain.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByNumber", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]domain.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByNumber indicates an expected call of FindByNumber.
func (mr *MockOperationRepositoryMockRecorder) FindByNumber(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByNumber", reflect.TypeOf((*MockOperationRepository)(nil).FindByNumber), arg0, arg1, arg2, arg3)
}

// MockResourceRepository is a mock of ResourceRepository interface.
type MockResourceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResourceRepositoryMockRecorder
}

// MockResourceRepositoryMockRecorder is the mock recorder for MockResourceRepository.
type MockResourceRepositoryMockRecorder struct {
	mock *MockResourceRepository
}

// NewMockResourceRepository creates a new mock instance.
func NewMockResourceRepository(ctrl *gomock.Controller) *MockResourceRepository {
	mock := &MockResourceRepository{ctrl: ctrl}
	mock.recorder = &MockResourceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceRepository) EXPECT() *MockResourceRepositoryMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockResourceRepository) All(arg0 context.Context) ([]domain.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].([]domain.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockResourceRepositoryMockRecorder) All(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockResourceRepository)(nil).All), arg0)
}

// Create mocks base method.
func (m *MockResourceRepository) Create(arg0 context.Context, arg1 *domain.Resource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockResourceRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResourceRepository)(nil).Create), arg0, arg1)
}

// FindByUID mocks base method.
func (m *MockResourceRepository) FindByUID(arg0 context.Context, arg1 string) (*domain.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUID", arg0, arg1)
	ret0, _ := ret[0].(*domain.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUID indicates an expected call of FindByUID.
func (mr *MockResourceRepositoryMockRecorder) FindByUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUID", reflect.TypeOf((*MockResourceRepository)(nil).FindByUID), arg0, arg1)
}

// Update mocks base method.
func (m *MockResourceRepository) Update(arg0 context.Context, arg1 *domain.Resource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockResourceRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResourceRepository)(nil).Update), arg0, arg1)
}
