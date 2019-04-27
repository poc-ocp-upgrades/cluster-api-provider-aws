package machine

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	clusterFilterKeyPrefix	= "kubernetes.io/cluster/"
	clusterFilterValue	= "owned"
)

func awsTagFilter(name string) *string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return aws.String(fmt.Sprint("tag:", name))
}
func clusterFilterKey(name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprint(clusterFilterKeyPrefix, name)
}
func clusterFilter(name string) *ec2.Filter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ec2.Filter{Name: awsTagFilter(clusterFilterKey(name)), Values: aws.StringSlice([]string{clusterFilterValue})}
}
