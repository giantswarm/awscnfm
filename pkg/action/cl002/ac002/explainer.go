package ac002

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Check if the number of pods on master and worker node with host network set matches the number we expect from k8scloudconfig.

	* Fetch all Tenant Cluster nodes and take the the first worker node and master node via label.
	* Compare the current pods with host network set with the expected amount of pods on master node.
	* Compare the current pods with host network set with the expected amount of pods on worker node.
	`

	return s, nil
}
