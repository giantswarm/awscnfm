package env

import (
	"os"
)

const (
	EnvVarReleaseVersion = "AWSCNFM_RELEASEVERSION"
)

// ReleaseVersion is the release version used to create the Tenant Cluster we
// want to test for conformity.
func ReleaseVersion() string {
	return os.Getenv(EnvVarReleaseVersion)
}
