package client

import (
	"context"
	"bytes"
	"net/http"
	"runtime"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/version"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
)

const (
	AwsCredsSecretIDKey	= "aws_access_key_id"
	AwsCredsSecretAccessKey	= "aws_secret_access_key"
)

type AwsClientBuilderFuncType func(client client.Client, secretName, namespace, region string) (Client, error)
type Client interface {
	DescribeImages(*ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error)
	DescribeVpcs(*ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)
	DescribeSubnets(*ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error)
	DescribeAvailabilityZones(*ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error)
	DescribeSecurityGroups(*ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error)
	RunInstances(*ec2.RunInstancesInput) (*ec2.Reservation, error)
	DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)
	TerminateInstances(*ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)
	DescribeVolumes(*ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error)
	RegisterInstancesWithLoadBalancer(*elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error)
	ELBv2DescribeLoadBalancers(*elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error)
	ELBv2DescribeTargetGroups(*elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error)
	ELBv2RegisterTargets(*elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error)
}
type awsClient struct {
	ec2Client	ec2iface.EC2API
	elbClient	elbiface.ELBAPI
	elbv2Client	elbv2iface.ELBV2API
}

func (c *awsClient) DescribeImages(input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.DescribeImages(input)
}
func (c *awsClient) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.DescribeVpcs(input)
}
func (c *awsClient) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.DescribeSubnets(input)
}
func (c *awsClient) DescribeAvailabilityZones(input *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.DescribeAvailabilityZones(input)
}
func (c *awsClient) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.DescribeSecurityGroups(input)
}
func (c *awsClient) RunInstances(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.RunInstances(input)
}
func (c *awsClient) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.DescribeInstances(input)
}
func (c *awsClient) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.TerminateInstances(input)
}
func (c *awsClient) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ec2Client.DescribeVolumes(input)
}
func (c *awsClient) RegisterInstancesWithLoadBalancer(input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.elbClient.RegisterInstancesWithLoadBalancer(input)
}
func (c *awsClient) ELBv2DescribeLoadBalancers(input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.elbv2Client.DescribeLoadBalancers(input)
}
func (c *awsClient) ELBv2DescribeTargetGroups(input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.elbv2Client.DescribeTargetGroups(input)
}
func (c *awsClient) ELBv2RegisterTargets(input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.elbv2Client.RegisterTargets(input)
}
func NewClient(ctrlRuntimeClient client.Client, secretName, namespace, region string) (Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	awsConfig := &aws.Config{Region: aws.String(region)}
	if secretName != "" {
		var secret corev1.Secret
		if err := ctrlRuntimeClient.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: secretName}, &secret); err != nil {
			return nil, err
		}
		accessKeyID, ok := secret.Data[AwsCredsSecretIDKey]
		if !ok {
			return nil, fmt.Errorf("AWS credentials secret %v did not contain key %v", secretName, AwsCredsSecretIDKey)
		}
		secretAccessKey, ok := secret.Data[AwsCredsSecretAccessKey]
		if !ok {
			return nil, fmt.Errorf("AWS credentials secret %v did not contain key %v", secretName, AwsCredsSecretAccessKey)
		}
		awsConfig.Credentials = credentials.NewStaticCredentials(string(accessKeyID), string(secretAccessKey), "")
	}
	s, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}
	s.Handlers.Build.PushBackNamed(addProviderVersionToUserAgent)
	return &awsClient{ec2Client: ec2.New(s), elbClient: elb.New(s), elbv2Client: elbv2.New(s)}, nil
}
func NewClientFromKeys(accessKey, secretAccessKey, region string) (Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	awsConfig := &aws.Config{Region: aws.String(region), Credentials: credentials.NewStaticCredentials(accessKey, secretAccessKey, "")}
	s, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}
	s.Handlers.Build.PushBackNamed(addProviderVersionToUserAgent)
	return &awsClient{ec2Client: ec2.New(s), elbClient: elb.New(s), elbv2Client: elbv2.New(s)}, nil
}

var addProviderVersionToUserAgent = request.NamedHandler{Name: "openshift.io/cluster-api-provider-aws", Fn: request.MakeAddToUserAgentHandler("openshift.io cluster-api-provider-aws", version.Version.String())}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
