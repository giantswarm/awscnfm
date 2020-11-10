package deleted

import "github.com/giantswarm/microerror"

var customResourceCleanupError = &microerror.Error{
	Kind: "customResourceCleanupError",
	Desc: "We do not expect any CR to be found anymore since we want to ensure that cluster deletion cleans up properly.",
}

// IsCustomResourceCleanup asserts customResourceCleanupError.
func IsCustomResourceCleanup(err error) bool {
	return microerror.Cause(err) == customResourceCleanupError
}

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
