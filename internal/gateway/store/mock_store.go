// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/superproj/onex/internal/gateway/store (interfaces: IStore,ChainStore,MinerStore,MinerSetStore)

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/superproj/onex/internal/gateway/model"
	meta "github.com/superproj/onex/internal/pkg/meta"
)

// MockIStore is a mock of IStore interface.
type MockIStore struct {
	ctrl     *gomock.Controller
	recorder *MockIStoreMockRecorder
}

// MockIStoreMockRecorder is the mock recorder for MockIStore.
type MockIStoreMockRecorder struct {
	mock *MockIStore
}

// NewMockIStore creates a new mock instance.
func NewMockIStore(ctrl *gomock.Controller) *MockIStore {
	mock := &MockIStore{ctrl: ctrl}
	mock.recorder = &MockIStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStore) EXPECT() *MockIStoreMockRecorder {
	return m.recorder
}

// Chains mocks base method.
func (m *MockIStore) Chains() ChainStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chains")
	ret0, _ := ret[0].(ChainStore)
	return ret0
}

// Chains indicates an expected call of Chains.
func (mr *MockIStoreMockRecorder) Chains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chains", reflect.TypeOf((*MockIStore)(nil).Chains))
}

// MinerSets mocks base method.
func (m *MockIStore) MinerSets() MinerSetStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinerSets")
	ret0, _ := ret[0].(MinerSetStore)
	return ret0
}

// MinerSets indicates an expected call of MinerSets.
func (mr *MockIStoreMockRecorder) MinerSets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinerSets", reflect.TypeOf((*MockIStore)(nil).MinerSets))
}

// Miners mocks base method.
func (m *MockIStore) Miners() MinerStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Miners")
	ret0, _ := ret[0].(MinerStore)
	return ret0
}

// Miners indicates an expected call of Miners.
func (mr *MockIStoreMockRecorder) Miners() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Miners", reflect.TypeOf((*MockIStore)(nil).Miners))
}

// TX mocks base method.
func (m *MockIStore) TX(arg0 context.Context, arg1 func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TX", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TX indicates an expected call of TX.
func (mr *MockIStoreMockRecorder) TX(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TX", reflect.TypeOf((*MockIStore)(nil).TX), arg0, arg1)
}

// MockChainStore is a mock of ChainStore interface.
type MockChainStore struct {
	ctrl     *gomock.Controller
	recorder *MockChainStoreMockRecorder
}

// MockChainStoreMockRecorder is the mock recorder for MockChainStore.
type MockChainStoreMockRecorder struct {
	mock *MockChainStore
}

// NewMockChainStore creates a new mock instance.
func NewMockChainStore(ctrl *gomock.Controller) *MockChainStore {
	mock := &MockChainStore{ctrl: ctrl}
	mock.recorder = &MockChainStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainStore) EXPECT() *MockChainStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChainStore) Create(arg0 context.Context, arg1 *model.ChainM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockChainStoreMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChainStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockChainStore) Delete(arg0 context.Context, arg1 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockChainStoreMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockChainStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockChainStore) Get(arg0 context.Context, arg1 map[string]any) (*model.ChainM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.ChainM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockChainStoreMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockChainStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockChainStore) List(arg0 context.Context, arg1 string, arg2 ...meta.ListOption) (int64, []*model.ChainM, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].([]*model.ChainM)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockChainStoreMockRecorder) List(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockChainStore)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockChainStore) Update(arg0 context.Context, arg1 *model.ChainM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockChainStoreMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockChainStore)(nil).Update), arg0, arg1)
}

// MockMinerStore is a mock of MinerStore interface.
type MockMinerStore struct {
	ctrl     *gomock.Controller
	recorder *MockMinerStoreMockRecorder
}

// MockMinerStoreMockRecorder is the mock recorder for MockMinerStore.
type MockMinerStoreMockRecorder struct {
	mock *MockMinerStore
}

// NewMockMinerStore creates a new mock instance.
func NewMockMinerStore(ctrl *gomock.Controller) *MockMinerStore {
	mock := &MockMinerStore{ctrl: ctrl}
	mock.recorder = &MockMinerStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMinerStore) EXPECT() *MockMinerStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMinerStore) Create(arg0 context.Context, arg1 *model.MinerM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMinerStoreMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMinerStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockMinerStore) Delete(arg0 context.Context, arg1 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMinerStoreMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMinerStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockMinerStore) Get(arg0 context.Context, arg1 map[string]any) (*model.MinerM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.MinerM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMinerStoreMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMinerStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockMinerStore) List(arg0 context.Context, arg1 string, arg2 ...meta.ListOption) (int64, []*model.MinerM, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].([]*model.MinerM)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockMinerStoreMockRecorder) List(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMinerStore)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockMinerStore) Update(arg0 context.Context, arg1 *model.MinerM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMinerStoreMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMinerStore)(nil).Update), arg0, arg1)
}

// MockMinerSetStore is a mock of MinerSetStore interface.
type MockMinerSetStore struct {
	ctrl     *gomock.Controller
	recorder *MockMinerSetStoreMockRecorder
}

// MockMinerSetStoreMockRecorder is the mock recorder for MockMinerSetStore.
type MockMinerSetStoreMockRecorder struct {
	mock *MockMinerSetStore
}

// NewMockMinerSetStore creates a new mock instance.
func NewMockMinerSetStore(ctrl *gomock.Controller) *MockMinerSetStore {
	mock := &MockMinerSetStore{ctrl: ctrl}
	mock.recorder = &MockMinerSetStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMinerSetStore) EXPECT() *MockMinerSetStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMinerSetStore) Create(arg0 context.Context, arg1 *model.MinerSetM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMinerSetStoreMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMinerSetStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockMinerSetStore) Delete(arg0 context.Context, arg1 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMinerSetStoreMockRecorder) Delete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMinerSetStore)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockMinerSetStore) Get(arg0 context.Context, arg1 map[string]any) (*model.MinerSetM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*model.MinerSetM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMinerSetStoreMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMinerSetStore)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockMinerSetStore) List(arg0 context.Context, arg1 string, arg2 ...meta.ListOption) (int64, []*model.MinerSetM, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].([]*model.MinerSetM)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockMinerSetStoreMockRecorder) List(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMinerSetStore)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockMinerSetStore) Update(arg0 context.Context, arg1 *model.MinerSetM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMinerSetStoreMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMinerSetStore)(nil).Update), arg0, arg1)
}