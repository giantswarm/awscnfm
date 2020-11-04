package networkpool

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

var customNetworkPoolMasterUnusedError = &microerror.Error{
	Kind: "customNetworkPoolMasterUnusedError",
	Desc: "Custom NetworkPool is not being used on the master subnet.",
}

var customNetworkPoolWorkerUnusedError = &microerror.Error{
	Kind: "customNetworkPoolWorkerUnusedError",
	Desc: "Custom NetworkPool is not being used on the worker subnet.",
}

func IsCustomNetworkPoolMasterUnused(err error) bool {
	return microerror.Cause(err) == customNetworkPoolMasterUnusedError
}

func IsCustomNetworkPoolWorkerUnused(err error) bool {
	return microerror.Cause(err) == customNetworkPoolWorkerUnusedError
}
