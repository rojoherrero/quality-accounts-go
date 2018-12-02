// Code generated by MockGen. DO NOT EDIT.
// Source: roleRepository.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/rojoherrero/quality-accounts/app/model"
	reflect "reflect"
)

// MockRoleRepository is a mock of RoleRepository interface
type MockRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoleRepositoryMockRecorder
}

// MockRoleRepositoryMockRecorder is the mock recorder for MockRoleRepository
type MockRoleRepositoryMockRecorder struct {
	mock *MockRoleRepository
}

// NewMockRoleRepository creates a new mock instance
func NewMockRoleRepository(ctrl *gomock.Controller) *MockRoleRepository {
	mock := &MockRoleRepository{ctrl: ctrl}
	mock.recorder = &MockRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoleRepository) EXPECT() *MockRoleRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockRoleRepository) Save(role model.RoleDepartment) error {
	ret := m.ctrl.Call(m, "Save", role)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockRoleRepositoryMockRecorder) Save(role interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockRoleRepository)(nil).Save), role)
}

// Update mocks base method
func (m *MockRoleRepository) Update(data model.RoleDepartmentUpdate) error {
	ret := m.ctrl.Call(m, "Update", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRoleRepositoryMockRecorder) Update(data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleRepository)(nil).Update), data)
}

// Paginate mocks base method
func (m *MockRoleRepository) Paginate(start, end int) (model.RolesDepartments, error) {
	ret := m.ctrl.Call(m, "Paginate", start, end)
	ret0, _ := ret[0].(model.RolesDepartments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Paginate indicates an expected call of Paginate
func (mr *MockRoleRepositoryMockRecorder) Paginate(start, end interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockRoleRepository)(nil).Paginate), start, end)
}

// Delete mocks base method
func (m *MockRoleRepository) Delete(code string) error {
	ret := m.ctrl.Call(m, "Delete", code)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRoleRepositoryMockRecorder) Delete(code interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleRepository)(nil).Delete), code)
}
