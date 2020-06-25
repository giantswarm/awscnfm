package ac001

func Explain() string {
	return `
	Check if the desired amount of Tenant Cluster master nodes are up and ready.

		* Fetch all G8sControlPlane CRs spec.replicas so that we know how many masters the Tenant Cluster is supposed to have.
		* Fetch the Tenant Cluster master nodes.
		* Compare the current and desired amount of master nodes.

	Check if the desired amount of Tenant Cluster worker nodes are up and ready.

		* Fetch all AWSMachineDeployment CRs spec.scaling.min so that we know how many workers the Tenant Cluster is supposed to have.
		* Fetch the Tenant Cluster worker nodes.
		* Compare the current and desired amount of worker nodes.
	`
}
