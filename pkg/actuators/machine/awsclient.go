package machine

import (
	"fmt"
	"github.com/openshift/cluster-api-actuator-pkg/pkg/types"
	machinev1beta1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	awsclient "sigs.k8s.io/cluster-api-provider-aws/pkg/client"
)

type AwsClientWrapper struct{ client awsclient.Client }

var _ types.CloudProviderClient = &AwsClientWrapper{}

func NewAwsClientWrapper(client awsclient.Client) *AwsClientWrapper {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &AwsClientWrapper{client: client}
}
func (client *AwsClientWrapper) GetRunningInstances(machine *machinev1beta1.Machine) ([]interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	runningInstances, err := getRunningInstances(machine, client.client)
	if err != nil {
		return nil, err
	}
	var instances []interface{}
	for _, instance := range runningInstances {
		instances = append(instances, instance)
	}
	return instances, nil
}
func (client *AwsClientWrapper) GetPublicDNSName(machine *machinev1beta1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := getRunningInstance(machine, client.client)
	if err != nil {
		return "", err
	}
	if *instance.PublicDnsName == "" {
		return "", fmt.Errorf("machine instance public DNS name not set")
	}
	return *instance.PublicDnsName, nil
}
func (client *AwsClientWrapper) GetPrivateIP(machine *machinev1beta1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := getRunningInstance(machine, client.client)
	if err != nil {
		return "", err
	}
	if *instance.PrivateIpAddress == "" {
		return "", fmt.Errorf("machine instance public DNS name not set")
	}
	return *instance.PrivateIpAddress, nil
}
func (client *AwsClientWrapper) GetSecurityGroups(machine *machinev1beta1.Machine) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := getRunningInstance(machine, client.client)
	if err != nil {
		return nil, err
	}
	var groups []string
	for _, groupIdentifier := range instance.SecurityGroups {
		if *groupIdentifier.GroupName != "" {
			groups = append(groups, *groupIdentifier.GroupName)
		}
	}
	return groups, nil
}
func (client *AwsClientWrapper) GetIAMRole(machine *machinev1beta1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := getRunningInstance(machine, client.client)
	if err != nil {
		return "", err
	}
	if instance.IamInstanceProfile == nil {
		return "", err
	}
	return *instance.IamInstanceProfile.Id, nil
}
func (client *AwsClientWrapper) GetTags(machine *machinev1beta1.Machine) (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := getRunningInstance(machine, client.client)
	if err != nil {
		return nil, err
	}
	tags := make(map[string]string, len(instance.Tags))
	for _, tag := range instance.Tags {
		tags[*tag.Key] = *tag.Value
	}
	return tags, nil
}
func (client *AwsClientWrapper) GetSubnet(machine *machinev1beta1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := getRunningInstance(machine, client.client)
	if err != nil {
		return "", err
	}
	if instance.SubnetId == nil {
		return "", err
	}
	return *instance.SubnetId, nil
}
func (client *AwsClientWrapper) GetAvailabilityZone(machine *machinev1beta1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := getRunningInstance(machine, client.client)
	if err != nil {
		return "", err
	}
	if instance.Placement == nil {
		return "", err
	}
	return *instance.Placement.AvailabilityZone, nil
}
func (client *AwsClientWrapper) GetVolumes(machine *machinev1beta1.Machine) (map[string]map[string]interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := getRunningInstance(machine, client.client)
	if err != nil {
		return nil, err
	}
	if instance.BlockDeviceMappings == nil {
		return nil, err
	}
	volumes := make(map[string]map[string]interface{}, len(instance.BlockDeviceMappings))
	for _, blockDeviceMapping := range instance.BlockDeviceMappings {
		volume, err := getVolume(client.client, *blockDeviceMapping.Ebs.VolumeId)
		if err != nil {
			return volumes, err
		}
		volumes[*blockDeviceMapping.DeviceName] = map[string]interface{}{"id": *volume.VolumeId, "iops": *volume.Iops, "size": *volume.Size, "type": *volume.VolumeType}
	}
	return volumes, nil
}
