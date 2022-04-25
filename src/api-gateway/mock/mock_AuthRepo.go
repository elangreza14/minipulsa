// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/elangreza14/minipulsa/api-gateway/port (interfaces: AuthRepo)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	minipulsa "github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAuthRepo is a mock of AuthRepo interface.
type MockAuthRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepoMockRecorder
}

// MockAuthRepoMockRecorder is the mock recorder for MockAuthRepo.
type MockAuthRepoMockRecorder struct {
	mock *MockAuthRepo
}

// NewMockAuthRepo creates a new mock instance.
func NewMockAuthRepo(ctrl *gomock.Controller) *MockAuthRepo {
	mock := &MockAuthRepo{ctrl: ctrl}
	mock.recorder = &MockAuthRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepo) EXPECT() *MockAuthRepoMockRecorder {
	return m.recorder
}

// LoginRegister mocks base method.
func (m *MockAuthRepo) LoginRegister(arg0 context.Context, arg1 *minipulsa.LoginRegisterRequest, arg2 ...grpc.CallOption) (*minipulsa.LoginRegisterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoginRegister", varargs...)
	ret0, _ := ret[0].(*minipulsa.LoginRegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginRegister indicates an expected call of LoginRegister.
func (mr *MockAuthRepoMockRecorder) LoginRegister(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginRegister", reflect.TypeOf((*MockAuthRepo)(nil).LoginRegister), varargs...)
}

// ValidateToken mocks base method.
func (m *MockAuthRepo) ValidateToken(arg0 context.Context, arg1 *minipulsa.ValidateTokenRequest, arg2 ...grpc.CallOption) (*minipulsa.ValidateTokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateToken", varargs...)
	ret0, _ := ret[0].(*minipulsa.ValidateTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAuthRepoMockRecorder) ValidateToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAuthRepo)(nil).ValidateToken), varargs...)
}
