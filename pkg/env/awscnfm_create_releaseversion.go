package env

import (
	"os"
)

const (
	EnvVarCreateReleaseVersion = "AWSCNFM_CREATE_RELEASEVERSION"
)

// CreateReleaseVersion is the release version used to create the Tenant Cluster
// we want to test for conformity.
func CreateReleaseVersion() string {
	return os.Getenv(EnvVarCreateReleaseVersion)
}
