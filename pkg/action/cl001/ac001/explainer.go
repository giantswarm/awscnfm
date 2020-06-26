package ac001

type ExplainerConfig struct {
}

type Explainer struct {
}

func NewExplainer(config ExplainerConfig) (*Explainer, error) {
	e := &Explainer{}

	return e, nil
}

func (e *Explainer) Explain() string {
	return `
Check if the desired amount of Tenant Cluster master nodes are up and ready.

	* Fetch all G8sControlPlane CRs spec.replicas so that we know how many masters the Tenant Cluster is supposed to have.
	* Fetch the Tenant Cluster master nodes.
	* Compare the current and desired amount of master nodes.
	`
}
