package env

import (
	"os"
)

const (
	EnvVarTenantCluster = "AWSCNFM_TENANTCLUSTER"
)

// TenantCluster is the Tenant Cluster ID used to create the Tenant Cluster
// specific rest config.
func TenantCluster() string {
	return os.Getenv(EnvVarTenantCluster)
}
