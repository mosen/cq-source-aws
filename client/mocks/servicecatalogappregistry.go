// Code generated by MockGen. DO NOT EDIT.
// Source: servicecatalogappregistry.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	servicecatalogappregistry "github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	gomock "github.com/golang/mock/gomock"
)

// MockServicecatalogappregistryClient is a mock of ServicecatalogappregistryClient interface.
type MockServicecatalogappregistryClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicecatalogappregistryClientMockRecorder
}

// MockServicecatalogappregistryClientMockRecorder is the mock recorder for MockServicecatalogappregistryClient.
type MockServicecatalogappregistryClientMockRecorder struct {
	mock *MockServicecatalogappregistryClient
}

// NewMockServicecatalogappregistryClient creates a new mock instance.
func NewMockServicecatalogappregistryClient(ctrl *gomock.Controller) *MockServicecatalogappregistryClient {
	mock := &MockServicecatalogappregistryClient{ctrl: ctrl}
	mock.recorder = &MockServicecatalogappregistryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicecatalogappregistryClient) EXPECT() *MockServicecatalogappregistryClientMockRecorder {
	return m.recorder
}

// GetApplication mocks base method.
func (m *MockServicecatalogappregistryClient) GetApplication(arg0 context.Context, arg1 *servicecatalogappregistry.GetApplicationInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetApplicationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetApplication")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApplication", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.GetApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockServicecatalogappregistryClientMockRecorder) GetApplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).GetApplication), varargs...)
}

// GetAssociatedResource mocks base method.
func (m *MockServicecatalogappregistryClient) GetAssociatedResource(arg0 context.Context, arg1 *servicecatalogappregistry.GetAssociatedResourceInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAssociatedResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAssociatedResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAssociatedResource", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.GetAssociatedResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssociatedResource indicates an expected call of GetAssociatedResource.
func (mr *MockServicecatalogappregistryClientMockRecorder) GetAssociatedResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssociatedResource", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).GetAssociatedResource), varargs...)
}

// GetAttributeGroup mocks base method.
func (m *MockServicecatalogappregistryClient) GetAttributeGroup(arg0 context.Context, arg1 *servicecatalogappregistry.GetAttributeGroupInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAttributeGroupOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAttributeGroup")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAttributeGroup", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.GetAttributeGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttributeGroup indicates an expected call of GetAttributeGroup.
func (mr *MockServicecatalogappregistryClientMockRecorder) GetAttributeGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttributeGroup", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).GetAttributeGroup), varargs...)
}

// GetConfiguration mocks base method.
func (m *MockServicecatalogappregistryClient) GetConfiguration(arg0 context.Context, arg1 *servicecatalogappregistry.GetConfigurationInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfiguration", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.GetConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockServicecatalogappregistryClientMockRecorder) GetConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).GetConfiguration), varargs...)
}

// ListApplications mocks base method.
func (m *MockServicecatalogappregistryClient) ListApplications(arg0 context.Context, arg1 *servicecatalogappregistry.ListApplicationsInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListApplicationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListApplications")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApplications", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockServicecatalogappregistryClientMockRecorder) ListApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListApplications), varargs...)
}

// ListAssociatedAttributeGroups mocks base method.
func (m *MockServicecatalogappregistryClient) ListAssociatedAttributeGroups(arg0 context.Context, arg1 *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAssociatedAttributeGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssociatedAttributeGroups", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssociatedAttributeGroups indicates an expected call of ListAssociatedAttributeGroups.
func (mr *MockServicecatalogappregistryClientMockRecorder) ListAssociatedAttributeGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssociatedAttributeGroups", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListAssociatedAttributeGroups), varargs...)
}

// ListAssociatedResources mocks base method.
func (m *MockServicecatalogappregistryClient) ListAssociatedResources(arg0 context.Context, arg1 *servicecatalogappregistry.ListAssociatedResourcesInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedResourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAssociatedResources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssociatedResources", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListAssociatedResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssociatedResources indicates an expected call of ListAssociatedResources.
func (mr *MockServicecatalogappregistryClientMockRecorder) ListAssociatedResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssociatedResources", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListAssociatedResources), varargs...)
}

// ListAttributeGroups mocks base method.
func (m *MockServicecatalogappregistryClient) ListAttributeGroups(arg0 context.Context, arg1 *servicecatalogappregistry.ListAttributeGroupsInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAttributeGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttributeGroups", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListAttributeGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttributeGroups indicates an expected call of ListAttributeGroups.
func (mr *MockServicecatalogappregistryClientMockRecorder) ListAttributeGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttributeGroups", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListAttributeGroups), varargs...)
}

// ListAttributeGroupsForApplication mocks base method.
func (m *MockServicecatalogappregistryClient) ListAttributeGroupsForApplication(arg0 context.Context, arg1 *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAttributeGroupsForApplication")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttributeGroupsForApplication", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttributeGroupsForApplication indicates an expected call of ListAttributeGroupsForApplication.
func (mr *MockServicecatalogappregistryClientMockRecorder) ListAttributeGroupsForApplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttributeGroupsForApplication", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListAttributeGroupsForApplication), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockServicecatalogappregistryClient) ListTagsForResource(arg0 context.Context, arg1 *servicecatalogappregistry.ListTagsForResourceInput, arg2 ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicecatalogappregistry.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*servicecatalogappregistry.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockServicecatalogappregistryClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockServicecatalogappregistryClient)(nil).ListTagsForResource), varargs...)
}
