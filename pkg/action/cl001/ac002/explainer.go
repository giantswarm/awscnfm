package ac002

import "context"

type ExplainerConfig struct {
}

type Explainer struct {
}

func NewExplainer(config ExplainerConfig) (*Explainer, error) {
	e := &Explainer{}

	return e, nil
}

func (e *Explainer) Explain(ctx context.Context) (string, error) {
	return `
Check if the desired amount of Tenant Cluster worker nodes are up and ready.

	* Fetch all AWSMachineDeployment CRs spec.scaling.min so that we know how many workers the Tenant Cluster is supposed to have.
	* Fetch the Tenant Cluster worker nodes.
	* Compare the current and desired amount of worker nodes.
	`, nil
}
