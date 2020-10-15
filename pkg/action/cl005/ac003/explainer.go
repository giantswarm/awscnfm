package ac003

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Upgrade the Tenant Cluster to the latest major version.

	* Fetch the Cluster CR.
	* Set the desired cluster-operator version in the CR labels.
	* Set the desired release version in the CR labels.
	* Update the Cluster CR in the Control Plane.

	`

	return s, nil
}
