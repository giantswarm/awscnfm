package ac013

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Check that all relevant CRs of the tenant cluster management got properly
cleaned up eventually during the transition of cluster deletion. This check
considers the following CRs.

	* Cluster
	* AWSCluster
	* G8sControlPlane
	* AWSControlPlane
	* MachineDeployment
	* AWSMachineDeployment

`

	return s, nil
}
