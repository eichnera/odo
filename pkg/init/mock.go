// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/init/interface.go

// Package init is a generated GoMock package.
package init

import (
	reflect "reflect"

	v1alpha2 "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	parser "github.com/devfile/library/pkg/devfile/parser"
	gomock "github.com/golang/mock/gomock"
	backend "github.com/redhat-developer/odo/pkg/init/backend"
	filesystem "github.com/redhat-developer/odo/pkg/testingutil/filesystem"
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

// DownloadDevfile mocks base method.
func (m *MockClient) DownloadDevfile(devfileLocation *backend.DevfileLocation, destDir string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadDevfile", devfileLocation, destDir)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadDevfile indicates an expected call of DownloadDevfile.
func (mr *MockClientMockRecorder) DownloadDevfile(devfileLocation, destDir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadDevfile", reflect.TypeOf((*MockClient)(nil).DownloadDevfile), devfileLocation, destDir)
}

// DownloadStarterProject mocks base method.
func (m *MockClient) DownloadStarterProject(project *v1alpha2.StarterProject, dest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadStarterProject", project, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadStarterProject indicates an expected call of DownloadStarterProject.
func (mr *MockClientMockRecorder) DownloadStarterProject(project, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadStarterProject", reflect.TypeOf((*MockClient)(nil).DownloadStarterProject), project, dest)
}

// InitDevfile mocks base method.
func (m *MockClient) InitDevfile(flags map[string]string, contextDir string, preInitHandlerFunc func(bool), newDevfileHandlerFunc func(parser.DevfileObj) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitDevfile", flags, contextDir, preInitHandlerFunc, newDevfileHandlerFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitDevfile indicates an expected call of InitDevfile.
func (mr *MockClientMockRecorder) InitDevfile(flags, contextDir, preInitHandlerFunc, newDevfileHandlerFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitDevfile", reflect.TypeOf((*MockClient)(nil).InitDevfile), flags, contextDir, preInitHandlerFunc, newDevfileHandlerFunc)
}

// PersonalizeDevfileConfig mocks base method.
func (m *MockClient) PersonalizeDevfileConfig(devfileobj parser.DevfileObj, flags map[string]string, fs filesystem.Filesystem, dir string) (parser.DevfileObj, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersonalizeDevfileConfig", devfileobj, flags, fs, dir)
	ret0, _ := ret[0].(parser.DevfileObj)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PersonalizeDevfileConfig indicates an expected call of PersonalizeDevfileConfig.
func (mr *MockClientMockRecorder) PersonalizeDevfileConfig(devfileobj, flags, fs, dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersonalizeDevfileConfig", reflect.TypeOf((*MockClient)(nil).PersonalizeDevfileConfig), devfileobj, flags, fs, dir)
}

// PersonalizeName mocks base method.
func (m *MockClient) PersonalizeName(devfile parser.DevfileObj, flags map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersonalizeName", devfile, flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PersonalizeName indicates an expected call of PersonalizeName.
func (mr *MockClientMockRecorder) PersonalizeName(devfile, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersonalizeName", reflect.TypeOf((*MockClient)(nil).PersonalizeName), devfile, flags)
}

// SelectAndPersonalizeDevfile mocks base method.
func (m *MockClient) SelectAndPersonalizeDevfile(flags map[string]string, contextDir string) (parser.DevfileObj, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAndPersonalizeDevfile", flags, contextDir)
	ret0, _ := ret[0].(parser.DevfileObj)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SelectAndPersonalizeDevfile indicates an expected call of SelectAndPersonalizeDevfile.
func (mr *MockClientMockRecorder) SelectAndPersonalizeDevfile(flags, contextDir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAndPersonalizeDevfile", reflect.TypeOf((*MockClient)(nil).SelectAndPersonalizeDevfile), flags, contextDir)
}

// SelectDevfile mocks base method.
func (m *MockClient) SelectDevfile(flags map[string]string, fs filesystem.Filesystem, dir string) (*backend.DevfileLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectDevfile", flags, fs, dir)
	ret0, _ := ret[0].(*backend.DevfileLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectDevfile indicates an expected call of SelectDevfile.
func (mr *MockClientMockRecorder) SelectDevfile(flags, fs, dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectDevfile", reflect.TypeOf((*MockClient)(nil).SelectDevfile), flags, fs, dir)
}

// SelectStarterProject mocks base method.
func (m *MockClient) SelectStarterProject(devfile parser.DevfileObj, flags map[string]string, fs filesystem.Filesystem, dir string) (*v1alpha2.StarterProject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectStarterProject", devfile, flags, fs, dir)
	ret0, _ := ret[0].(*v1alpha2.StarterProject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectStarterProject indicates an expected call of SelectStarterProject.
func (mr *MockClientMockRecorder) SelectStarterProject(devfile, flags, fs, dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectStarterProject", reflect.TypeOf((*MockClient)(nil).SelectStarterProject), devfile, flags, fs, dir)
}

// Validate mocks base method.
func (m *MockClient) Validate(flags map[string]string, fs filesystem.Filesystem, dir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", flags, fs, dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockClientMockRecorder) Validate(flags, fs, dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockClient)(nil).Validate), flags, fs, dir)
}
