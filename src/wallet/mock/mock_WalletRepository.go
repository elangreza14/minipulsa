// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/elangreza14/minipulsa/wallet/core/wallet (interfaces: WalletRepository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/elangreza14/minipulsa/wallet/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockWalletRepository is a mock of WalletRepository interface.
type MockWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWalletRepositoryMockRecorder
}

// MockWalletRepositoryMockRecorder is the mock recorder for MockWalletRepository.
type MockWalletRepositoryMockRecorder struct {
	mock *MockWalletRepository
}

// NewMockWalletRepository creates a new mock instance.
func NewMockWalletRepository(ctrl *gomock.Controller) *MockWalletRepository {
	mock := &MockWalletRepository{ctrl: ctrl}
	mock.recorder = &MockWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletRepository) EXPECT() *MockWalletRepositoryMockRecorder {
	return m.recorder
}

// GetWalletByUserID mocks base method.
func (m *MockWalletRepository) GetWalletByUserID(arg0 context.Context, arg1 int64) (*entity.DBWallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletByUserID", arg0, arg1)
	ret0, _ := ret[0].(*entity.DBWallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletByUserID indicates an expected call of GetWalletByUserID.
func (mr *MockWalletRepositoryMockRecorder) GetWalletByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletByUserID", reflect.TypeOf((*MockWalletRepository)(nil).GetWalletByUserID), arg0, arg1)
}

// GetWalletHistories mocks base method.
func (m *MockWalletRepository) GetWalletHistories(arg0 context.Context, arg1 int64) (*[]entity.DBWalletHistories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletHistories", arg0, arg1)
	ret0, _ := ret[0].(*[]entity.DBWalletHistories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletHistories indicates an expected call of GetWalletHistories.
func (mr *MockWalletRepositoryMockRecorder) GetWalletHistories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletHistories", reflect.TypeOf((*MockWalletRepository)(nil).GetWalletHistories), arg0, arg1)
}

// GetWalletHistoryByReqUseWallet mocks base method.
func (m *MockWalletRepository) GetWalletHistoryByReqUseWallet(arg0 context.Context, arg1 entity.ReqUseWallet) (*entity.DBWalletHistoriesDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletHistoryByReqUseWallet", arg0, arg1)
	ret0, _ := ret[0].(*entity.DBWalletHistoriesDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletHistoryByReqUseWallet indicates an expected call of GetWalletHistoryByReqUseWallet.
func (mr *MockWalletRepositoryMockRecorder) GetWalletHistoryByReqUseWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletHistoryByReqUseWallet", reflect.TypeOf((*MockWalletRepository)(nil).GetWalletHistoryByReqUseWallet), arg0, arg1)
}

// InsertWallet mocks base method.
func (m *MockWalletRepository) InsertWallet(arg0 context.Context, arg1 entity.ReqUseWallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertWallet indicates an expected call of InsertWallet.
func (mr *MockWalletRepositoryMockRecorder) InsertWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWallet", reflect.TypeOf((*MockWalletRepository)(nil).InsertWallet), arg0, arg1)
}

// UpdateWalletByUserID mocks base method.
func (m *MockWalletRepository) UpdateWalletByUserID(arg0 context.Context, arg1 entity.ReqUseWallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWalletByUserID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWalletByUserID indicates an expected call of UpdateWalletByUserID.
func (mr *MockWalletRepositoryMockRecorder) UpdateWalletByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWalletByUserID", reflect.TypeOf((*MockWalletRepository)(nil).UpdateWalletByUserID), arg0, arg1)
}
