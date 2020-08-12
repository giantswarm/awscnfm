package ac006

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Check if the number of host network pods on tenant cluster master nodes matches the number we expect from k8scloudconfig.

	* Fetch all Tenant Cluster nodes and take the the first master node by label.
	* Compare the current pods with host network set with the expected amount of pods on master node.
	`

	return s, nil
}
