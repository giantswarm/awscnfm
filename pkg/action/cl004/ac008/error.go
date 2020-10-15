package ac008

import "github.com/giantswarm/microerror"

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
