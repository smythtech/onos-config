// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/store/snapshot/device/store.go

// Package store is a generated GoMock package.
package store

import (
	gomock "github.com/golang/mock/gomock"
	device "github.com/onosproject/onos-config/api/types/device"
	device0 "github.com/onosproject/onos-config/api/types/snapshot/device"
	stream "github.com/onosproject/onos-config/pkg/store/stream"
	reflect "reflect"
)

// MockDeviceSnapshotStore is a mock of Store interface
type MockDeviceSnapshotStore struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceSnapshotStoreMockRecorder
}

// MockDeviceSnapshotStoreMockRecorder is the mock recorder for MockDeviceSnapshotStore
type MockDeviceSnapshotStoreMockRecorder struct {
	mock *MockDeviceSnapshotStore
}

// NewMockDeviceSnapshotStore creates a new mock instance
func NewMockDeviceSnapshotStore(ctrl *gomock.Controller) *MockDeviceSnapshotStore {
	mock := &MockDeviceSnapshotStore{ctrl: ctrl}
	mock.recorder = &MockDeviceSnapshotStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceSnapshotStore) EXPECT() *MockDeviceSnapshotStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDeviceSnapshotStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockDeviceSnapshotStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).Close))
}

// Get mocks base method
func (m *MockDeviceSnapshotStore) Get(id device0.ID) (*device0.DeviceSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*device0.DeviceSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDeviceSnapshotStoreMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).Get), id)
}

// Create mocks base method
func (m *MockDeviceSnapshotStore) Create(snapshot *device0.DeviceSnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", snapshot)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockDeviceSnapshotStoreMockRecorder) Create(snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).Create), snapshot)
}

// Update mocks base method
func (m *MockDeviceSnapshotStore) Update(snapshot *device0.DeviceSnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", snapshot)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDeviceSnapshotStoreMockRecorder) Update(snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).Update), snapshot)
}

// Delete mocks base method
func (m *MockDeviceSnapshotStore) Delete(snapshot *device0.DeviceSnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", snapshot)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockDeviceSnapshotStoreMockRecorder) Delete(snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).Delete), snapshot)
}

// List mocks base method
func (m *MockDeviceSnapshotStore) List(arg0 chan<- *device0.DeviceSnapshot) (stream.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(stream.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDeviceSnapshotStoreMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).List), arg0)
}

// Watch mocks base method
func (m *MockDeviceSnapshotStore) Watch(arg0 chan<- stream.Event) (stream.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(stream.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockDeviceSnapshotStoreMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).Watch), arg0)
}

// Store mocks base method
func (m *MockDeviceSnapshotStore) Store(snapshot *device0.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", snapshot)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockDeviceSnapshotStoreMockRecorder) Store(snapshot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).Store), snapshot)
}

// Load mocks base method
func (m *MockDeviceSnapshotStore) Load(deviceID device.VersionedID) (*device0.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", deviceID)
	ret0, _ := ret[0].(*device0.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockDeviceSnapshotStoreMockRecorder) Load(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).Load), deviceID)
}

// LoadAll mocks base method
func (m *MockDeviceSnapshotStore) LoadAll(ch chan<- *device0.Snapshot) (stream.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAll", ch)
	ret0, _ := ret[0].(stream.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAll indicates an expected call of LoadAll
func (mr *MockDeviceSnapshotStoreMockRecorder) LoadAll(ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAll", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).LoadAll), ch)
}

// WatchAll mocks base method
func (m *MockDeviceSnapshotStore) WatchAll(arg0 chan<- stream.Event) (stream.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAll", arg0)
	ret0, _ := ret[0].(stream.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAll indicates an expected call of WatchAll
func (mr *MockDeviceSnapshotStoreMockRecorder) WatchAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAll", reflect.TypeOf((*MockDeviceSnapshotStore)(nil).WatchAll), arg0)
}
