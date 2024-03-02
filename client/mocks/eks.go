// Code generated by MockGen. DO NOT EDIT.
// Source: eks.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	eks "github.com/aws/aws-sdk-go-v2/service/eks"
	gomock "github.com/golang/mock/gomock"
)

// MockEksClient is a mock of EksClient interface.
type MockEksClient struct {
	ctrl     *gomock.Controller
	recorder *MockEksClientMockRecorder
}

// MockEksClientMockRecorder is the mock recorder for MockEksClient.
type MockEksClientMockRecorder struct {
	mock *MockEksClient
}

// NewMockEksClient creates a new mock instance.
func NewMockEksClient(ctrl *gomock.Controller) *MockEksClient {
	mock := &MockEksClient{ctrl: ctrl}
	mock.recorder = &MockEksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEksClient) EXPECT() *MockEksClientMockRecorder {
	return m.recorder
}

// DescribeAddon mocks base method.
func (m *MockEksClient) DescribeAddon(arg0 context.Context, arg1 *eks.DescribeAddonInput, arg2 ...func(*eks.Options)) (*eks.DescribeAddonOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAddon")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAddon", varargs...)
	ret0, _ := ret[0].(*eks.DescribeAddonOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAddon indicates an expected call of DescribeAddon.
func (mr *MockEksClientMockRecorder) DescribeAddon(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAddon", reflect.TypeOf((*MockEksClient)(nil).DescribeAddon), varargs...)
}

// DescribeAddonConfiguration mocks base method.
func (m *MockEksClient) DescribeAddonConfiguration(arg0 context.Context, arg1 *eks.DescribeAddonConfigurationInput, arg2 ...func(*eks.Options)) (*eks.DescribeAddonConfigurationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAddonConfiguration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAddonConfiguration", varargs...)
	ret0, _ := ret[0].(*eks.DescribeAddonConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAddonConfiguration indicates an expected call of DescribeAddonConfiguration.
func (mr *MockEksClientMockRecorder) DescribeAddonConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAddonConfiguration", reflect.TypeOf((*MockEksClient)(nil).DescribeAddonConfiguration), varargs...)
}

// DescribeAddonVersions mocks base method.
func (m *MockEksClient) DescribeAddonVersions(arg0 context.Context, arg1 *eks.DescribeAddonVersionsInput, arg2 ...func(*eks.Options)) (*eks.DescribeAddonVersionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeAddonVersions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAddonVersions", varargs...)
	ret0, _ := ret[0].(*eks.DescribeAddonVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAddonVersions indicates an expected call of DescribeAddonVersions.
func (mr *MockEksClientMockRecorder) DescribeAddonVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAddonVersions", reflect.TypeOf((*MockEksClient)(nil).DescribeAddonVersions), varargs...)
}

// DescribeCluster mocks base method.
func (m *MockEksClient) DescribeCluster(arg0 context.Context, arg1 *eks.DescribeClusterInput, arg2 ...func(*eks.Options)) (*eks.DescribeClusterOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCluster")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCluster", varargs...)
	ret0, _ := ret[0].(*eks.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCluster indicates an expected call of DescribeCluster.
func (mr *MockEksClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCluster", reflect.TypeOf((*MockEksClient)(nil).DescribeCluster), varargs...)
}

// DescribeEksAnywhereSubscription mocks base method.
func (m *MockEksClient) DescribeEksAnywhereSubscription(arg0 context.Context, arg1 *eks.DescribeEksAnywhereSubscriptionInput, arg2 ...func(*eks.Options)) (*eks.DescribeEksAnywhereSubscriptionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEksAnywhereSubscription")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEksAnywhereSubscription", varargs...)
	ret0, _ := ret[0].(*eks.DescribeEksAnywhereSubscriptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEksAnywhereSubscription indicates an expected call of DescribeEksAnywhereSubscription.
func (mr *MockEksClientMockRecorder) DescribeEksAnywhereSubscription(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEksAnywhereSubscription", reflect.TypeOf((*MockEksClient)(nil).DescribeEksAnywhereSubscription), varargs...)
}

// DescribeFargateProfile mocks base method.
func (m *MockEksClient) DescribeFargateProfile(arg0 context.Context, arg1 *eks.DescribeFargateProfileInput, arg2 ...func(*eks.Options)) (*eks.DescribeFargateProfileOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeFargateProfile")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFargateProfile", varargs...)
	ret0, _ := ret[0].(*eks.DescribeFargateProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFargateProfile indicates an expected call of DescribeFargateProfile.
func (mr *MockEksClientMockRecorder) DescribeFargateProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFargateProfile", reflect.TypeOf((*MockEksClient)(nil).DescribeFargateProfile), varargs...)
}

// DescribeIdentityProviderConfig mocks base method.
func (m *MockEksClient) DescribeIdentityProviderConfig(arg0 context.Context, arg1 *eks.DescribeIdentityProviderConfigInput, arg2 ...func(*eks.Options)) (*eks.DescribeIdentityProviderConfigOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeIdentityProviderConfig")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeIdentityProviderConfig", varargs...)
	ret0, _ := ret[0].(*eks.DescribeIdentityProviderConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeIdentityProviderConfig indicates an expected call of DescribeIdentityProviderConfig.
func (mr *MockEksClientMockRecorder) DescribeIdentityProviderConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeIdentityProviderConfig", reflect.TypeOf((*MockEksClient)(nil).DescribeIdentityProviderConfig), varargs...)
}

// DescribeNodegroup mocks base method.
func (m *MockEksClient) DescribeNodegroup(arg0 context.Context, arg1 *eks.DescribeNodegroupInput, arg2 ...func(*eks.Options)) (*eks.DescribeNodegroupOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeNodegroup")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNodegroup", varargs...)
	ret0, _ := ret[0].(*eks.DescribeNodegroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNodegroup indicates an expected call of DescribeNodegroup.
func (mr *MockEksClientMockRecorder) DescribeNodegroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNodegroup", reflect.TypeOf((*MockEksClient)(nil).DescribeNodegroup), varargs...)
}

// DescribePodIdentityAssociation mocks base method.
func (m *MockEksClient) DescribePodIdentityAssociation(arg0 context.Context, arg1 *eks.DescribePodIdentityAssociationInput, arg2 ...func(*eks.Options)) (*eks.DescribePodIdentityAssociationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribePodIdentityAssociation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePodIdentityAssociation", varargs...)
	ret0, _ := ret[0].(*eks.DescribePodIdentityAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePodIdentityAssociation indicates an expected call of DescribePodIdentityAssociation.
func (mr *MockEksClientMockRecorder) DescribePodIdentityAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePodIdentityAssociation", reflect.TypeOf((*MockEksClient)(nil).DescribePodIdentityAssociation), varargs...)
}

// DescribeUpdate mocks base method.
func (m *MockEksClient) DescribeUpdate(arg0 context.Context, arg1 *eks.DescribeUpdateInput, arg2 ...func(*eks.Options)) (*eks.DescribeUpdateOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeUpdate")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUpdate", varargs...)
	ret0, _ := ret[0].(*eks.DescribeUpdateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUpdate indicates an expected call of DescribeUpdate.
func (mr *MockEksClientMockRecorder) DescribeUpdate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUpdate", reflect.TypeOf((*MockEksClient)(nil).DescribeUpdate), varargs...)
}

// ListAddons mocks base method.
func (m *MockEksClient) ListAddons(arg0 context.Context, arg1 *eks.ListAddonsInput, arg2 ...func(*eks.Options)) (*eks.ListAddonsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAddons")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAddons", varargs...)
	ret0, _ := ret[0].(*eks.ListAddonsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAddons indicates an expected call of ListAddons.
func (mr *MockEksClientMockRecorder) ListAddons(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAddons", reflect.TypeOf((*MockEksClient)(nil).ListAddons), varargs...)
}

// ListClusters mocks base method.
func (m *MockEksClient) ListClusters(arg0 context.Context, arg1 *eks.ListClustersInput, arg2 ...func(*eks.Options)) (*eks.ListClustersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListClusters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*eks.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockEksClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEksClient)(nil).ListClusters), varargs...)
}

// ListEksAnywhereSubscriptions mocks base method.
func (m *MockEksClient) ListEksAnywhereSubscriptions(arg0 context.Context, arg1 *eks.ListEksAnywhereSubscriptionsInput, arg2 ...func(*eks.Options)) (*eks.ListEksAnywhereSubscriptionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListEksAnywhereSubscriptions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEksAnywhereSubscriptions", varargs...)
	ret0, _ := ret[0].(*eks.ListEksAnywhereSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEksAnywhereSubscriptions indicates an expected call of ListEksAnywhereSubscriptions.
func (mr *MockEksClientMockRecorder) ListEksAnywhereSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEksAnywhereSubscriptions", reflect.TypeOf((*MockEksClient)(nil).ListEksAnywhereSubscriptions), varargs...)
}

// ListFargateProfiles mocks base method.
func (m *MockEksClient) ListFargateProfiles(arg0 context.Context, arg1 *eks.ListFargateProfilesInput, arg2 ...func(*eks.Options)) (*eks.ListFargateProfilesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListFargateProfiles")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFargateProfiles", varargs...)
	ret0, _ := ret[0].(*eks.ListFargateProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFargateProfiles indicates an expected call of ListFargateProfiles.
func (mr *MockEksClientMockRecorder) ListFargateProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFargateProfiles", reflect.TypeOf((*MockEksClient)(nil).ListFargateProfiles), varargs...)
}

// ListIdentityProviderConfigs mocks base method.
func (m *MockEksClient) ListIdentityProviderConfigs(arg0 context.Context, arg1 *eks.ListIdentityProviderConfigsInput, arg2 ...func(*eks.Options)) (*eks.ListIdentityProviderConfigsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListIdentityProviderConfigs")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIdentityProviderConfigs", varargs...)
	ret0, _ := ret[0].(*eks.ListIdentityProviderConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIdentityProviderConfigs indicates an expected call of ListIdentityProviderConfigs.
func (mr *MockEksClientMockRecorder) ListIdentityProviderConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIdentityProviderConfigs", reflect.TypeOf((*MockEksClient)(nil).ListIdentityProviderConfigs), varargs...)
}

// ListNodegroups mocks base method.
func (m *MockEksClient) ListNodegroups(arg0 context.Context, arg1 *eks.ListNodegroupsInput, arg2 ...func(*eks.Options)) (*eks.ListNodegroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListNodegroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNodegroups", varargs...)
	ret0, _ := ret[0].(*eks.ListNodegroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodegroups indicates an expected call of ListNodegroups.
func (mr *MockEksClientMockRecorder) ListNodegroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodegroups", reflect.TypeOf((*MockEksClient)(nil).ListNodegroups), varargs...)
}

// ListPodIdentityAssociations mocks base method.
func (m *MockEksClient) ListPodIdentityAssociations(arg0 context.Context, arg1 *eks.ListPodIdentityAssociationsInput, arg2 ...func(*eks.Options)) (*eks.ListPodIdentityAssociationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPodIdentityAssociations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPodIdentityAssociations", varargs...)
	ret0, _ := ret[0].(*eks.ListPodIdentityAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodIdentityAssociations indicates an expected call of ListPodIdentityAssociations.
func (mr *MockEksClientMockRecorder) ListPodIdentityAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodIdentityAssociations", reflect.TypeOf((*MockEksClient)(nil).ListPodIdentityAssociations), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockEksClient) ListTagsForResource(arg0 context.Context, arg1 *eks.ListTagsForResourceInput, arg2 ...func(*eks.Options)) (*eks.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
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
	ret0, _ := ret[0].(*eks.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockEksClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockEksClient)(nil).ListTagsForResource), varargs...)
}

// ListUpdates mocks base method.
func (m *MockEksClient) ListUpdates(arg0 context.Context, arg1 *eks.ListUpdatesInput, arg2 ...func(*eks.Options)) (*eks.ListUpdatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &eks.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListUpdates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUpdates", varargs...)
	ret0, _ := ret[0].(*eks.ListUpdatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUpdates indicates an expected call of ListUpdates.
func (mr *MockEksClientMockRecorder) ListUpdates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUpdates", reflect.TypeOf((*MockEksClient)(nil).ListUpdates), varargs...)
}