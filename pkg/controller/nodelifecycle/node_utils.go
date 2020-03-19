package nodelifecycle

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
	"time"

	v1 "k8s.io/api/core/v1"
)

// callKubeletHealthz call kubelet api using default port 10250;
// default timeout 1s
func callKubeletHealthz(node *v1.Node) (bool, error) {
	kubeletEndpoint := fmt.Sprintf("https://%s:10250/healthz", node.Name)
	req, err := http.NewRequest("GET", kubeletEndpoint, nil)
	if err != nil {
		return false, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 50 * time.Millisecond,
	}

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 400 {
		result, _ := ioutil.ReadAll(res.Body)
		klog.V(4).Infof("kubelet healthz API return Not OK, response: %s, node: %s", string(result), node.Name)
		return false, nil
	}

	return true, nil
}
