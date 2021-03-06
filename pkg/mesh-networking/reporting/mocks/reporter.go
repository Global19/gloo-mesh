// Code generated by MockGen. DO NOT EDIT.
// Source: ./reporter.go

// Package mock_reporting is a generated GoMock package.
package mock_reporting

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	ezkube "github.com/solo-io/skv2/pkg/ezkube"
)

// MockReporter is a mock of Reporter interface
type MockReporter struct {
	ctrl     *gomock.Controller
	recorder *MockReporterMockRecorder
}

// MockReporterMockRecorder is the mock recorder for MockReporter
type MockReporterMockRecorder struct {
	mock *MockReporter
}

// NewMockReporter creates a new mock instance
func NewMockReporter(ctrl *gomock.Controller) *MockReporter {
	mock := &MockReporter{ctrl: ctrl}
	mock.recorder = &MockReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReporter) EXPECT() *MockReporterMockRecorder {
	return m.recorder
}

// ReportTrafficPolicyToTrafficTarget mocks base method
func (m *MockReporter) ReportTrafficPolicyToTrafficTarget(trafficTarget *v1alpha2.TrafficTarget, trafficPolicy ezkube.ResourceId, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportTrafficPolicyToTrafficTarget", trafficTarget, trafficPolicy, err)
}

// ReportTrafficPolicyToTrafficTarget indicates an expected call of ReportTrafficPolicyToTrafficTarget
func (mr *MockReporterMockRecorder) ReportTrafficPolicyToTrafficTarget(trafficTarget, trafficPolicy, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportTrafficPolicyToTrafficTarget", reflect.TypeOf((*MockReporter)(nil).ReportTrafficPolicyToTrafficTarget), trafficTarget, trafficPolicy, err)
}

// ReportAccessPolicyToTrafficTarget mocks base method
func (m *MockReporter) ReportAccessPolicyToTrafficTarget(trafficTarget *v1alpha2.TrafficTarget, accessPolicy ezkube.ResourceId, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportAccessPolicyToTrafficTarget", trafficTarget, accessPolicy, err)
}

// ReportAccessPolicyToTrafficTarget indicates an expected call of ReportAccessPolicyToTrafficTarget
func (mr *MockReporterMockRecorder) ReportAccessPolicyToTrafficTarget(trafficTarget, accessPolicy, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportAccessPolicyToTrafficTarget", reflect.TypeOf((*MockReporter)(nil).ReportAccessPolicyToTrafficTarget), trafficTarget, accessPolicy, err)
}

// ReportVirtualMeshToMesh mocks base method
func (m *MockReporter) ReportVirtualMeshToMesh(mesh *v1alpha2.Mesh, virtualMesh ezkube.ResourceId, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportVirtualMeshToMesh", mesh, virtualMesh, err)
}

// ReportVirtualMeshToMesh indicates an expected call of ReportVirtualMeshToMesh
func (mr *MockReporterMockRecorder) ReportVirtualMeshToMesh(mesh, virtualMesh, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportVirtualMeshToMesh", reflect.TypeOf((*MockReporter)(nil).ReportVirtualMeshToMesh), mesh, virtualMesh, err)
}

// ReportFailoverServiceToMesh mocks base method
func (m *MockReporter) ReportFailoverServiceToMesh(mesh *v1alpha2.Mesh, failoverService ezkube.ResourceId, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportFailoverServiceToMesh", mesh, failoverService, err)
}

// ReportFailoverServiceToMesh indicates an expected call of ReportFailoverServiceToMesh
func (mr *MockReporterMockRecorder) ReportFailoverServiceToMesh(mesh, failoverService, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportFailoverServiceToMesh", reflect.TypeOf((*MockReporter)(nil).ReportFailoverServiceToMesh), mesh, failoverService, err)
}

// ReportFailoverService mocks base method
func (m *MockReporter) ReportFailoverService(failoverService ezkube.ResourceId, errs []error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportFailoverService", failoverService, errs)
}

// ReportFailoverService indicates an expected call of ReportFailoverService
func (mr *MockReporterMockRecorder) ReportFailoverService(failoverService, errs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportFailoverService", reflect.TypeOf((*MockReporter)(nil).ReportFailoverService), failoverService, errs)
}
