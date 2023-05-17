package label

import corev1 "k8s.io/api/core/v1"

const (
	// MasterNodeRole label denotes K8s cluster master node role.
	MasterNodeRole = "node-role.kubernetes.io/master"
	// WorkerNodeRole label denotes K8s cluster worker node role.
	WorkerNodeRole = "node-role.kubernetes.io/worker"
)

// MasterNodeRoles labels denote K8s cluster master node role.
var MasterNodeRoles = []string{
	"node-role.kubernetes.io/master",
	"node-role.kubernetes.io/control-plane",
}

func IsMaster(node corev1.Node) bool {
	for _, role := range MasterNodeRoles {
		if _, ok := node.Labels[role]; ok {
			return true
		}
	}

	return false
}
