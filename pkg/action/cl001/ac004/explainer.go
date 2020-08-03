package ac004

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Delete the tenant cluster of the current cluster scope by triggering the
deletion of the Cluster CR. This should ensure the following.

	* Trigger deletion to all other CRs associated with the tenant cluster.
	* Execute cleanup logic in all involved operators.
	* Remove all cloud provider resources.
	* Remove all CRs associated with the tenant cluster.

`

	return s, nil
}
