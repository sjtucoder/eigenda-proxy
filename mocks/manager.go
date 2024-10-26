// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigenda-proxy/store (interfaces: IManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	commitments "github.com/Layr-Labs/eigenda-proxy/commitments"
	gomock "github.com/golang/mock/gomock"
)

// MockIManager is a mock of IManager interface.
type MockIManager struct {
	ctrl     *gomock.Controller
	recorder *MockIManagerMockRecorder
}

// MockIManagerMockRecorder is the mock recorder for MockIManager.
type MockIManagerMockRecorder struct {
	mock *MockIManager
}

// NewMockIManager creates a new mock instance.
func NewMockIManager(ctrl *gomock.Controller) *MockIManager {
	mock := &MockIManager{ctrl: ctrl}
	mock.recorder = &MockIManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIManager) EXPECT() *MockIManagerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIManager) Get(arg0 context.Context, arg1 []byte, arg2 commitments.CommitmentMode) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIManagerMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIManager)(nil).Get), arg0, arg1, arg2)
}

// Put mocks base method.
func (m *MockIManager) Put(arg0 context.Context, arg1 commitments.CommitmentMode, arg2, arg3 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockIManagerMockRecorder) Put(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockIManager)(nil).Put), arg0, arg1, arg2, arg3)
}
