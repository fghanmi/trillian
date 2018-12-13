// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/trillian/storage/cache (interfaces: NodeStorage)

// Package cache is a generated GoMock package.
package cache

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/google/trillian/storage"
	storagepb "github.com/google/trillian/storage/storagepb"
	reflect "reflect"
)

// MockNodeStorage is a mock of NodeStorage interface
type MockNodeStorage struct {
	ctrl     *gomock.Controller
	recorder *MockNodeStorageMockRecorder
}

// MockNodeStorageMockRecorder is the mock recorder for MockNodeStorage
type MockNodeStorageMockRecorder struct {
	mock *MockNodeStorage
}

// NewMockNodeStorage creates a new mock instance
func NewMockNodeStorage(ctrl *gomock.Controller) *MockNodeStorage {
	mock := &MockNodeStorage{ctrl: ctrl}
	mock.recorder = &MockNodeStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeStorage) EXPECT() *MockNodeStorageMockRecorder {
	return m.recorder
}

// GetSubtree mocks base method
func (m *MockNodeStorage) GetSubtree(arg0 storage.NodeID) (*storagepb.SubtreeProto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubtree", arg0)
	ret0, _ := ret[0].(*storagepb.SubtreeProto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubtree indicates an expected call of GetSubtree
func (mr *MockNodeStorageMockRecorder) GetSubtree(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubtree", reflect.TypeOf((*MockNodeStorage)(nil).GetSubtree), arg0)
}

// SetSubtrees mocks base method
func (m *MockNodeStorage) SetSubtrees(arg0 []*storagepb.SubtreeProto) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSubtrees", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSubtrees indicates an expected call of SetSubtrees
func (mr *MockNodeStorageMockRecorder) SetSubtrees(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubtrees", reflect.TypeOf((*MockNodeStorage)(nil).SetSubtrees), arg0)
}
