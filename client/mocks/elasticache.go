// Code generated by MockGen. DO NOT EDIT.
// Source: elasticache.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	elasticache "github.com/aws/aws-sdk-go-v2/service/elasticache"
	gomock "github.com/golang/mock/gomock"
)

// MockElasticacheClient is a mock of ElasticacheClient interface.
type MockElasticacheClient struct {
	ctrl     *gomock.Controller
	recorder *MockElasticacheClientMockRecorder
}

// MockElasticacheClientMockRecorder is the mock recorder for MockElasticacheClient.
type MockElasticacheClientMockRecorder struct {
	mock *MockElasticacheClient
}

// NewMockElasticacheClient creates a new mock instance.
func NewMockElasticacheClient(ctrl *gomock.Controller) *MockElasticacheClient {
	mock := &MockElasticacheClient{ctrl: ctrl}
	mock.recorder = &MockElasticacheClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockElasticacheClient) EXPECT() *MockElasticacheClientMockRecorder {
	return m.recorder
}

// DescribeCacheClusters mocks base method.
func (m *MockElasticacheClient) DescribeCacheClusters(arg0 context.Context, arg1 *elasticache.DescribeCacheClustersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheClustersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCacheClusters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheClusters", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheClusters indicates an expected call of DescribeCacheClusters.
func (mr *MockElasticacheClientMockRecorder) DescribeCacheClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheClusters", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheClusters), varargs...)
}

// DescribeCacheEngineVersions mocks base method.
func (m *MockElasticacheClient) DescribeCacheEngineVersions(arg0 context.Context, arg1 *elasticache.DescribeCacheEngineVersionsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheEngineVersionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCacheEngineVersions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheEngineVersions", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheEngineVersions indicates an expected call of DescribeCacheEngineVersions.
func (mr *MockElasticacheClientMockRecorder) DescribeCacheEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheEngineVersions", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheEngineVersions), varargs...)
}

// DescribeCacheParameterGroups mocks base method.
func (m *MockElasticacheClient) DescribeCacheParameterGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheParameterGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheParameterGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCacheParameterGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheParameterGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheParameterGroups indicates an expected call of DescribeCacheParameterGroups.
func (mr *MockElasticacheClientMockRecorder) DescribeCacheParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheParameterGroups", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheParameterGroups), varargs...)
}

// DescribeCacheParameters mocks base method.
func (m *MockElasticacheClient) DescribeCacheParameters(arg0 context.Context, arg1 *elasticache.DescribeCacheParametersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheParametersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCacheParameters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheParameters", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheParameters indicates an expected call of DescribeCacheParameters.
func (mr *MockElasticacheClientMockRecorder) DescribeCacheParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheParameters", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheParameters), varargs...)
}

// DescribeCacheSecurityGroups mocks base method.
func (m *MockElasticacheClient) DescribeCacheSecurityGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheSecurityGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCacheSecurityGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheSecurityGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheSecurityGroups indicates an expected call of DescribeCacheSecurityGroups.
func (mr *MockElasticacheClientMockRecorder) DescribeCacheSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheSecurityGroups", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheSecurityGroups), varargs...)
}

// DescribeCacheSubnetGroups mocks base method.
func (m *MockElasticacheClient) DescribeCacheSubnetGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheSubnetGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCacheSubnetGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheSubnetGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCacheSubnetGroups indicates an expected call of DescribeCacheSubnetGroups.
func (mr *MockElasticacheClientMockRecorder) DescribeCacheSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheSubnetGroups", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeCacheSubnetGroups), varargs...)
}

// DescribeEngineDefaultParameters mocks base method.
func (m *MockElasticacheClient) DescribeEngineDefaultParameters(arg0 context.Context, arg1 *elasticache.DescribeEngineDefaultParametersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeEngineDefaultParametersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEngineDefaultParameters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEngineDefaultParameters", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeEngineDefaultParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEngineDefaultParameters indicates an expected call of DescribeEngineDefaultParameters.
func (mr *MockElasticacheClientMockRecorder) DescribeEngineDefaultParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEngineDefaultParameters", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeEngineDefaultParameters), varargs...)
}

// DescribeEvents mocks base method.
func (m *MockElasticacheClient) DescribeEvents(arg0 context.Context, arg1 *elasticache.DescribeEventsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeEventsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeEvents")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEvents", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEvents indicates an expected call of DescribeEvents.
func (mr *MockElasticacheClientMockRecorder) DescribeEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEvents", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeEvents), varargs...)
}

// DescribeGlobalReplicationGroups mocks base method.
func (m *MockElasticacheClient) DescribeGlobalReplicationGroups(arg0 context.Context, arg1 *elasticache.DescribeGlobalReplicationGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeGlobalReplicationGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGlobalReplicationGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeGlobalReplicationGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeGlobalReplicationGroups indicates an expected call of DescribeGlobalReplicationGroups.
func (mr *MockElasticacheClientMockRecorder) DescribeGlobalReplicationGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGlobalReplicationGroups", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeGlobalReplicationGroups), varargs...)
}

// DescribeReplicationGroups mocks base method.
func (m *MockElasticacheClient) DescribeReplicationGroups(arg0 context.Context, arg1 *elasticache.DescribeReplicationGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReplicationGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeReplicationGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReplicationGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReplicationGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReplicationGroups indicates an expected call of DescribeReplicationGroups.
func (mr *MockElasticacheClientMockRecorder) DescribeReplicationGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReplicationGroups", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeReplicationGroups), varargs...)
}

// DescribeReservedCacheNodes mocks base method.
func (m *MockElasticacheClient) DescribeReservedCacheNodes(arg0 context.Context, arg1 *elasticache.DescribeReservedCacheNodesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeReservedCacheNodes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReservedCacheNodes", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReservedCacheNodesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReservedCacheNodes indicates an expected call of DescribeReservedCacheNodes.
func (mr *MockElasticacheClientMockRecorder) DescribeReservedCacheNodes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReservedCacheNodes", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeReservedCacheNodes), varargs...)
}

// DescribeReservedCacheNodesOfferings mocks base method.
func (m *MockElasticacheClient) DescribeReservedCacheNodesOfferings(arg0 context.Context, arg1 *elasticache.DescribeReservedCacheNodesOfferingsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeReservedCacheNodesOfferings")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReservedCacheNodesOfferings", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReservedCacheNodesOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeReservedCacheNodesOfferings indicates an expected call of DescribeReservedCacheNodesOfferings.
func (mr *MockElasticacheClientMockRecorder) DescribeReservedCacheNodesOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReservedCacheNodesOfferings", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeReservedCacheNodesOfferings), varargs...)
}

// DescribeServerlessCacheSnapshots mocks base method.
func (m *MockElasticacheClient) DescribeServerlessCacheSnapshots(arg0 context.Context, arg1 *elasticache.DescribeServerlessCacheSnapshotsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeServerlessCacheSnapshotsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeServerlessCacheSnapshots")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServerlessCacheSnapshots", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeServerlessCacheSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServerlessCacheSnapshots indicates an expected call of DescribeServerlessCacheSnapshots.
func (mr *MockElasticacheClientMockRecorder) DescribeServerlessCacheSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServerlessCacheSnapshots", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeServerlessCacheSnapshots), varargs...)
}

// DescribeServerlessCaches mocks base method.
func (m *MockElasticacheClient) DescribeServerlessCaches(arg0 context.Context, arg1 *elasticache.DescribeServerlessCachesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeServerlessCachesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeServerlessCaches")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServerlessCaches", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeServerlessCachesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServerlessCaches indicates an expected call of DescribeServerlessCaches.
func (mr *MockElasticacheClientMockRecorder) DescribeServerlessCaches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServerlessCaches", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeServerlessCaches), varargs...)
}

// DescribeServiceUpdates mocks base method.
func (m *MockElasticacheClient) DescribeServiceUpdates(arg0 context.Context, arg1 *elasticache.DescribeServiceUpdatesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeServiceUpdatesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeServiceUpdates")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServiceUpdates", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeServiceUpdatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServiceUpdates indicates an expected call of DescribeServiceUpdates.
func (mr *MockElasticacheClientMockRecorder) DescribeServiceUpdates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServiceUpdates", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeServiceUpdates), varargs...)
}

// DescribeSnapshots mocks base method.
func (m *MockElasticacheClient) DescribeSnapshots(arg0 context.Context, arg1 *elasticache.DescribeSnapshotsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeSnapshotsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeSnapshots")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSnapshots", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSnapshots indicates an expected call of DescribeSnapshots.
func (mr *MockElasticacheClientMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshots", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeSnapshots), varargs...)
}

// DescribeUpdateActions mocks base method.
func (m *MockElasticacheClient) DescribeUpdateActions(arg0 context.Context, arg1 *elasticache.DescribeUpdateActionsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUpdateActionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeUpdateActions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUpdateActions", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUpdateActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUpdateActions indicates an expected call of DescribeUpdateActions.
func (mr *MockElasticacheClientMockRecorder) DescribeUpdateActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUpdateActions", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeUpdateActions), varargs...)
}

// DescribeUserGroups mocks base method.
func (m *MockElasticacheClient) DescribeUserGroups(arg0 context.Context, arg1 *elasticache.DescribeUserGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUserGroupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeUserGroups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUserGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUserGroups indicates an expected call of DescribeUserGroups.
func (mr *MockElasticacheClientMockRecorder) DescribeUserGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserGroups", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeUserGroups), varargs...)
}

// DescribeUsers mocks base method.
func (m *MockElasticacheClient) DescribeUsers(arg0 context.Context, arg1 *elasticache.DescribeUsersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUsersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeUsers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUsers", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUsers indicates an expected call of DescribeUsers.
func (mr *MockElasticacheClientMockRecorder) DescribeUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUsers", reflect.TypeOf((*MockElasticacheClient)(nil).DescribeUsers), varargs...)
}

// ListAllowedNodeTypeModifications mocks base method.
func (m *MockElasticacheClient) ListAllowedNodeTypeModifications(arg0 context.Context, arg1 *elasticache.ListAllowedNodeTypeModificationsInput, arg2 ...func(*elasticache.Options)) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListAllowedNodeTypeModifications")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAllowedNodeTypeModifications", varargs...)
	ret0, _ := ret[0].(*elasticache.ListAllowedNodeTypeModificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllowedNodeTypeModifications indicates an expected call of ListAllowedNodeTypeModifications.
func (mr *MockElasticacheClientMockRecorder) ListAllowedNodeTypeModifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedNodeTypeModifications", reflect.TypeOf((*MockElasticacheClient)(nil).ListAllowedNodeTypeModifications), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockElasticacheClient) ListTagsForResource(arg0 context.Context, arg1 *elasticache.ListTagsForResourceInput, arg2 ...func(*elasticache.Options)) (*elasticache.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &elasticache.Options{}
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
	ret0, _ := ret[0].(*elasticache.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockElasticacheClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockElasticacheClient)(nil).ListTagsForResource), varargs...)
}
