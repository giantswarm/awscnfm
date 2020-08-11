package ac009

import "github.com/giantswarm/microerror"

var customResourceCleanupError = &microerror.Error{
	Kind: "customResourceCleanupError",
	Desc: "We do not expect any CR to be found anymore since we want to ensure that cluster deletion cleans up properly.",
}

// IsCustomResourceCleanup asserts customResourceCleanupError.
func IsCustomResourceCleanup(err error) bool {
	return microerror.Cause(err) == customResourceCleanupError
}
