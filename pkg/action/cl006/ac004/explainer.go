package ac004

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Check if the Tenant Cluster got successfully upgraded to HA. 

	* Check if the all 3 master ready are in ready state.
	* Return an error if we see other than 3 master nodes.

	`

	return s, nil
}
