package ac003

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Upgrade the Tenant Cluster to HA.

	* Fetch the Cluster CR.
	* Set replicas to 3.
	* Update the Cluster CR in the Control Plane.

	`

	return s, nil
}
