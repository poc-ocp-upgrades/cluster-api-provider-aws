package machine

import (
	"fmt"
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"golang.org/x/net/context"
	"k8s.io/apimachinery/pkg/types"
	providerconfigv1 "sigs.k8s.io/cluster-api-provider-aws/pkg/apis/awsproviderconfig/v1beta1"
	awsclient "sigs.k8s.io/cluster-api-provider-aws/pkg/client"
)

func getRunningInstance(machine *machinev1.Machine, client awsclient.Client) (*ec2.Instance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instances, err := getRunningInstances(machine, client)
	if err != nil {
		return nil, err
	}
	if len(instances) == 0 {
		return nil, fmt.Errorf("no instance found for machine: %s", machine.Name)
	}
	sortInstances(instances)
	return instances[0], nil
}
func getRunningInstances(machine *machinev1.Machine, client awsclient.Client) ([]*ec2.Instance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runningInstanceStateFilter := []*string{aws.String(ec2.InstanceStateNameRunning), aws.String(ec2.InstanceStateNamePending)}
	return getInstances(machine, client, runningInstanceStateFilter)
}
func getStoppedInstances(machine *machinev1.Machine, client awsclient.Client) ([]*ec2.Instance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stoppedInstanceStateFilter := []*string{aws.String(ec2.InstanceStateNameStopped), aws.String(ec2.InstanceStateNameStopping)}
	return getInstances(machine, client, stoppedInstanceStateFilter)
}
func getInstances(machine *machinev1.Machine, client awsclient.Client, instanceStateFilter []*string) ([]*ec2.Instance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterID, ok := getClusterID(machine)
	if !ok {
		return []*ec2.Instance{}, fmt.Errorf("unable to get cluster ID for machine: %q", machine.Name)
	}
	requestFilters := []*ec2.Filter{{Name: awsTagFilter("Name"), Values: aws.StringSlice([]string{machine.Name})}, clusterFilter(clusterID)}
	if instanceStateFilter != nil {
		requestFilters = append(requestFilters, &ec2.Filter{Name: aws.String("instance-state-name"), Values: instanceStateFilter})
	}
	request := &ec2.DescribeInstancesInput{Filters: requestFilters}
	result, err := client.DescribeInstances(request)
	if err != nil {
		return []*ec2.Instance{}, err
	}
	instances := make([]*ec2.Instance, 0, len(result.Reservations))
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, instance)
		}
	}
	return instances, nil
}
func getVolume(client awsclient.Client, volumeID string) (*ec2.Volume, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	request := &ec2.DescribeVolumesInput{VolumeIds: []*string{&volumeID}}
	result, err := client.DescribeVolumes(request)
	if err != nil {
		return &ec2.Volume{}, err
	}
	if len(result.Volumes) != 1 {
		return &ec2.Volume{}, fmt.Errorf("unable to get volume ID: %q", volumeID)
	}
	return result.Volumes[0], nil
}
func terminateInstances(client awsclient.Client, instances []*ec2.Instance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceIDs := []*string{}
	for _, instance := range instances {
		glog.Infof("Cleaning up extraneous instance for machine: %v, state: %v, launchTime: %v", *instance.InstanceId, *instance.State.Name, *instance.LaunchTime)
		instanceIDs = append(instanceIDs, instance.InstanceId)
	}
	for _, instanceID := range instanceIDs {
		glog.Infof("Terminating %v instance", *instanceID)
	}
	terminateInstancesRequest := &ec2.TerminateInstancesInput{InstanceIds: instanceIDs}
	_, err := client.TerminateInstances(terminateInstancesRequest)
	if err != nil {
		glog.Errorf("Error terminating instances: %v", err)
		return fmt.Errorf("error terminating instances: %v", err)
	}
	return nil
}
func providerConfigFromMachine(machine *machinev1.Machine, codec *providerconfigv1.AWSProviderConfigCodec) (*providerconfigv1.AWSMachineProviderConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if machine.Spec.ProviderSpec.Value == nil {
		return nil, fmt.Errorf("unable to find machine provider config: Spec.ProviderSpec.Value is not set")
	}
	var config providerconfigv1.AWSMachineProviderConfig
	if err := codec.DecodeProviderSpec(&machine.Spec.ProviderSpec, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
func (a *Actuator) isMaster(machine *machinev1.Machine) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if machine.Status.NodeRef == nil {
		glog.Errorf("NodeRef not found in machine %s", machine.Name)
		return false, nil
	}
	node := &corev1.Node{}
	nodeKey := types.NamespacedName{Namespace: machine.Status.NodeRef.Namespace, Name: machine.Status.NodeRef.Name}
	err := a.client.Get(context.Background(), nodeKey, node)
	if err != nil {
		return false, fmt.Errorf("failed to get node from machine %s", machine.Name)
	}
	if _, exists := node.Labels["node-role.kubernetes.io/master"]; exists {
		return true, nil
	}
	return false, nil
}

type updateConditionCheck func(oldReason, oldMessage, newReason, newMessage string) bool

func updateConditionAlways(_, _, _, _ string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func updateConditionNever(_, _, _, _ string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func updateConditionIfReasonOrMessageChange(oldReason, oldMessage, newReason, newMessage string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return oldReason != newReason || oldMessage != newMessage
}
func shouldUpdateCondition(oldStatus corev1.ConditionStatus, oldReason, oldMessage string, newStatus corev1.ConditionStatus, newReason, newMessage string, updateConditionCheck updateConditionCheck) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if oldStatus != newStatus {
		return true
	}
	return updateConditionCheck(oldReason, oldMessage, newReason, newMessage)
}
func setAWSMachineProviderCondition(conditions []providerconfigv1.AWSMachineProviderCondition, conditionType providerconfigv1.AWSMachineProviderConditionType, status corev1.ConditionStatus, reason string, message string, updateConditionCheck updateConditionCheck) []providerconfigv1.AWSMachineProviderCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := metav1.Now()
	existingCondition := findAWSMachineProviderCondition(conditions, conditionType)
	if existingCondition == nil {
		if status == corev1.ConditionTrue {
			conditions = append(conditions, providerconfigv1.AWSMachineProviderCondition{Type: conditionType, Status: status, Reason: reason, Message: message, LastTransitionTime: now, LastProbeTime: now})
		}
	} else {
		if shouldUpdateCondition(existingCondition.Status, existingCondition.Reason, existingCondition.Message, status, reason, message, updateConditionCheck) {
			if existingCondition.Status != status {
				existingCondition.LastTransitionTime = now
			}
			existingCondition.Status = status
			existingCondition.Reason = reason
			existingCondition.Message = message
			existingCondition.LastProbeTime = now
		}
	}
	return conditions
}
func findAWSMachineProviderCondition(conditions []providerconfigv1.AWSMachineProviderCondition, conditionType providerconfigv1.AWSMachineProviderConditionType) *providerconfigv1.AWSMachineProviderCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, condition := range conditions {
		if condition.Type == conditionType {
			return &conditions[i]
		}
	}
	return nil
}
