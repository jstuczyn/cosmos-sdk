// Code generated by MockGen. DO NOT EDIT.
// Source: core/appmodule/module.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAppModule is a mock of AppModule interface.
type MockAppModule struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleMockRecorder
}

// MockAppModuleMockRecorder is the mock recorder for MockAppModule.
type MockAppModuleMockRecorder struct {
	mock *MockAppModule
}

// NewMockAppModule creates a new mock instance.
func NewMockAppModule(ctrl *gomock.Controller) *MockAppModule {
	mock := &MockAppModule{ctrl: ctrl}
	mock.recorder = &MockAppModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppModule) EXPECT() *MockAppModuleMockRecorder {
	return m.recorder
}

// IsAppModule mocks base method.
func (m *MockAppModule) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockAppModuleMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockAppModule)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockAppModule) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockAppModuleMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockAppModule)(nil).IsOnePerModuleType))
}

// MockHasServices is a mock of HasServices interface.
type MockHasServices struct {
	ctrl     *gomock.Controller
	recorder *MockHasServicesMockRecorder
}

// MockHasServicesMockRecorder is the mock recorder for MockHasServices.
type MockHasServicesMockRecorder struct {
	mock *MockHasServices
}

// NewMockHasServices creates a new mock instance.
func NewMockHasServices(ctrl *gomock.Controller) *MockHasServices {
	mock := &MockHasServices{ctrl: ctrl}
	mock.recorder = &MockHasServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasServices) EXPECT() *MockHasServicesMockRecorder {
	return m.recorder
}

// IsAppModule mocks base method.
func (m *MockHasServices) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockHasServicesMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockHasServices)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockHasServices) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockHasServicesMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockHasServices)(nil).IsOnePerModuleType))
}

// RegisterServices mocks base method.
func (m *MockHasServices) RegisterServices(arg0 grpc.ServiceRegistrar) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterServices", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterServices indicates an expected call of RegisterServices.
func (mr *MockHasServicesMockRecorder) RegisterServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServices", reflect.TypeOf((*MockHasServices)(nil).RegisterServices), arg0)
}

// MockHasPrepareCheckState is a mock of HasPrepareCheckState interface.
type MockHasPrepareCheckState struct {
	ctrl     *gomock.Controller
	recorder *MockHasPrepareCheckStateMockRecorder
}

// MockHasPrepareCheckStateMockRecorder is the mock recorder for MockHasPrepareCheckState.
type MockHasPrepareCheckStateMockRecorder struct {
	mock *MockHasPrepareCheckState
}

// NewMockHasPrepareCheckState creates a new mock instance.
func NewMockHasPrepareCheckState(ctrl *gomock.Controller) *MockHasPrepareCheckState {
	mock := &MockHasPrepareCheckState{ctrl: ctrl}
	mock.recorder = &MockHasPrepareCheckStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasPrepareCheckState) EXPECT() *MockHasPrepareCheckStateMockRecorder {
	return m.recorder
}

// IsAppModule mocks base method.
func (m *MockHasPrepareCheckState) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockHasPrepareCheckStateMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockHasPrepareCheckState)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockHasPrepareCheckState) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockHasPrepareCheckStateMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockHasPrepareCheckState)(nil).IsOnePerModuleType))
}

// PrepareCheckState mocks base method.
func (m *MockHasPrepareCheckState) PrepareCheckState(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrepareCheckState", arg0)
}

// PrepareCheckState indicates an expected call of PrepareCheckState.
func (mr *MockHasPrepareCheckStateMockRecorder) PrepareCheckState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCheckState", reflect.TypeOf((*MockHasPrepareCheckState)(nil).PrepareCheckState), arg0)
}

// MockHasPrecommit is a mock of HasPrecommit interface.
type MockHasPrecommit struct {
	ctrl     *gomock.Controller
	recorder *MockHasPrecommitMockRecorder
}

// MockHasPrecommitMockRecorder is the mock recorder for MockHasPrecommit.
type MockHasPrecommitMockRecorder struct {
	mock *MockHasPrecommit
}

// NewMockHasPrecommit creates a new mock instance.
func NewMockHasPrecommit(ctrl *gomock.Controller) *MockHasPrecommit {
	mock := &MockHasPrecommit{ctrl: ctrl}
	mock.recorder = &MockHasPrecommitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasPrecommit) EXPECT() *MockHasPrecommitMockRecorder {
	return m.recorder
}

// IsAppModule mocks base method.
func (m *MockHasPrecommit) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockHasPrecommitMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockHasPrecommit)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockHasPrecommit) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockHasPrecommitMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockHasPrecommit)(nil).IsOnePerModuleType))
}

// Precommit mocks base method.
func (m *MockHasPrecommit) Precommit(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Precommit", arg0)
}

// Precommit indicates an expected call of Precommit.
func (mr *MockHasPrecommitMockRecorder) Precommit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Precommit", reflect.TypeOf((*MockHasPrecommit)(nil).Precommit), arg0)
}

// MockHasBeginBlocker is a mock of HasBeginBlocker interface.
type MockHasBeginBlocker struct {
	ctrl     *gomock.Controller
	recorder *MockHasBeginBlockerMockRecorder
}

// MockHasBeginBlockerMockRecorder is the mock recorder for MockHasBeginBlocker.
type MockHasBeginBlockerMockRecorder struct {
	mock *MockHasBeginBlocker
}

// NewMockHasBeginBlocker creates a new mock instance.
func NewMockHasBeginBlocker(ctrl *gomock.Controller) *MockHasBeginBlocker {
	mock := &MockHasBeginBlocker{ctrl: ctrl}
	mock.recorder = &MockHasBeginBlockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasBeginBlocker) EXPECT() *MockHasBeginBlockerMockRecorder {
	return m.recorder
}

// BeginBlock mocks base method.
func (m *MockHasBeginBlocker) BeginBlock(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginBlock indicates an expected call of BeginBlock.
func (mr *MockHasBeginBlockerMockRecorder) BeginBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockHasBeginBlocker)(nil).BeginBlock), arg0)
}

// IsAppModule mocks base method.
func (m *MockHasBeginBlocker) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockHasBeginBlockerMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockHasBeginBlocker)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockHasBeginBlocker) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockHasBeginBlockerMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockHasBeginBlocker)(nil).IsOnePerModuleType))
}

// MockHasEndBlocker is a mock of HasEndBlocker interface.
type MockHasEndBlocker struct {
	ctrl     *gomock.Controller
	recorder *MockHasEndBlockerMockRecorder
}

// MockHasEndBlockerMockRecorder is the mock recorder for MockHasEndBlocker.
type MockHasEndBlockerMockRecorder struct {
	mock *MockHasEndBlocker
}

// NewMockHasEndBlocker creates a new mock instance.
func NewMockHasEndBlocker(ctrl *gomock.Controller) *MockHasEndBlocker {
	mock := &MockHasEndBlocker{ctrl: ctrl}
	mock.recorder = &MockHasEndBlockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasEndBlocker) EXPECT() *MockHasEndBlockerMockRecorder {
	return m.recorder
}

// EndBlock mocks base method.
func (m *MockHasEndBlocker) EndBlock(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EndBlock indicates an expected call of EndBlock.
func (mr *MockHasEndBlockerMockRecorder) EndBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndBlock", reflect.TypeOf((*MockHasEndBlocker)(nil).EndBlock), arg0)
}

// IsAppModule mocks base method.
func (m *MockHasEndBlocker) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockHasEndBlockerMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockHasEndBlocker)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockHasEndBlocker) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockHasEndBlockerMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockHasEndBlocker)(nil).IsOnePerModuleType))
}
