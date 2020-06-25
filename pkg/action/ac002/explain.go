package ac002

func Explain() string {
	return `
Check if the desired amount of Tenant Cluster worker nodes are up and ready.

	* Fetch all AWSMachineDeployment CRs spec.scaling.min so that we know how many workers the Tenant Cluster is supposed to have.
	* Fetch the Tenant Cluster worker nodes.
	* Compare the current and desired amount of worker nodes.
	`
}
