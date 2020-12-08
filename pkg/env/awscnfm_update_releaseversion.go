package env

import (
	"os"
)

const (
	EnvVarUpdateReleaseVersion = "AWSCNFM_UPDATE_RELEASEVERSION"
)

// UpdateReleaseVersion is the release version used to update the Tenant Cluster
// we want to test for conformity.
func UpdateReleaseVersion() string {
	return os.Getenv(EnvVarUpdateReleaseVersion)
}
