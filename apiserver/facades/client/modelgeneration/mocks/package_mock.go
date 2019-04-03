// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/modelgeneration (interfaces: State,Model,Generation,Application)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	modelgeneration "github.com/juju/juju/apiserver/facades/client/modelgeneration"
	charm_v6 "gopkg.in/juju/charm.v6"
	names_v2 "gopkg.in/juju/names.v2"
	reflect "reflect"
)

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// Application mocks base method
func (m *MockState) Application(arg0 string) (modelgeneration.Application, error) {
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(modelgeneration.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application
func (mr *MockStateMockRecorder) Application(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockState)(nil).Application), arg0)
}

// ControllerTag mocks base method
func (m *MockState) ControllerTag() names_v2.ControllerTag {
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names_v2.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag
func (mr *MockStateMockRecorder) ControllerTag() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockState)(nil).ControllerTag))
}

// Model mocks base method
func (m *MockState) Model() (modelgeneration.Model, error) {
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(modelgeneration.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model
func (mr *MockStateMockRecorder) Model() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockState)(nil).Model))
}

// MockModel is a mock of Model interface
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// AddGeneration mocks base method
func (m *MockModel) AddGeneration(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AddGeneration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddGeneration indicates an expected call of AddGeneration
func (mr *MockModelMockRecorder) AddGeneration(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGeneration", reflect.TypeOf((*MockModel)(nil).AddGeneration), arg0, arg1)
}

// Branch mocks base method
func (m *MockModel) Branch(arg0 string) (modelgeneration.Generation, error) {
	ret := m.ctrl.Call(m, "Branch", arg0)
	ret0, _ := ret[0].(modelgeneration.Generation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Branch indicates an expected call of Branch
func (mr *MockModelMockRecorder) Branch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Branch", reflect.TypeOf((*MockModel)(nil).Branch), arg0)
}

// MockGeneration is a mock of Generation interface
type MockGeneration struct {
	ctrl     *gomock.Controller
	recorder *MockGenerationMockRecorder
}

// MockGenerationMockRecorder is the mock recorder for MockGeneration
type MockGenerationMockRecorder struct {
	mock *MockGeneration
}

// NewMockGeneration creates a new mock instance
func NewMockGeneration(ctrl *gomock.Controller) *MockGeneration {
	mock := &MockGeneration{ctrl: ctrl}
	mock.recorder = &MockGenerationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGeneration) EXPECT() *MockGenerationMockRecorder {
	return m.recorder
}

// AssignAllUnits mocks base method
func (m *MockGeneration) AssignAllUnits(arg0 string) error {
	ret := m.ctrl.Call(m, "AssignAllUnits", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignAllUnits indicates an expected call of AssignAllUnits
func (mr *MockGenerationMockRecorder) AssignAllUnits(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignAllUnits", reflect.TypeOf((*MockGeneration)(nil).AssignAllUnits), arg0)
}

// AssignUnit mocks base method
func (m *MockGeneration) AssignUnit(arg0 string) error {
	ret := m.ctrl.Call(m, "AssignUnit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignUnit indicates an expected call of AssignUnit
func (mr *MockGenerationMockRecorder) AssignUnit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignUnit", reflect.TypeOf((*MockGeneration)(nil).AssignUnit), arg0)
}

// AssignedUnits mocks base method
func (m *MockGeneration) AssignedUnits() map[string][]string {
	ret := m.ctrl.Call(m, "AssignedUnits")
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// AssignedUnits indicates an expected call of AssignedUnits
func (mr *MockGenerationMockRecorder) AssignedUnits() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedUnits", reflect.TypeOf((*MockGeneration)(nil).AssignedUnits))
}

// BranchName mocks base method
func (m *MockGeneration) BranchName() string {
	ret := m.ctrl.Call(m, "BranchName")
	ret0, _ := ret[0].(string)
	return ret0
}

// BranchName indicates an expected call of BranchName
func (mr *MockGenerationMockRecorder) BranchName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BranchName", reflect.TypeOf((*MockGeneration)(nil).BranchName))
}

// Commit mocks base method
func (m *MockGeneration) Commit(arg0 string) (int, error) {
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit
func (mr *MockGenerationMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockGeneration)(nil).Commit), arg0)
}

// Created mocks base method
func (m *MockGeneration) Created() int64 {
	ret := m.ctrl.Call(m, "Created")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Created indicates an expected call of Created
func (mr *MockGenerationMockRecorder) Created() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Created", reflect.TypeOf((*MockGeneration)(nil).Created))
}

// CreatedBy mocks base method
func (m *MockGeneration) CreatedBy() string {
	ret := m.ctrl.Call(m, "CreatedBy")
	ret0, _ := ret[0].(string)
	return ret0
}

// CreatedBy indicates an expected call of CreatedBy
func (mr *MockGenerationMockRecorder) CreatedBy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedBy", reflect.TypeOf((*MockGeneration)(nil).CreatedBy))
}

// Refresh mocks base method
func (m *MockGeneration) Refresh() error {
	ret := m.ctrl.Call(m, "Refresh")
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh
func (mr *MockGenerationMockRecorder) Refresh() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockGeneration)(nil).Refresh))
}

// MockApplication is a mock of Application interface
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// CharmConfig mocks base method
func (m *MockApplication) CharmConfig(arg0 string) (charm_v6.Settings, error) {
	ret := m.ctrl.Call(m, "CharmConfig", arg0)
	ret0, _ := ret[0].(charm_v6.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmConfig indicates an expected call of CharmConfig
func (mr *MockApplicationMockRecorder) CharmConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmConfig", reflect.TypeOf((*MockApplication)(nil).CharmConfig), arg0)
}
