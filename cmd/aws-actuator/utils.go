package main

import (
	"bytes"
	"fmt"
	"github.com/openshift/cluster-api-actuator-pkg/pkg/e2e/framework"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	yaml "gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/record"
	"os/exec"
	machineactuator "sigs.k8s.io/cluster-api-provider-aws/pkg/actuators/machine"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/apis/awsproviderconfig/v1beta1"
	awsclient "sigs.k8s.io/cluster-api-provider-aws/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

type manifestParams struct{ ClusterID string }

func readMachineManifest(manifestParams *manifestParams, manifestLoc string) (*machinev1.Machine, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	machine := &machinev1.Machine{}
	manifestBytes, err := ioutil.ReadFile(manifestLoc)
	if err != nil {
		return nil, fmt.Errorf("unable to read %v: %v", manifestLoc, err)
	}
	t, err := template.New("machineuserdata").Parse(string(manifestBytes))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, *manifestParams)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(buf.Bytes(), &machine); err != nil {
		return nil, fmt.Errorf("unable to unmarshal %v: %v", manifestLoc, err)
	}
	return machine, nil
}
func readClusterResources(manifestParams *manifestParams, clusterLoc, machineLoc, awsCredentialSecretLoc, userDataLoc string) (*machinev1.Cluster, *machinev1.Machine, *apiv1.Secret, *apiv1.Secret, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	machine, err := readMachineManifest(manifestParams, machineLoc)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	cluster := &machinev1.Cluster{}
	bytes, err := ioutil.ReadFile(clusterLoc)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("cluster manifest %q: %v", clusterLoc, err)
	}
	if err := yaml.Unmarshal(bytes, &cluster); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("cluster manifest %q: %v", clusterLoc, err)
	}
	var awsCredentialsSecret *apiv1.Secret
	if awsCredentialSecretLoc != "" {
		awsCredentialsSecret = &apiv1.Secret{}
		bytes, err := ioutil.ReadFile(awsCredentialSecretLoc)
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("aws credentials manifest %q: %v", awsCredentialSecretLoc, err)
		}
		if err = yaml.Unmarshal(bytes, &awsCredentialsSecret); err != nil {
			return nil, nil, nil, nil, fmt.Errorf("aws credentials manifest %q: %v", awsCredentialSecretLoc, err)
		}
	}
	var userDataSecret *apiv1.Secret
	if userDataLoc != "" {
		userDataSecret = &apiv1.Secret{}
		bytes, err := ioutil.ReadFile(userDataLoc)
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("user data manifest %q: %v", userDataLoc, err)
		}
		if err = yaml.Unmarshal(bytes, &userDataSecret); err != nil {
			return nil, nil, nil, nil, fmt.Errorf("user data manifest %q: %v", userDataLoc, err)
		}
	}
	return cluster, machine, awsCredentialsSecret, userDataSecret, nil
}
func createSecretAndWait(f *framework.Framework, secret *apiv1.Secret) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := f.KubeClient.CoreV1().Secrets(secret.Namespace).Create(secret)
	if err != nil {
		return err
	}
	err = wait.Poll(framework.PollInterval, framework.PoolTimeout, func() (bool, error) {
		_, err := f.KubeClient.CoreV1().Secrets(secret.Namespace).Get(secret.Name, metav1.GetOptions{})
		return err == nil, nil
	})
	return err
}
func createActuator(machine *machinev1.Machine, awsCredentials, userData *apiv1.Secret) (*machineactuator.Actuator, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	objList := []runtime.Object{machine}
	if awsCredentials != nil {
		objList = append(objList, awsCredentials)
	}
	if userData != nil {
		objList = append(objList, userData)
	}
	fakeClient := fake.NewFakeClient(objList...)
	codec, err := v1beta1.NewCodec()
	if err != nil {
		return nil, err
	}
	params := machineactuator.ActuatorParams{Client: fakeClient, AwsClientBuilder: awsclient.NewClient, Codec: codec, EventRecorder: &record.FakeRecorder{}}
	actuator, err := machineactuator.NewActuator(params)
	if err != nil {
		return nil, err
	}
	return actuator, nil
}
func cmdRun(binaryPath string, args ...string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := exec.Command(binaryPath, args...)
	return cmd.CombinedOutput()
}
