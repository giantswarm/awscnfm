package ac008

import "github.com/giantswarm/microerror"

var notFoundError = &microerror.Error{
	Kind: "notFoundError",
}

// IsNotFound asserts notFoundError.
func IsNotFound(err error) bool {
	return microerror.Cause(err) == notFoundError
}

var tooManyCRsError = &microerror.Error{
	Kind: "tooManyCRsError",
	Desc: "There is only a single AWSCluster CR allowed with the current implementation.",
}

// IsTooManyCRsError asserts tooManyCRsError.
func IsTooManyCRsError(err error) bool {
	return microerror.Cause(err) == tooManyCRsError
}
