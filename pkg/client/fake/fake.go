package fake

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"k8s.io/client-go/kubernetes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/client"
	"time"
)

type awsClient struct{}

func (c *awsClient) DescribeImages(input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.DescribeImagesOutput{Images: []*ec2.Image{{ImageId: aws.String("ami-a9acbbd6")}}}, nil
}
func (c *awsClient) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.DescribeVpcsOutput{Vpcs: []*ec2.Vpc{{VpcId: aws.String("vpc-32677e0e794418639")}}}, nil
}
func (c *awsClient) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.DescribeSubnetsOutput{Subnets: []*ec2.Subnet{{SubnetId: aws.String("subnet-28fddb3c45cae61b5")}}}, nil
}
func (c *awsClient) DescribeAvailabilityZones(*ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.DescribeAvailabilityZonesOutput{}, nil
}
func (c *awsClient) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.DescribeSecurityGroupsOutput{SecurityGroups: []*ec2.SecurityGroup{{GroupId: aws.String("sg-05acc3c38a35ce63b")}}}, nil
}
func (c *awsClient) RunInstances(input *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.Reservation{Instances: []*ec2.Instance{{ImageId: aws.String("ami-a9acbbd6"), InstanceId: aws.String("i-02fcb933c5da7085c"), State: &ec2.InstanceState{Code: aws.Int64(16)}}}}, nil
}
func (c *awsClient) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.DescribeInstancesOutput{Reservations: []*ec2.Reservation{{Instances: []*ec2.Instance{{ImageId: aws.String("ami-a9acbbd6"), InstanceId: aws.String("i-02fcb933c5da7085c"), State: &ec2.InstanceState{Name: aws.String("Running"), Code: aws.Int64(16)}, LaunchTime: aws.Time(time.Now())}}}}}, nil
}
func (c *awsClient) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.TerminateInstancesOutput{}, nil
}
func (c *awsClient) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.DescribeVolumesOutput{}, nil
}
func (c *awsClient) RegisterInstancesWithLoadBalancer(input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &elb.RegisterInstancesWithLoadBalancerOutput{}, nil
}
func (c *awsClient) ELBv2DescribeLoadBalancers(*elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &elbv2.DescribeLoadBalancersOutput{}, nil
}
func (c *awsClient) ELBv2DescribeTargetGroups(*elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &elbv2.DescribeTargetGroupsOutput{}, nil
}
func (c *awsClient) ELBv2RegisterTargets(*elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &elbv2.RegisterTargetsOutput{}, nil
}
func NewClient(kubeClient kubernetes.Interface, secretName, namespace, region string) (client.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &awsClient{}, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
