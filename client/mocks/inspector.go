// Code generated by MockGen. DO NOT EDIT.
// Source: inspector.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	inspector "github.com/aws/aws-sdk-go-v2/service/inspector"
	gomock "github.com/golang/mock/gomock"
)

// MockInspectorClient is a mock of InspectorClient interface.
type MockInspectorClient struct {
	ctrl     *gomock.Controller
	recorder *MockInspectorClientMockRecorder
}

// MockInspectorClientMockRecorder is the mock recorder for MockInspectorClient.
type MockInspectorClientMockRecorder struct {
	mock *MockInspectorClient
}

// NewMockInspectorClient creates a new mock instance.
func NewMockInspectorClient(ctrl *gomock.Controller) *MockInspectorClient {
	mock := &MockInspectorClient{ctrl: ctrl}
	mock.recorder = &MockInspectorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInspectorClient) EXPECT() *MockInspectorClientMockRecorder {
	return m.recorder
}

// DescribeAssessmentRuns mocks base method.
func (m *MockInspectorClient) DescribeAssessmentRuns(arg0 context.Context, arg1 *inspector.DescribeAssessmentRunsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeAssessmentRunsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAssessmentRuns")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAssessmentRuns", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeAssessmentRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAssessmentRuns indicates an expected call of DescribeAssessmentRuns.
func (mr *MockInspectorClientMockRecorder) DescribeAssessmentRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAssessmentRuns", reflect.TypeOf((*MockInspectorClient)(nil).DescribeAssessmentRuns), varargs...)
}

// DescribeAssessmentTargets mocks base method.
func (m *MockInspectorClient) DescribeAssessmentTargets(arg0 context.Context, arg1 *inspector.DescribeAssessmentTargetsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeAssessmentTargetsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAssessmentTargets")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAssessmentTargets", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeAssessmentTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAssessmentTargets indicates an expected call of DescribeAssessmentTargets.
func (mr *MockInspectorClientMockRecorder) DescribeAssessmentTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAssessmentTargets", reflect.TypeOf((*MockInspectorClient)(nil).DescribeAssessmentTargets), varargs...)
}

// DescribeAssessmentTemplates mocks base method.
func (m *MockInspectorClient) DescribeAssessmentTemplates(arg0 context.Context, arg1 *inspector.DescribeAssessmentTemplatesInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeAssessmentTemplatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAssessmentTemplates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAssessmentTemplates", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeAssessmentTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAssessmentTemplates indicates an expected call of DescribeAssessmentTemplates.
func (mr *MockInspectorClientMockRecorder) DescribeAssessmentTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAssessmentTemplates", reflect.TypeOf((*MockInspectorClient)(nil).DescribeAssessmentTemplates), varargs...)
}

// DescribeCrossAccountAccessRole mocks base method.
func (m *MockInspectorClient) DescribeCrossAccountAccessRole(arg0 context.Context, arg1 *inspector.DescribeCrossAccountAccessRoleInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCrossAccountAccessRole")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCrossAccountAccessRole", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeCrossAccountAccessRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCrossAccountAccessRole indicates an expected call of DescribeCrossAccountAccessRole.
func (mr *MockInspectorClientMockRecorder) DescribeCrossAccountAccessRole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCrossAccountAccessRole", reflect.TypeOf((*MockInspectorClient)(nil).DescribeCrossAccountAccessRole), varargs...)
}

// DescribeExclusions mocks base method.
func (m *MockInspectorClient) DescribeExclusions(arg0 context.Context, arg1 *inspector.DescribeExclusionsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeExclusionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeExclusions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeExclusions", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeExclusionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeExclusions indicates an expected call of DescribeExclusions.
func (mr *MockInspectorClientMockRecorder) DescribeExclusions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeExclusions", reflect.TypeOf((*MockInspectorClient)(nil).DescribeExclusions), varargs...)
}

// DescribeFindings mocks base method.
func (m *MockInspectorClient) DescribeFindings(arg0 context.Context, arg1 *inspector.DescribeFindingsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeFindingsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeFindings")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFindings", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFindings indicates an expected call of DescribeFindings.
func (mr *MockInspectorClientMockRecorder) DescribeFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFindings", reflect.TypeOf((*MockInspectorClient)(nil).DescribeFindings), varargs...)
}

// DescribeResourceGroups mocks base method.
func (m *MockInspectorClient) DescribeResourceGroups(arg0 context.Context, arg1 *inspector.DescribeResourceGroupsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeResourceGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeResourceGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeResourceGroups", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeResourceGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeResourceGroups indicates an expected call of DescribeResourceGroups.
func (mr *MockInspectorClientMockRecorder) DescribeResourceGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeResourceGroups", reflect.TypeOf((*MockInspectorClient)(nil).DescribeResourceGroups), varargs...)
}

// DescribeRulesPackages mocks base method.
func (m *MockInspectorClient) DescribeRulesPackages(arg0 context.Context, arg1 *inspector.DescribeRulesPackagesInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeRulesPackagesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeRulesPackages")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRulesPackages", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeRulesPackagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRulesPackages indicates an expected call of DescribeRulesPackages.
func (mr *MockInspectorClientMockRecorder) DescribeRulesPackages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRulesPackages", reflect.TypeOf((*MockInspectorClient)(nil).DescribeRulesPackages), varargs...)
}

// GetAssessmentReport mocks base method.
func (m *MockInspectorClient) GetAssessmentReport(arg0 context.Context, arg1 *inspector.GetAssessmentReportInput, arg2 ...func(*inspector.Options)) (*inspector.GetAssessmentReportOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAssessmentReport")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAssessmentReport", varargs...)
	ret0, _ := ret[0].(*inspector.GetAssessmentReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssessmentReport indicates an expected call of GetAssessmentReport.
func (mr *MockInspectorClientMockRecorder) GetAssessmentReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssessmentReport", reflect.TypeOf((*MockInspectorClient)(nil).GetAssessmentReport), varargs...)
}

// GetExclusionsPreview mocks base method.
func (m *MockInspectorClient) GetExclusionsPreview(arg0 context.Context, arg1 *inspector.GetExclusionsPreviewInput, arg2 ...func(*inspector.Options)) (*inspector.GetExclusionsPreviewOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetExclusionsPreview")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetExclusionsPreview", varargs...)
	ret0, _ := ret[0].(*inspector.GetExclusionsPreviewOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExclusionsPreview indicates an expected call of GetExclusionsPreview.
func (mr *MockInspectorClientMockRecorder) GetExclusionsPreview(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExclusionsPreview", reflect.TypeOf((*MockInspectorClient)(nil).GetExclusionsPreview), varargs...)
}

// GetTelemetryMetadata mocks base method.
func (m *MockInspectorClient) GetTelemetryMetadata(arg0 context.Context, arg1 *inspector.GetTelemetryMetadataInput, arg2 ...func(*inspector.Options)) (*inspector.GetTelemetryMetadataOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetTelemetryMetadata")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTelemetryMetadata", varargs...)
	ret0, _ := ret[0].(*inspector.GetTelemetryMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTelemetryMetadata indicates an expected call of GetTelemetryMetadata.
func (mr *MockInspectorClientMockRecorder) GetTelemetryMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTelemetryMetadata", reflect.TypeOf((*MockInspectorClient)(nil).GetTelemetryMetadata), varargs...)
}

// ListAssessmentRunAgents mocks base method.
func (m *MockInspectorClient) ListAssessmentRunAgents(arg0 context.Context, arg1 *inspector.ListAssessmentRunAgentsInput, arg2 ...func(*inspector.Options)) (*inspector.ListAssessmentRunAgentsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAssessmentRunAgents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssessmentRunAgents", varargs...)
	ret0, _ := ret[0].(*inspector.ListAssessmentRunAgentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssessmentRunAgents indicates an expected call of ListAssessmentRunAgents.
func (mr *MockInspectorClientMockRecorder) ListAssessmentRunAgents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssessmentRunAgents", reflect.TypeOf((*MockInspectorClient)(nil).ListAssessmentRunAgents), varargs...)
}

// ListAssessmentRuns mocks base method.
func (m *MockInspectorClient) ListAssessmentRuns(arg0 context.Context, arg1 *inspector.ListAssessmentRunsInput, arg2 ...func(*inspector.Options)) (*inspector.ListAssessmentRunsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAssessmentRuns")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssessmentRuns", varargs...)
	ret0, _ := ret[0].(*inspector.ListAssessmentRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssessmentRuns indicates an expected call of ListAssessmentRuns.
func (mr *MockInspectorClientMockRecorder) ListAssessmentRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssessmentRuns", reflect.TypeOf((*MockInspectorClient)(nil).ListAssessmentRuns), varargs...)
}

// ListAssessmentTargets mocks base method.
func (m *MockInspectorClient) ListAssessmentTargets(arg0 context.Context, arg1 *inspector.ListAssessmentTargetsInput, arg2 ...func(*inspector.Options)) (*inspector.ListAssessmentTargetsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAssessmentTargets")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssessmentTargets", varargs...)
	ret0, _ := ret[0].(*inspector.ListAssessmentTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssessmentTargets indicates an expected call of ListAssessmentTargets.
func (mr *MockInspectorClientMockRecorder) ListAssessmentTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssessmentTargets", reflect.TypeOf((*MockInspectorClient)(nil).ListAssessmentTargets), varargs...)
}

// ListAssessmentTemplates mocks base method.
func (m *MockInspectorClient) ListAssessmentTemplates(arg0 context.Context, arg1 *inspector.ListAssessmentTemplatesInput, arg2 ...func(*inspector.Options)) (*inspector.ListAssessmentTemplatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAssessmentTemplates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssessmentTemplates", varargs...)
	ret0, _ := ret[0].(*inspector.ListAssessmentTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssessmentTemplates indicates an expected call of ListAssessmentTemplates.
func (mr *MockInspectorClientMockRecorder) ListAssessmentTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssessmentTemplates", reflect.TypeOf((*MockInspectorClient)(nil).ListAssessmentTemplates), varargs...)
}

// ListEventSubscriptions mocks base method.
func (m *MockInspectorClient) ListEventSubscriptions(arg0 context.Context, arg1 *inspector.ListEventSubscriptionsInput, arg2 ...func(*inspector.Options)) (*inspector.ListEventSubscriptionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListEventSubscriptions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEventSubscriptions", varargs...)
	ret0, _ := ret[0].(*inspector.ListEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventSubscriptions indicates an expected call of ListEventSubscriptions.
func (mr *MockInspectorClientMockRecorder) ListEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventSubscriptions", reflect.TypeOf((*MockInspectorClient)(nil).ListEventSubscriptions), varargs...)
}

// ListExclusions mocks base method.
func (m *MockInspectorClient) ListExclusions(arg0 context.Context, arg1 *inspector.ListExclusionsInput, arg2 ...func(*inspector.Options)) (*inspector.ListExclusionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListExclusions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListExclusions", varargs...)
	ret0, _ := ret[0].(*inspector.ListExclusionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExclusions indicates an expected call of ListExclusions.
func (mr *MockInspectorClientMockRecorder) ListExclusions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExclusions", reflect.TypeOf((*MockInspectorClient)(nil).ListExclusions), varargs...)
}

// ListFindings mocks base method.
func (m *MockInspectorClient) ListFindings(arg0 context.Context, arg1 *inspector.ListFindingsInput, arg2 ...func(*inspector.Options)) (*inspector.ListFindingsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListFindings")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindings", varargs...)
	ret0, _ := ret[0].(*inspector.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFindings indicates an expected call of ListFindings.
func (mr *MockInspectorClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindings", reflect.TypeOf((*MockInspectorClient)(nil).ListFindings), varargs...)
}

// ListRulesPackages mocks base method.
func (m *MockInspectorClient) ListRulesPackages(arg0 context.Context, arg1 *inspector.ListRulesPackagesInput, arg2 ...func(*inspector.Options)) (*inspector.ListRulesPackagesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListRulesPackages")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRulesPackages", varargs...)
	ret0, _ := ret[0].(*inspector.ListRulesPackagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRulesPackages indicates an expected call of ListRulesPackages.
func (mr *MockInspectorClientMockRecorder) ListRulesPackages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRulesPackages", reflect.TypeOf((*MockInspectorClient)(nil).ListRulesPackages), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockInspectorClient) ListTagsForResource(arg0 context.Context, arg1 *inspector.ListTagsForResourceInput, arg2 ...func(*inspector.Options)) (*inspector.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &inspector.Options{}
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
	ret0, _ := ret[0].(*inspector.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockInspectorClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockInspectorClient)(nil).ListTagsForResource), varargs...)
}
