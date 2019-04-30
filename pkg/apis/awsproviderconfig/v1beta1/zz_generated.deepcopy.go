package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AWSMachineProviderCondition) DeepCopyInto(out *AWSMachineProviderCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *AWSMachineProviderCondition) DeepCopy() *AWSMachineProviderCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *AWSMachineProviderConfig) DeepCopyInto(out *AWSMachineProviderConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AMI.DeepCopyInto(&out.AMI)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagSpecification, len(*in))
		copy(*out, *in)
	}
	if in.IAMInstanceProfile != nil {
		in, out := &in.IAMInstanceProfile, &out.IAMInstanceProfile
		*out = new(AWSResourceReference)
		(*in).DeepCopyInto(*out)
	}
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
		*out = new(bool)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]AWSResourceReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Subnet.DeepCopyInto(&out.Subnet)
	out.Placement = in.Placement
	if in.LoadBalancers != nil {
		in, out := &in.LoadBalancers, &out.LoadBalancers
		*out = make([]LoadBalancerReference, len(*in))
		copy(*out, *in)
	}
	if in.BlockDevices != nil {
		in, out := &in.BlockDevices, &out.BlockDevices
		*out = make([]BlockDeviceMappingSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *AWSMachineProviderConfig) DeepCopy() *AWSMachineProviderConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *AWSMachineProviderConfig) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *AWSMachineProviderConfigList) DeepCopyInto(out *AWSMachineProviderConfigList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSMachineProviderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *AWSMachineProviderConfigList) DeepCopy() *AWSMachineProviderConfigList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderConfigList)
	in.DeepCopyInto(out)
	return out
}
func (in *AWSMachineProviderConfigList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *AWSMachineProviderStatus) DeepCopyInto(out *AWSMachineProviderStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AWSMachineProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *AWSMachineProviderStatus) DeepCopy() *AWSMachineProviderStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *AWSMachineProviderStatus) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *AWSResourceReference) DeepCopyInto(out *AWSResourceReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]Filter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *AWSResourceReference) DeepCopy() *AWSResourceReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AWSResourceReference)
	in.DeepCopyInto(out)
	return out
}
func (in *BlockDeviceMappingSpec) DeepCopyInto(out *BlockDeviceMappingSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.EBS != nil {
		in, out := &in.EBS, &out.EBS
		*out = new(EBSBlockDeviceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
	return
}
func (in *BlockDeviceMappingSpec) DeepCopy() *BlockDeviceMappingSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BlockDeviceMappingSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *EBSBlockDeviceSpec) DeepCopyInto(out *EBSBlockDeviceSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	return
}
func (in *EBSBlockDeviceSpec) DeepCopy() *EBSBlockDeviceSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(EBSBlockDeviceSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *Filter) DeepCopyInto(out *Filter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *Filter) DeepCopy() *Filter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}
func (in *LoadBalancerReference) DeepCopyInto(out *LoadBalancerReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *LoadBalancerReference) DeepCopy() *LoadBalancerReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LoadBalancerReference)
	in.DeepCopyInto(out)
	return out
}
func (in *Placement) DeepCopyInto(out *Placement) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *Placement) DeepCopy() *Placement {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}
func (in *TagSpecification) DeepCopyInto(out *TagSpecification) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *TagSpecification) DeepCopy() *TagSpecification {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TagSpecification)
	in.DeepCopyInto(out)
	return out
}
