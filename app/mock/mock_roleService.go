// Code generated by MockGen. DO NOT EDIT.
// Source: roleService.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/rojoherrero/quality-accounts/app/model"
	reflect "reflect"
)

// MockRoleService is a mock of RoleService interface
type MockRoleService struct {
	ctrl     *gomock.Controller
	recorder *MockRoleServiceMockRecorder
}

// MockRoleServiceMockRecorder is the mock recorder for MockRoleService
type MockRoleServiceMockRecorder struct {
	mock *MockRoleService
}

// NewMockRoleService creates a new mock instance
func NewMockRoleService(ctrl *gomock.Controller) *MockRoleService {
	mock := &MockRoleService{ctrl: ctrl}
	mock.recorder = &MockRoleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleService) EXPECT() *MockRoleServiceMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockRoleService) Save(ctx context.Context, role []model.Role) error {
	ret := m.ctrl.Call(m, "Save", ctx, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockRoleServiceMockRecorder) Save(ctx, role interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRoleService)(nil).Save), ctx, role)
}

// Update mocks base method
func (m *MockRoleService) Update(ctx context.Context, data model.Role, oldCode string) error {
	ret := m.ctrl.Call(m, "Update", ctx, data, oldCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRoleServiceMockRecorder) Update(ctx, data, oldCode interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleService)(nil).Update), ctx, data, oldCode)
}

// Paginate mocks base method
func (m *MockRoleService) Paginate(ctx context.Context, start, end int) ([]model.Role, error) {
	ret := m.ctrl.Call(m, "Paginate", ctx, start, end)
	ret0, _ := ret[0].([]model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Paginate indicates an expected call of Paginate
func (mr *MockRoleServiceMockRecorder) Paginate(ctx, start, end interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockRoleService)(nil).Paginate), ctx, start, end)
}

// Delete mocks base method
func (m *MockRoleService) Delete(ctx context.Context, id int64) error {
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRoleServiceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleService)(nil).Delete), ctx, id)
}
