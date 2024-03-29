// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
)

//go:generate mockgen -package=mocks -destination=../mocks/route53recoverycontrolconfig.go -source=route53recoverycontrolconfig.go Route53recoverycontrolconfigClient
type Route53recoverycontrolconfigClient interface {
	DescribeCluster(context.Context, *route53recoverycontrolconfig.DescribeClusterInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DescribeClusterOutput, error)
	DescribeControlPanel(context.Context, *route53recoverycontrolconfig.DescribeControlPanelInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DescribeControlPanelOutput, error)
	DescribeRoutingControl(context.Context, *route53recoverycontrolconfig.DescribeRoutingControlInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DescribeRoutingControlOutput, error)
	DescribeSafetyRule(context.Context, *route53recoverycontrolconfig.DescribeSafetyRuleInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.DescribeSafetyRuleOutput, error)
	GetResourcePolicy(context.Context, *route53recoverycontrolconfig.GetResourcePolicyInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.GetResourcePolicyOutput, error)
	ListAssociatedRoute53HealthChecks(context.Context, *route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput, error)
	ListClusters(context.Context, *route53recoverycontrolconfig.ListClustersInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListClustersOutput, error)
	ListControlPanels(context.Context, *route53recoverycontrolconfig.ListControlPanelsInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListControlPanelsOutput, error)
	ListRoutingControls(context.Context, *route53recoverycontrolconfig.ListRoutingControlsInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListRoutingControlsOutput, error)
	ListSafetyRules(context.Context, *route53recoverycontrolconfig.ListSafetyRulesInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListSafetyRulesOutput, error)
	ListTagsForResource(context.Context, *route53recoverycontrolconfig.ListTagsForResourceInput, ...func(*route53recoverycontrolconfig.Options)) (*route53recoverycontrolconfig.ListTagsForResourceOutput, error)
}
