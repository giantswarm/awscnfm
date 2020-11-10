package valid

import "github.com/giantswarm/microerror"

var invalidIDError = &microerror.Error{
	Kind: "invalidIDError",
}

// IsInvalidID asserts invalidIDError.
func IsInvalidID(err error) bool {
	return microerror.Cause(err) == invalidIDError
}
