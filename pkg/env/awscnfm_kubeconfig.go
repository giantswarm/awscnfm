package env

import (
	"os"
)

const (
	EnvVarControlPlaneKubeConfig = "AWSCNFM_CONTROLPLANE_KUBECONFIG"
)

const (
	DefaultControlPlaneKubeConfig = "~/.kube/config"
)

// ControlPlaneKubeConfig is the local path to the kube config file used to
// create the Control Plane specific rest config.
func ControlPlaneKubeConfig() string {
	return os.Getenv(EnvVarControlPlaneKubeConfig)
}
