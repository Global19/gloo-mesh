// Code generated by MockGen. DO NOT EDIT.
// Source: ./appmesh_client.go

// Package mock_appmesh is a generated GoMock package.
package mock_appmesh

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.aws.solo.io/v1alpha1"
	output "github.com/solo-io/skv2/contrib/pkg/output"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// UpsertVirtualNodes mocks base method
func (m *MockClient) UpsertVirtualNodes(ctx context.Context, mesh string, existingVirtualNodes sets.String, desiredVirtualNodes []*v1alpha1.VirtualNode, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpsertVirtualNodes", ctx, mesh, existingVirtualNodes, desiredVirtualNodes, errHandler)
}

// UpsertVirtualNodes indicates an expected call of UpsertVirtualNodes
func (mr *MockClientMockRecorder) UpsertVirtualNodes(ctx, mesh, existingVirtualNodes, desiredVirtualNodes, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertVirtualNodes", reflect.TypeOf((*MockClient)(nil).UpsertVirtualNodes), ctx, mesh, existingVirtualNodes, desiredVirtualNodes, errHandler)
}

// DeleteVirtualNodes mocks base method
func (m *MockClient) DeleteVirtualNodes(ctx context.Context, mesh string, staleVirtualNodes sets.String, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteVirtualNodes", ctx, mesh, staleVirtualNodes, errHandler)
}

// DeleteVirtualNodes indicates an expected call of DeleteVirtualNodes
func (mr *MockClientMockRecorder) DeleteVirtualNodes(ctx, mesh, staleVirtualNodes, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualNodes", reflect.TypeOf((*MockClient)(nil).DeleteVirtualNodes), ctx, mesh, staleVirtualNodes, errHandler)
}

// ListVirtualNodes mocks base method
func (m *MockClient) ListVirtualNodes(ctx context.Context, meshName string) (sets.String, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVirtualNodes", ctx, meshName)
	ret0, _ := ret[0].(sets.String)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualNodes indicates an expected call of ListVirtualNodes
func (mr *MockClientMockRecorder) ListVirtualNodes(ctx, meshName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualNodes", reflect.TypeOf((*MockClient)(nil).ListVirtualNodes), ctx, meshName)
}

// UpsertVirtualRouters mocks base method
func (m *MockClient) UpsertVirtualRouters(ctx context.Context, mesh string, existingVirtualRouters sets.String, desiredVirtualRouters []*v1alpha1.VirtualRouter, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpsertVirtualRouters", ctx, mesh, existingVirtualRouters, desiredVirtualRouters, errHandler)
}

// UpsertVirtualRouters indicates an expected call of UpsertVirtualRouters
func (mr *MockClientMockRecorder) UpsertVirtualRouters(ctx, mesh, existingVirtualRouters, desiredVirtualRouters, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertVirtualRouters", reflect.TypeOf((*MockClient)(nil).UpsertVirtualRouters), ctx, mesh, existingVirtualRouters, desiredVirtualRouters, errHandler)
}

// DeleteVirtualRouters mocks base method
func (m *MockClient) DeleteVirtualRouters(ctx context.Context, mesh string, staleVirtualRouters sets.String, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteVirtualRouters", ctx, mesh, staleVirtualRouters, errHandler)
}

// DeleteVirtualRouters indicates an expected call of DeleteVirtualRouters
func (mr *MockClientMockRecorder) DeleteVirtualRouters(ctx, mesh, staleVirtualRouters, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualRouters", reflect.TypeOf((*MockClient)(nil).DeleteVirtualRouters), ctx, mesh, staleVirtualRouters, errHandler)
}

// ListVirtualRouters mocks base method
func (m *MockClient) ListVirtualRouters(ctx context.Context, meshName string) (sets.String, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVirtualRouters", ctx, meshName)
	ret0, _ := ret[0].(sets.String)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualRouters indicates an expected call of ListVirtualRouters
func (mr *MockClientMockRecorder) ListVirtualRouters(ctx, meshName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualRouters", reflect.TypeOf((*MockClient)(nil).ListVirtualRouters), ctx, meshName)
}

// UpsertRoutes mocks base method
func (m *MockClient) UpsertRoutes(ctx context.Context, mesh string, existingRoutes sets.String, desiredRoutes []*v1alpha1.Route, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpsertRoutes", ctx, mesh, existingRoutes, desiredRoutes, errHandler)
}

// UpsertRoutes indicates an expected call of UpsertRoutes
func (mr *MockClientMockRecorder) UpsertRoutes(ctx, mesh, existingRoutes, desiredRoutes, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRoutes", reflect.TypeOf((*MockClient)(nil).UpsertRoutes), ctx, mesh, existingRoutes, desiredRoutes, errHandler)
}

// DeleteRoutes mocks base method
func (m *MockClient) DeleteRoutes(ctx context.Context, mesh string, staleRoutes sets.String, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteRoutes", ctx, mesh, staleRoutes, errHandler)
}

// DeleteRoutes indicates an expected call of DeleteRoutes
func (mr *MockClientMockRecorder) DeleteRoutes(ctx, mesh, staleRoutes, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoutes", reflect.TypeOf((*MockClient)(nil).DeleteRoutes), ctx, mesh, staleRoutes, errHandler)
}

// ListRoutes mocks base method
func (m *MockClient) ListRoutes(ctx context.Context, meshName string) (sets.String, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoutes", ctx, meshName)
	ret0, _ := ret[0].(sets.String)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoutes indicates an expected call of ListRoutes
func (mr *MockClientMockRecorder) ListRoutes(ctx, meshName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoutes", reflect.TypeOf((*MockClient)(nil).ListRoutes), ctx, meshName)
}

// UpsertVirtualServices mocks base method
func (m *MockClient) UpsertVirtualServices(ctx context.Context, mesh string, existingVirtualServices sets.String, desiredVirtualServices []*v1alpha1.VirtualService, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpsertVirtualServices", ctx, mesh, existingVirtualServices, desiredVirtualServices, errHandler)
}

// UpsertVirtualServices indicates an expected call of UpsertVirtualServices
func (mr *MockClientMockRecorder) UpsertVirtualServices(ctx, mesh, existingVirtualServices, desiredVirtualServices, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertVirtualServices", reflect.TypeOf((*MockClient)(nil).UpsertVirtualServices), ctx, mesh, existingVirtualServices, desiredVirtualServices, errHandler)
}

// DeleteVirtualServices mocks base method
func (m *MockClient) DeleteVirtualServices(ctx context.Context, mesh string, staleVirtualServices sets.String, errHandler output.ErrorHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteVirtualServices", ctx, mesh, staleVirtualServices, errHandler)
}

// DeleteVirtualServices indicates an expected call of DeleteVirtualServices
func (mr *MockClientMockRecorder) DeleteVirtualServices(ctx, mesh, staleVirtualServices, errHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualServices", reflect.TypeOf((*MockClient)(nil).DeleteVirtualServices), ctx, mesh, staleVirtualServices, errHandler)
}

// ListVirtualServices mocks base method
func (m *MockClient) ListVirtualServices(ctx context.Context, meshName string) (sets.String, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVirtualServices", ctx, meshName)
	ret0, _ := ret[0].(sets.String)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualServices indicates an expected call of ListVirtualServices
func (mr *MockClientMockRecorder) ListVirtualServices(ctx, meshName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualServices", reflect.TypeOf((*MockClient)(nil).ListVirtualServices), ctx, meshName)
}