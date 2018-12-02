// Code generated by MockGen. DO NOT EDIT.
// Source: departmentService.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/rojoherrero/quality-accounts/app/model"
	reflect "reflect"
)

// MockDepartmentService is a mock of DepartmentService interface
type MockDepartmentService struct {
	ctrl     *gomock.Controller
	recorder *MockDepartmentServiceMockRecorder
}

// MockDepartmentServiceMockRecorder is the mock recorder for MockDepartmentService
type MockDepartmentServiceMockRecorder struct {
	mock *MockDepartmentService
}

// NewMockDepartmentService creates a new mock instance
func NewMockDepartmentService(ctrl *gomock.Controller) *MockDepartmentService {
	mock := &MockDepartmentService{ctrl: ctrl}
	mock.recorder = &MockDepartmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDepartmentService) EXPECT() *MockDepartmentServiceMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockDepartmentService) Save(dept model.RoleDepartment) error {
	ret := m.ctrl.Call(m, "Save", dept)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockDepartmentServiceMockRecorder) Save(dept interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDepartmentService)(nil).Save), dept)
}

// Update mocks base method
func (m *MockDepartmentService) Update(data model.RoleDepartmentUpdate) error {
	ret := m.ctrl.Call(m, "Update", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDepartmentServiceMockRecorder) Update(data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDepartmentService)(nil).Update), data)
}

// Paginate mocks base method
func (m *MockDepartmentService) Paginate(start, end int) (model.RolesDepartments, error) {
	ret := m.ctrl.Call(m, "Paginate", start, end)
	ret0, _ := ret[0].(model.RolesDepartments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Paginate indicates an expected call of Paginate
func (mr *MockDepartmentServiceMockRecorder) Paginate(start, end interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockDepartmentService)(nil).Paginate), start, end)
}

// Delete mocks base method
func (m *MockDepartmentService) Delete(code string) error {
	ret := m.ctrl.Call(m, "Delete", code)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDepartmentServiceMockRecorder) Delete(code interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDepartmentService)(nil).Delete), code)
}