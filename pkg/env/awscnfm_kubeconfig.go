package env

import (
	"os"
)

const (
	EnvVarKubeConfig = "AWSCNFM_KUBECONFIG"
)

const (
	DefaultKubeConfig = "~/.kube/config"
)

// KubeConfig is the local path to the kube config file used to create the
// Control Plane specific rest config.
func KubeConfig() string {
	return os.Getenv(EnvVarKubeConfig)
}
