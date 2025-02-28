// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/shop.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/UserNameShouldBeHere/AvitoTask/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockShopStorage is a mock of ShopStorage interface.
type MockShopStorage struct {
	ctrl     *gomock.Controller
	recorder *MockShopStorageMockRecorder
}

// MockShopStorageMockRecorder is the mock recorder for MockShopStorage.
type MockShopStorageMockRecorder struct {
	mock *MockShopStorage
}

// NewMockShopStorage creates a new mock instance.
func NewMockShopStorage(ctrl *gomock.Controller) *MockShopStorage {
	mock := &MockShopStorage{ctrl: ctrl}
	mock.recorder = &MockShopStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShopStorage) EXPECT() *MockShopStorageMockRecorder {
	return m.recorder
}

// BuyItem mocks base method.
func (m *MockShopStorage) BuyItem(ctx context.Context, username, itemName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyItem", ctx, username, itemName)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyItem indicates an expected call of BuyItem.
func (mr *MockShopStorageMockRecorder) BuyItem(ctx, username, itemName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyItem", reflect.TypeOf((*MockShopStorage)(nil).BuyItem), ctx, username, itemName)
}

// GetInfo mocks base method.
func (m *MockShopStorage) GetInfo(ctx context.Context, username string) (domain.InventoryInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", ctx, username)
	ret0, _ := ret[0].(domain.InventoryInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockShopStorageMockRecorder) GetInfo(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockShopStorage)(nil).GetInfo), ctx, username)
}

// SendCoin mocks base method.
func (m *MockShopStorage) SendCoin(ctx context.Context, transaction domain.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoin", ctx, transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoin indicates an expected call of SendCoin.
func (mr *MockShopStorageMockRecorder) SendCoin(ctx, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoin", reflect.TypeOf((*MockShopStorage)(nil).SendCoin), ctx, transaction)
}
