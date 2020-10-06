package env

import (
	"os"
)

const (
	EnvVarTenantClusterKubeConfig = "AWSCNFM_TENANTCLUSTER_KUBECONFIG"
)

// TenantClusterKubeConfig is the local path to the kube config file used to
// create the Control Plane specific rest config.
func TenantClusterKubeConfig() string {
	return os.Getenv(EnvVarTenantClusterKubeConfig)
}
