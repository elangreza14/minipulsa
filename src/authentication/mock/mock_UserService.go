// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/elangreza14/minipulsa/authentication/core/user (interfaces: UserService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/elangreza14/minipulsa/authentication/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// LoginRegister mocks base method.
func (m *MockUserService) LoginRegister(arg0 context.Context, arg1 entity.ReqPostPutUser) (string, *int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginRegister", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoginRegister indicates an expected call of LoginRegister.
func (mr *MockUserServiceMockRecorder) LoginRegister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginRegister", reflect.TypeOf((*MockUserService)(nil).LoginRegister), arg0, arg1)
}

// ValidateToken mocks base method.
func (m *MockUserService) ValidateToken(arg0 context.Context, arg1 string) (*int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", arg0, arg1)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockUserServiceMockRecorder) ValidateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockUserService)(nil).ValidateToken), arg0, arg1)
}
