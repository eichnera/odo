// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/init/asker/interface.go

// Package asker is a generated GoMock package.
package asker

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/redhat-developer/odo/pkg/api"
	registry "github.com/redhat-developer/odo/pkg/registry"
)

// MockAsker is a mock of Asker interface.
type MockAsker struct {
	ctrl     *gomock.Controller
	recorder *MockAskerMockRecorder
}

// MockAskerMockRecorder is the mock recorder for MockAsker.
type MockAskerMockRecorder struct {
	mock *MockAsker
}

// NewMockAsker creates a new mock instance.
func NewMockAsker(ctrl *gomock.Controller) *MockAsker {
	mock := &MockAsker{ctrl: ctrl}
	mock.recorder = &MockAskerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsker) EXPECT() *MockAskerMockRecorder {
	return m.recorder
}

// AskAddEnvVar mocks base method.
func (m *MockAsker) AskAddEnvVar() (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskAddEnvVar")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AskAddEnvVar indicates an expected call of AskAddEnvVar.
func (mr *MockAskerMockRecorder) AskAddEnvVar() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskAddEnvVar", reflect.TypeOf((*MockAsker)(nil).AskAddEnvVar))
}

// AskAddPort mocks base method.
func (m *MockAsker) AskAddPort() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskAddPort")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskAddPort indicates an expected call of AskAddPort.
func (mr *MockAskerMockRecorder) AskAddPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskAddPort", reflect.TypeOf((*MockAsker)(nil).AskAddPort))
}

// AskContainerName mocks base method.
func (m *MockAsker) AskContainerName(containers []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskContainerName", containers)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskContainerName indicates an expected call of AskContainerName.
func (mr *MockAskerMockRecorder) AskContainerName(containers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskContainerName", reflect.TypeOf((*MockAsker)(nil).AskContainerName), containers)
}

// AskCorrect mocks base method.
func (m *MockAsker) AskCorrect() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskCorrect")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskCorrect indicates an expected call of AskCorrect.
func (mr *MockAskerMockRecorder) AskCorrect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskCorrect", reflect.TypeOf((*MockAsker)(nil).AskCorrect))
}

// AskLanguage mocks base method.
func (m *MockAsker) AskLanguage(langs []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskLanguage", langs)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskLanguage indicates an expected call of AskLanguage.
func (mr *MockAskerMockRecorder) AskLanguage(langs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskLanguage", reflect.TypeOf((*MockAsker)(nil).AskLanguage), langs)
}

// AskName mocks base method.
func (m *MockAsker) AskName(defaultName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskName", defaultName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskName indicates an expected call of AskName.
func (mr *MockAskerMockRecorder) AskName(defaultName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskName", reflect.TypeOf((*MockAsker)(nil).AskName), defaultName)
}

// AskPersonalizeConfiguration mocks base method.
func (m *MockAsker) AskPersonalizeConfiguration(configuration ContainerConfiguration) (OperationOnContainer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskPersonalizeConfiguration", configuration)
	ret0, _ := ret[0].(OperationOnContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskPersonalizeConfiguration indicates an expected call of AskPersonalizeConfiguration.
func (mr *MockAskerMockRecorder) AskPersonalizeConfiguration(configuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskPersonalizeConfiguration", reflect.TypeOf((*MockAsker)(nil).AskPersonalizeConfiguration), configuration)
}

// AskStarterProject mocks base method.
func (m *MockAsker) AskStarterProject(projects []string) (bool, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskStarterProject", projects)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AskStarterProject indicates an expected call of AskStarterProject.
func (mr *MockAskerMockRecorder) AskStarterProject(projects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskStarterProject", reflect.TypeOf((*MockAsker)(nil).AskStarterProject), projects)
}

// AskType mocks base method.
func (m *MockAsker) AskType(types registry.TypesWithDetails) (bool, api.DevfileStack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskType", types)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(api.DevfileStack)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AskType indicates an expected call of AskType.
func (mr *MockAskerMockRecorder) AskType(types interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskType", reflect.TypeOf((*MockAsker)(nil).AskType), types)
}

// AskVersion mocks base method.
func (m *MockAsker) AskVersion(versions []api.DevfileStackVersion) (bool, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AskVersion", versions)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AskVersion indicates an expected call of AskVersion.
func (mr *MockAskerMockRecorder) AskVersion(versions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskVersion", reflect.TypeOf((*MockAsker)(nil).AskVersion), versions)
}
