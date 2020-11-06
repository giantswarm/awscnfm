package installed

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

// appsNotInstalled is an error indicating that all apps are not installed
var appsNotInstalled = &microerror.Error{
	Kind: "appsNotInstalled",
	Desc: "Some applications are not installed",
}

// IsAppsNotInstalled asserts appsNotInstalled error.
func IsAppsNotInstalled(err error) bool {
	return microerror.Cause(err) == appsNotInstalled
}
