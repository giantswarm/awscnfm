package ac007

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Check if the number of host network pods on tenant cluster worker nodes matches the number we expect from k8scloudconfig.

	* Fetch all Tenant Cluster nodes and take the the first worker node by label.
	* Compare the current pods with host network set with the expected amount of pods on worker node.
	* See also https://github.com/giantswarm/k8scloudconfig/blob/529491d591e039da1ffde03fef070101c8d4a95c/files/conf/setup-kubelet-environment#L21. 
	`

	return s, nil
}
