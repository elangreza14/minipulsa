// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/elangreza14/minipulsa/api-gateway/core/product (interfaces: ProductService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/elangreza14/minipulsa/api-gateway/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// GetProducts mocks base method.
func (m *MockProductService) GetProducts(arg0 context.Context) (*entity.HTTPResGetProducts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", arg0)
	ret0, _ := ret[0].(*entity.HTTPResGetProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductServiceMockRecorder) GetProducts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProductService)(nil).GetProducts), arg0)
}
