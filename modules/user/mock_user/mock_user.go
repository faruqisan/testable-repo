// Code generated by MockGen. DO NOT EDIT.
// Source: modules/user/user.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
        gomock "github.com/golang/mock/gomock"
        reflect "reflect"
)

// MockNSQ is a mock of NSQ interface
type MockNSQ struct {
        ctrl     *gomock.Controller
        recorder *MockNSQMockRecorder
}

// MockNSQMockRecorder is the mock recorder for MockNSQ
type MockNSQMockRecorder struct {
        mock *MockNSQ
}

// NewMockNSQ creates a new mock instance
func NewMockNSQ(ctrl *gomock.Controller) *MockNSQ {
        mock := &MockNSQ{ctrl: ctrl}
        mock.recorder = &MockNSQMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNSQ) EXPECT() *MockNSQMockRecorder {
        return m.recorder
}

// Publish mocks base method
func (m *MockNSQ) Publish(data string) error {
        ret := m.ctrl.Call(m, "Publish", data)
        ret0, _ := ret[0].(error)
        return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockNSQMockRecorder) Publish(data interface{}) *gomock.Call {
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockNSQ)(nil).Publish), data)
}

// MockRedis is a mock of Redis interface
type MockRedis struct {
        ctrl     *gomock.Controller
        recorder *MockRedisMockRecorder
}

// MockRedisMockRecorder is the mock recorder for MockRedis
type MockRedisMockRecorder struct {
        mock *MockRedis
}

// NewMockRedis creates a new mock instance
func NewMockRedis(ctrl *gomock.Controller) *MockRedis {
        mock := &MockRedis{ctrl: ctrl}
        mock.recorder = &MockRedisMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedis) EXPECT() *MockRedisMockRecorder {
        return m.recorder
}

// Set mocks base method
func (m *MockRedis) Set(key string, data interface{}) {
        m.ctrl.Call(m, "Set", key, data)
}

// Set indicates an expected call of Set
func (mr *MockRedisMockRecorder) Set(key, data interface{}) *gomock.Call {
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedis)(nil).Set), key, data)
}

// Get mocks base method
func (m *MockRedis) Get(key string) {
        m.ctrl.Call(m, "Get", key)
}

// Get indicates an expected call of Get
func (mr *MockRedisMockRecorder) Get(key interface{}) *gomock.Call {
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedis)(nil).Get), key)
}