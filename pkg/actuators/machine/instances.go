package machine

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/glog"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	providerconfigv1 "sigs.k8s.io/cluster-api-provider-aws/pkg/apis/awsproviderconfig/v1beta1"
	awsclient "sigs.k8s.io/cluster-api-provider-aws/pkg/client"
	"sort"
	"time"
)

func removeDuplicatedTags(tags []*ec2.Tag) []*ec2.Tag {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := make(map[string]bool)
	result := []*ec2.Tag{}
	for _, entry := range tags {
		if _, value := m[*entry.Key]; !value {
			m[*entry.Key] = true
			result = append(result, entry)
		}
	}
	return result
}
func removeStoppedMachine(machine *machinev1.Machine, client awsclient.Client) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instances, err := getStoppedInstances(machine, client)
	if err != nil {
		glog.Errorf("Error getting stopped instances: %v", err)
		return fmt.Errorf("error getting stopped instances: %v", err)
	}
	if len(instances) == 0 {
		glog.Infof("No stopped instances found for machine %v", machine.Name)
		return nil
	}
	return terminateInstances(client, instances)
}
func buildEC2Filters(inputFilters []providerconfigv1.Filter) []*ec2.Filter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	filters := make([]*ec2.Filter, len(inputFilters))
	for i, f := range inputFilters {
		values := make([]*string, len(f.Values))
		for j, v := range f.Values {
			values[j] = aws.String(v)
		}
		filters[i] = &ec2.Filter{Name: aws.String(f.Name), Values: values}
	}
	return filters
}
func getSecurityGroupsIDs(securityGroups []providerconfigv1.AWSResourceReference, client awsclient.Client) ([]*string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var securityGroupIDs []*string
	for _, g := range securityGroups {
		if g.ID != nil {
			securityGroupIDs = append(securityGroupIDs, g.ID)
		} else if g.Filters != nil {
			glog.Info("Describing security groups based on filters")
			describeSecurityGroupsRequest := ec2.DescribeSecurityGroupsInput{Filters: buildEC2Filters(g.Filters)}
			describeSecurityGroupsResult, err := client.DescribeSecurityGroups(&describeSecurityGroupsRequest)
			if err != nil {
				glog.Errorf("error describing security groups: %v", err)
				return nil, fmt.Errorf("error describing security groups: %v", err)
			}
			for _, g := range describeSecurityGroupsResult.SecurityGroups {
				groupID := *g.GroupId
				securityGroupIDs = append(securityGroupIDs, &groupID)
			}
		}
	}
	if len(securityGroups) == 0 {
		glog.Info("No security group found")
	}
	return securityGroupIDs, nil
}
func getSubnetIDs(subnet providerconfigv1.AWSResourceReference, availabilityZone string, client awsclient.Client) ([]*string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var subnetIDs []*string
	if subnet.ID != nil {
		subnetIDs = append(subnetIDs, subnet.ID)
	} else {
		var filters []providerconfigv1.Filter
		if availabilityZone != "" {
			_, err := client.DescribeAvailabilityZones(&ec2.DescribeAvailabilityZonesInput{ZoneNames: []*string{aws.String(availabilityZone)}})
			if err != nil {
				glog.Errorf("error describing availability zones: %v", err)
				return nil, fmt.Errorf("error describing availability zones: %v", err)
			}
			filters = append(filters, providerconfigv1.Filter{Name: "availabilityZone", Values: []string{availabilityZone}})
		}
		filters = append(filters, subnet.Filters...)
		glog.Info("Describing subnets based on filters")
		describeSubnetRequest := ec2.DescribeSubnetsInput{Filters: buildEC2Filters(filters)}
		describeSubnetResult, err := client.DescribeSubnets(&describeSubnetRequest)
		if err != nil {
			glog.Errorf("error describing subnetes: %v", err)
			return nil, fmt.Errorf("error describing subnets: %v", err)
		}
		for _, n := range describeSubnetResult.Subnets {
			subnetID := *n.SubnetId
			subnetIDs = append(subnetIDs, &subnetID)
		}
	}
	if len(subnetIDs) == 0 {
		return nil, fmt.Errorf("no subnet IDs were found")
	}
	return subnetIDs, nil
}
func getAMI(AMI providerconfigv1.AWSResourceReference, client awsclient.Client) (*string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if AMI.ID != nil {
		amiID := AMI.ID
		glog.Infof("Using AMI %s", *amiID)
		return amiID, nil
	}
	if len(AMI.Filters) > 0 {
		glog.Info("Describing AMI based on filters")
		describeImagesRequest := ec2.DescribeImagesInput{Filters: buildEC2Filters(AMI.Filters)}
		describeAMIResult, err := client.DescribeImages(&describeImagesRequest)
		if err != nil {
			glog.Errorf("error describing AMI: %v", err)
			return nil, fmt.Errorf("error describing AMI: %v", err)
		}
		if len(describeAMIResult.Images) < 1 {
			glog.Errorf("no image for given filters not found")
			return nil, fmt.Errorf("no image for given filters not found")
		}
		latestImage := describeAMIResult.Images[0]
		latestTime, err := time.Parse(time.RFC3339, *latestImage.CreationDate)
		if err != nil {
			glog.Errorf("unable to parse time for %q AMI: %v", *latestImage.ImageId, err)
			return nil, fmt.Errorf("unable to parse time for %q AMI: %v", *latestImage.ImageId, err)
		}
		for _, image := range describeAMIResult.Images[1:] {
			imageTime, err := time.Parse(time.RFC3339, *image.CreationDate)
			if err != nil {
				glog.Errorf("unable to parse time for %q AMI: %v", *image.ImageId, err)
				return nil, fmt.Errorf("unable to parse time for %q AMI: %v", *image.ImageId, err)
			}
			if latestTime.Before(imageTime) {
				latestImage = image
				latestTime = imageTime
			}
		}
		return latestImage.ImageId, nil
	}
	return nil, fmt.Errorf("AMI ID or AMI filters need to be specified")
}
func getBlockDeviceMappings(blockDeviceMappings []providerconfigv1.BlockDeviceMappingSpec, AMI string, client awsclient.Client) ([]*ec2.BlockDeviceMapping, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(blockDeviceMappings) == 0 {
		return []*ec2.BlockDeviceMapping{}, nil
	}
	describeImagesRequest := ec2.DescribeImagesInput{ImageIds: []*string{&AMI}}
	describeAMIResult, err := client.DescribeImages(&describeImagesRequest)
	if err != nil {
		glog.Errorf("Error describing AMI: %v", err)
		return nil, fmt.Errorf("error describing AMI: %v", err)
	}
	if len(describeAMIResult.Images) < 1 {
		glog.Errorf("No image for given AMI was found")
		return nil, fmt.Errorf("no image for given AMI not found")
	}
	deviceName := describeAMIResult.Images[0].RootDeviceName
	volumeSize := blockDeviceMappings[0].EBS.VolumeSize
	volumeType := blockDeviceMappings[0].EBS.VolumeType
	blockDeviceMapping := ec2.BlockDeviceMapping{DeviceName: deviceName, Ebs: &ec2.EbsBlockDevice{VolumeSize: volumeSize, VolumeType: volumeType}}
	if *volumeType == "io1" {
		blockDeviceMapping.Ebs.Iops = blockDeviceMappings[0].EBS.Iops
	}
	return []*ec2.BlockDeviceMapping{&blockDeviceMapping}, nil
}
func launchInstance(machine *machinev1.Machine, machineProviderConfig *providerconfigv1.AWSMachineProviderConfig, userData []byte, client awsclient.Client) (*ec2.Instance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	amiID, err := getAMI(machineProviderConfig.AMI, client)
	if err != nil {
		return nil, fmt.Errorf("error getting AMI: %v,", err)
	}
	securityGroupsIDs, err := getSecurityGroupsIDs(machineProviderConfig.SecurityGroups, client)
	if err != nil {
		return nil, fmt.Errorf("error getting security groups IDs: %v,", err)
	}
	subnetIDs, err := getSubnetIDs(machineProviderConfig.Subnet, machineProviderConfig.Placement.AvailabilityZone, client)
	if err != nil {
		return nil, fmt.Errorf("error getting subnet IDs: %v,", err)
	}
	if len(subnetIDs) > 1 {
		glog.Warningf("More than one subnet id returned, only first one will be used")
	}
	var networkInterfaces = []*ec2.InstanceNetworkInterfaceSpecification{{DeviceIndex: aws.Int64(machineProviderConfig.DeviceIndex), AssociatePublicIpAddress: machineProviderConfig.PublicIP, SubnetId: subnetIDs[0], Groups: securityGroupsIDs}}
	blockDeviceMappings, err := getBlockDeviceMappings(machineProviderConfig.BlockDevices, *amiID, client)
	if err != nil {
		return nil, fmt.Errorf("error getting blockDeviceMappings: %v,", err)
	}
	clusterID, ok := getClusterID(machine)
	if !ok {
		glog.Errorf("Unable to get cluster ID for machine: %q", machine.Name)
		return nil, err
	}
	rawTagList := []*ec2.Tag{}
	for _, tag := range machineProviderConfig.Tags {
		rawTagList = append(rawTagList, &ec2.Tag{Key: aws.String(tag.Name), Value: aws.String(tag.Value)})
	}
	rawTagList = append(rawTagList, []*ec2.Tag{{Key: aws.String("kubernetes.io/cluster/" + clusterID), Value: aws.String("owned")}, {Key: aws.String("Name"), Value: aws.String(machine.Name)}}...)
	tagList := removeDuplicatedTags(rawTagList)
	tagInstance := &ec2.TagSpecification{ResourceType: aws.String("instance"), Tags: tagList}
	tagVolume := &ec2.TagSpecification{ResourceType: aws.String("volume"), Tags: tagList}
	userDataEnc := base64.StdEncoding.EncodeToString(userData)
	var iamInstanceProfile *ec2.IamInstanceProfileSpecification
	if machineProviderConfig.IAMInstanceProfile != nil && machineProviderConfig.IAMInstanceProfile.ID != nil {
		iamInstanceProfile = &ec2.IamInstanceProfileSpecification{Name: aws.String(*machineProviderConfig.IAMInstanceProfile.ID)}
	}
	var placement *ec2.Placement
	if machineProviderConfig.Placement.AvailabilityZone != "" && machineProviderConfig.Subnet.ID == nil {
		placement = &ec2.Placement{AvailabilityZone: aws.String(machineProviderConfig.Placement.AvailabilityZone)}
	}
	inputConfig := ec2.RunInstancesInput{ImageId: amiID, InstanceType: aws.String(machineProviderConfig.InstanceType), MinCount: aws.Int64(1), MaxCount: aws.Int64(1), KeyName: machineProviderConfig.KeyName, IamInstanceProfile: iamInstanceProfile, TagSpecifications: []*ec2.TagSpecification{tagInstance, tagVolume}, NetworkInterfaces: networkInterfaces, UserData: &userDataEnc, Placement: placement}
	if len(blockDeviceMappings) > 0 {
		inputConfig.BlockDeviceMappings = blockDeviceMappings
	}
	runResult, err := client.RunInstances(&inputConfig)
	if err != nil {
		glog.Errorf("Error creating EC2 instance: %v", err)
		return nil, fmt.Errorf("error creating EC2 instance: %v", err)
	}
	if runResult == nil || len(runResult.Instances) != 1 {
		glog.Errorf("Unexpected reservation creating instances: %v", runResult)
		return nil, fmt.Errorf("unexpected reservation creating instance")
	}
	return runResult.Instances[0], nil
}

type instanceList []*ec2.Instance

func (il instanceList) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(il)
}
func (il instanceList) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	il[i], il[j] = il[j], il[i]
}
func (il instanceList) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if il[i].LaunchTime == nil && il[j].LaunchTime == nil {
		return false
	}
	if il[i].LaunchTime != nil && il[j].LaunchTime == nil {
		return false
	}
	if il[i].LaunchTime == nil && il[j].LaunchTime != nil {
		return true
	}
	return (*il[i].LaunchTime).After(*il[j].LaunchTime)
}
func sortInstances(instances []*ec2.Instance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sort.Sort(instanceList(instances))
}
