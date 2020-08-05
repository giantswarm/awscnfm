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
	`

	return s, nil
}
