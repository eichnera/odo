// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/alizer/interface.go

// Package alizer is a generated GoMock package.
package alizer

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/redhat-developer/alizer/go/pkg/apis/model"
	api "github.com/redhat-developer/odo/pkg/api"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DetectFramework mocks base method.
func (m *MockClient) DetectFramework(ctx context.Context, path string) (model.DevFileType, string, api.Registry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectFramework", ctx, path)
	ret0, _ := ret[0].(model.DevFileType)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(api.Registry)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// DetectFramework indicates an expected call of DetectFramework.
func (mr *MockClientMockRecorder) DetectFramework(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectFramework", reflect.TypeOf((*MockClient)(nil).DetectFramework), ctx, path)
}

// DetectName mocks base method.
func (m *MockClient) DetectName(path string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectName", path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectName indicates an expected call of DetectName.
func (mr *MockClientMockRecorder) DetectName(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectName", reflect.TypeOf((*MockClient)(nil).DetectName), path)
}

// DetectPorts mocks base method.
func (m *MockClient) DetectPorts(path string) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectPorts", path)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectPorts indicates an expected call of DetectPorts.
func (mr *MockClientMockRecorder) DetectPorts(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectPorts", reflect.TypeOf((*MockClient)(nil).DetectPorts), path)
}
