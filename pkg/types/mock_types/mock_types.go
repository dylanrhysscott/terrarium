// Code generated by MockGen. DO NOT EDIT.
// Source: ../types.go

// Package mock_types is a generated GoMock package.
package mock_types

import (
	context "context"
	http "net/http"
	reflect "reflect"

	types "github.com/dylanrhysscott/terrarium/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAPIErrorWriter is a mock of APIErrorWriter interface.
type MockAPIErrorWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAPIErrorWriterMockRecorder
}

// MockAPIErrorWriterMockRecorder is the mock recorder for MockAPIErrorWriter.
type MockAPIErrorWriterMockRecorder struct {
	mock *MockAPIErrorWriter
}

// NewMockAPIErrorWriter creates a new mock instance.
func NewMockAPIErrorWriter(ctrl *gomock.Controller) *MockAPIErrorWriter {
	mock := &MockAPIErrorWriter{ctrl: ctrl}
	mock.recorder = &MockAPIErrorWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIErrorWriter) EXPECT() *MockAPIErrorWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockAPIErrorWriter) Write(rw http.ResponseWriter, err error, statusCode int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", rw, err, statusCode)
}

// Write indicates an expected call of Write.
func (mr *MockAPIErrorWriterMockRecorder) Write(rw, err, statusCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockAPIErrorWriter)(nil).Write), rw, err, statusCode)
}

// MockAPIResponseWriter is a mock of APIResponseWriter interface.
type MockAPIResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockAPIResponseWriterMockRecorder
}

// MockAPIResponseWriterMockRecorder is the mock recorder for MockAPIResponseWriter.
type MockAPIResponseWriterMockRecorder struct {
	mock *MockAPIResponseWriter
}

// NewMockAPIResponseWriter creates a new mock instance.
func NewMockAPIResponseWriter(ctrl *gomock.Controller) *MockAPIResponseWriter {
	mock := &MockAPIResponseWriter{ctrl: ctrl}
	mock.recorder = &MockAPIResponseWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIResponseWriter) EXPECT() *MockAPIResponseWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockAPIResponseWriter) Write(rw http.ResponseWriter, data interface{}, statusCode int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", rw, data, statusCode)
}

// Write indicates an expected call of Write.
func (mr *MockAPIResponseWriterMockRecorder) Write(rw, data, statusCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockAPIResponseWriter)(nil).Write), rw, data, statusCode)
}

// MockOrganizationStore is a mock of OrganizationStore interface.
type MockOrganizationStore struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationStoreMockRecorder
}

// MockOrganizationStoreMockRecorder is the mock recorder for MockOrganizationStore.
type MockOrganizationStoreMockRecorder struct {
	mock *MockOrganizationStore
}

// NewMockOrganizationStore creates a new mock instance.
func NewMockOrganizationStore(ctrl *gomock.Controller) *MockOrganizationStore {
	mock := &MockOrganizationStore{ctrl: ctrl}
	mock.recorder = &MockOrganizationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationStore) EXPECT() *MockOrganizationStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrganizationStore) Create(name, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrganizationStoreMockRecorder) Create(name, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrganizationStore)(nil).Create), name, email)
}

// Delete mocks base method.
func (m *MockOrganizationStore) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrganizationStoreMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrganizationStore)(nil).Delete), id)
}

// Init mocks base method.
func (m *MockOrganizationStore) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockOrganizationStoreMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockOrganizationStore)(nil).Init))
}

// ReadAll mocks base method.
func (m *MockOrganizationStore) ReadAll(limit, offset int) ([]*types.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", limit, offset)
	ret0, _ := ret[0].([]*types.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockOrganizationStoreMockRecorder) ReadAll(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockOrganizationStore)(nil).ReadAll), limit, offset)
}

// ReadOne mocks base method.
func (m *MockOrganizationStore) ReadOne(id string) (*types.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOne", id)
	ret0, _ := ret[0].(*types.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadOne indicates an expected call of ReadOne.
func (mr *MockOrganizationStoreMockRecorder) ReadOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOne", reflect.TypeOf((*MockOrganizationStore)(nil).ReadOne), id)
}

// Update mocks base method.
func (m *MockOrganizationStore) Update(id, name, email string) (*types.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, name, email)
	ret0, _ := ret[0].(*types.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockOrganizationStoreMockRecorder) Update(id, name, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrganizationStore)(nil).Update), id, name, email)
}

// MockTerrariumDriver is a mock of TerrariumDriver interface.
type MockTerrariumDriver struct {
	ctrl     *gomock.Controller
	recorder *MockTerrariumDriverMockRecorder
}

// MockTerrariumDriverMockRecorder is the mock recorder for MockTerrariumDriver.
type MockTerrariumDriverMockRecorder struct {
	mock *MockTerrariumDriver
}

// NewMockTerrariumDriver creates a new mock instance.
func NewMockTerrariumDriver(ctrl *gomock.Controller) *MockTerrariumDriver {
	mock := &MockTerrariumDriver{ctrl: ctrl}
	mock.recorder = &MockTerrariumDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTerrariumDriver) EXPECT() *MockTerrariumDriverMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockTerrariumDriver) Connect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockTerrariumDriverMockRecorder) Connect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockTerrariumDriver)(nil).Connect), ctx)
}

// Organizations mocks base method.
func (m *MockTerrariumDriver) Organizations() types.OrganizationStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Organizations")
	ret0, _ := ret[0].(types.OrganizationStore)
	return ret0
}

// Organizations indicates an expected call of Organizations.
func (mr *MockTerrariumDriverMockRecorder) Organizations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Organizations", reflect.TypeOf((*MockTerrariumDriver)(nil).Organizations))
}
