package running

import "github.com/giantswarm/microerror"

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidFlagsError = &microerror.Error{
	Kind: "invalidFlagsError",
}

// IsInvalidFlags asserts invalidFlagsError.
func IsInvalidFlags(err error) bool {
	return microerror.Cause(err) == invalidFlagsError
}

// appsNotRunning is an error indicating that some apps are not running.
var appsNotRunning = &microerror.Error{
	Kind: "appsNotRunning",
	Desc: "Some applications are not running",
}

// IsAppsNotRunning asserts appsNotRunning error.
func IsAppsNotInstalled(err error) bool {
	return microerror.Cause(err) == appsNotRunning
}
