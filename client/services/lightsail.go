// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

//go:generate mockgen -package=mocks -destination=../mocks/lightsail.go -source=lightsail.go LightsailClient
type LightsailClient interface {
	GetActiveNames(context.Context, *lightsail.GetActiveNamesInput, ...func(*lightsail.Options)) (*lightsail.GetActiveNamesOutput, error)
	GetAlarms(context.Context, *lightsail.GetAlarmsInput, ...func(*lightsail.Options)) (*lightsail.GetAlarmsOutput, error)
	GetAutoSnapshots(context.Context, *lightsail.GetAutoSnapshotsInput, ...func(*lightsail.Options)) (*lightsail.GetAutoSnapshotsOutput, error)
	GetBlueprints(context.Context, *lightsail.GetBlueprintsInput, ...func(*lightsail.Options)) (*lightsail.GetBlueprintsOutput, error)
	GetBucketAccessKeys(context.Context, *lightsail.GetBucketAccessKeysInput, ...func(*lightsail.Options)) (*lightsail.GetBucketAccessKeysOutput, error)
	GetBucketBundles(context.Context, *lightsail.GetBucketBundlesInput, ...func(*lightsail.Options)) (*lightsail.GetBucketBundlesOutput, error)
	GetBucketMetricData(context.Context, *lightsail.GetBucketMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetBucketMetricDataOutput, error)
	GetBuckets(context.Context, *lightsail.GetBucketsInput, ...func(*lightsail.Options)) (*lightsail.GetBucketsOutput, error)
	GetBundles(context.Context, *lightsail.GetBundlesInput, ...func(*lightsail.Options)) (*lightsail.GetBundlesOutput, error)
	GetCertificates(context.Context, *lightsail.GetCertificatesInput, ...func(*lightsail.Options)) (*lightsail.GetCertificatesOutput, error)
	GetCloudFormationStackRecords(context.Context, *lightsail.GetCloudFormationStackRecordsInput, ...func(*lightsail.Options)) (*lightsail.GetCloudFormationStackRecordsOutput, error)
	GetContactMethods(context.Context, *lightsail.GetContactMethodsInput, ...func(*lightsail.Options)) (*lightsail.GetContactMethodsOutput, error)
	GetContainerAPIMetadata(context.Context, *lightsail.GetContainerAPIMetadataInput, ...func(*lightsail.Options)) (*lightsail.GetContainerAPIMetadataOutput, error)
	GetContainerImages(context.Context, *lightsail.GetContainerImagesInput, ...func(*lightsail.Options)) (*lightsail.GetContainerImagesOutput, error)
	GetContainerLog(context.Context, *lightsail.GetContainerLogInput, ...func(*lightsail.Options)) (*lightsail.GetContainerLogOutput, error)
	GetContainerServiceDeployments(context.Context, *lightsail.GetContainerServiceDeploymentsInput, ...func(*lightsail.Options)) (*lightsail.GetContainerServiceDeploymentsOutput, error)
	GetContainerServiceMetricData(context.Context, *lightsail.GetContainerServiceMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetContainerServiceMetricDataOutput, error)
	GetContainerServicePowers(context.Context, *lightsail.GetContainerServicePowersInput, ...func(*lightsail.Options)) (*lightsail.GetContainerServicePowersOutput, error)
	GetContainerServices(context.Context, *lightsail.GetContainerServicesInput, ...func(*lightsail.Options)) (*lightsail.GetContainerServicesOutput, error)
	GetCostEstimate(context.Context, *lightsail.GetCostEstimateInput, ...func(*lightsail.Options)) (*lightsail.GetCostEstimateOutput, error)
	GetDisk(context.Context, *lightsail.GetDiskInput, ...func(*lightsail.Options)) (*lightsail.GetDiskOutput, error)
	GetDiskSnapshot(context.Context, *lightsail.GetDiskSnapshotInput, ...func(*lightsail.Options)) (*lightsail.GetDiskSnapshotOutput, error)
	GetDiskSnapshots(context.Context, *lightsail.GetDiskSnapshotsInput, ...func(*lightsail.Options)) (*lightsail.GetDiskSnapshotsOutput, error)
	GetDisks(context.Context, *lightsail.GetDisksInput, ...func(*lightsail.Options)) (*lightsail.GetDisksOutput, error)
	GetDistributionBundles(context.Context, *lightsail.GetDistributionBundlesInput, ...func(*lightsail.Options)) (*lightsail.GetDistributionBundlesOutput, error)
	GetDistributionLatestCacheReset(context.Context, *lightsail.GetDistributionLatestCacheResetInput, ...func(*lightsail.Options)) (*lightsail.GetDistributionLatestCacheResetOutput, error)
	GetDistributionMetricData(context.Context, *lightsail.GetDistributionMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetDistributionMetricDataOutput, error)
	GetDistributions(context.Context, *lightsail.GetDistributionsInput, ...func(*lightsail.Options)) (*lightsail.GetDistributionsOutput, error)
	GetDomain(context.Context, *lightsail.GetDomainInput, ...func(*lightsail.Options)) (*lightsail.GetDomainOutput, error)
	GetDomains(context.Context, *lightsail.GetDomainsInput, ...func(*lightsail.Options)) (*lightsail.GetDomainsOutput, error)
	GetExportSnapshotRecords(context.Context, *lightsail.GetExportSnapshotRecordsInput, ...func(*lightsail.Options)) (*lightsail.GetExportSnapshotRecordsOutput, error)
	GetInstance(context.Context, *lightsail.GetInstanceInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceOutput, error)
	GetInstanceAccessDetails(context.Context, *lightsail.GetInstanceAccessDetailsInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstanceMetricData(context.Context, *lightsail.GetInstanceMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceMetricDataOutput, error)
	GetInstancePortStates(context.Context, *lightsail.GetInstancePortStatesInput, ...func(*lightsail.Options)) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstanceSnapshot(context.Context, *lightsail.GetInstanceSnapshotInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceSnapshotOutput, error)
	GetInstanceSnapshots(context.Context, *lightsail.GetInstanceSnapshotsInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetInstanceState(context.Context, *lightsail.GetInstanceStateInput, ...func(*lightsail.Options)) (*lightsail.GetInstanceStateOutput, error)
	GetInstances(context.Context, *lightsail.GetInstancesInput, ...func(*lightsail.Options)) (*lightsail.GetInstancesOutput, error)
	GetKeyPair(context.Context, *lightsail.GetKeyPairInput, ...func(*lightsail.Options)) (*lightsail.GetKeyPairOutput, error)
	GetKeyPairs(context.Context, *lightsail.GetKeyPairsInput, ...func(*lightsail.Options)) (*lightsail.GetKeyPairsOutput, error)
	GetLoadBalancer(context.Context, *lightsail.GetLoadBalancerInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerOutput, error)
	GetLoadBalancerMetricData(context.Context, *lightsail.GetLoadBalancerMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerMetricDataOutput, error)
	GetLoadBalancerTlsCertificates(context.Context, *lightsail.GetLoadBalancerTlsCertificatesInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
	GetLoadBalancerTlsPolicies(context.Context, *lightsail.GetLoadBalancerTlsPoliciesInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerTlsPoliciesOutput, error)
	GetLoadBalancers(context.Context, *lightsail.GetLoadBalancersInput, ...func(*lightsail.Options)) (*lightsail.GetLoadBalancersOutput, error)
	GetOperation(context.Context, *lightsail.GetOperationInput, ...func(*lightsail.Options)) (*lightsail.GetOperationOutput, error)
	GetOperations(context.Context, *lightsail.GetOperationsInput, ...func(*lightsail.Options)) (*lightsail.GetOperationsOutput, error)
	GetOperationsForResource(context.Context, *lightsail.GetOperationsForResourceInput, ...func(*lightsail.Options)) (*lightsail.GetOperationsForResourceOutput, error)
	GetRegions(context.Context, *lightsail.GetRegionsInput, ...func(*lightsail.Options)) (*lightsail.GetRegionsOutput, error)
	GetRelationalDatabase(context.Context, *lightsail.GetRelationalDatabaseInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseOutput, error)
	GetRelationalDatabaseBlueprints(context.Context, *lightsail.GetRelationalDatabaseBlueprintsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error)
	GetRelationalDatabaseBundles(context.Context, *lightsail.GetRelationalDatabaseBundlesInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseBundlesOutput, error)
	GetRelationalDatabaseEvents(context.Context, *lightsail.GetRelationalDatabaseEventsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseEventsOutput, error)
	GetRelationalDatabaseLogEvents(context.Context, *lightsail.GetRelationalDatabaseLogEventsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
	GetRelationalDatabaseLogStreams(context.Context, *lightsail.GetRelationalDatabaseLogStreamsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
	GetRelationalDatabaseMasterUserPassword(context.Context, *lightsail.GetRelationalDatabaseMasterUserPasswordInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error)
	GetRelationalDatabaseMetricData(context.Context, *lightsail.GetRelationalDatabaseMetricDataInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseMetricDataOutput, error)
	GetRelationalDatabaseParameters(context.Context, *lightsail.GetRelationalDatabaseParametersInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseParametersOutput, error)
	GetRelationalDatabaseSnapshot(context.Context, *lightsail.GetRelationalDatabaseSnapshotInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotOutput, error)
	GetRelationalDatabaseSnapshots(context.Context, *lightsail.GetRelationalDatabaseSnapshotsInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
	GetRelationalDatabases(context.Context, *lightsail.GetRelationalDatabasesInput, ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabasesOutput, error)
	GetStaticIp(context.Context, *lightsail.GetStaticIpInput, ...func(*lightsail.Options)) (*lightsail.GetStaticIpOutput, error)
	GetStaticIps(context.Context, *lightsail.GetStaticIpsInput, ...func(*lightsail.Options)) (*lightsail.GetStaticIpsOutput, error)
}
