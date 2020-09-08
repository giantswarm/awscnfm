package ac002

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Check if the Tenant Cluster is up and we can connect to it.

	* List all Tenant Cluster nodes.
	* Listing all Tenant Cluster nodes without errors means the apiserver is up.
	* Proceed with further actions.

A cluster creation takes up to 30 minutes. This aligns with our cluster creation metric in cluster-operator,
see https://github.com/giantswarm/cluster-operator/blob/master/service/collector/cluster_transition.go#L135.

More information about cluster transitions: https://intranet.giantswarm.io/docs/monitoring/metrics/cluster-transitions/.
	`

	return s, nil
}
