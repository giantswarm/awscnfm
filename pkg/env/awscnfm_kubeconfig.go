package env

import (
	"os"
)

const (
	EnvVarKubeConfig = "AWSCNFM_KUBECONFIG"
)

// KubeConfig is the local path to the kube config file used to create the
// Control Plane specific rest config.
func KubeConfig() string {
	return os.Getenv(EnvVarKubeConfig)
}
