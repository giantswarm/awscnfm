package plan

import "github.com/giantswarm/microerror"

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var commandNotFoundError = &microerror.Error{
	Kind: "commandNotFoundError",
	Desc: "There is no cobra command registered for the action defined by the current conformance test plan.",
}

// IsCommandNotFound asserts commandNotFoundError.
func IsCommandNotFound(err error) bool {
	return microerror.Cause(err) == commandNotFoundError
}
