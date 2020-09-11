package ac004

import (
	"context"
)

func (e *Explainer) explain(ctx context.Context) (string, error) {
	s := `
Check if the Tenant Cluster got successfully upgraded. Note that this
particular action is not meant to be reliably used for different purposes
than for the plan exection of cluster scope cl002. Executing this action
against a Tenant Cluster that got already upgraded may lead to wrong results
in case you want to assert an additional Tenant Cluster upgrade.

	* Fetch the AWSCluster CR.
	* Check if the latest cluster status condition is "Updated".
	* Return an error if we see other cluster status conditions than "Updated".

	`

	return s, nil
}
