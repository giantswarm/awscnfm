package hasetup

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

var wrongMasterNodesError = &microerror.Error{
	Kind: "wrongMasterNodesError",
	Desc: "We want to see the 3 master nodes to verify that the Tenant Cluster HA upgrade was finished successfully.",
}

// IswrongMasterNodesError asserts wrongMasterNodesError.
func IswrongWrongMasterNodes(err error) bool {
	return microerror.Cause(err) == wrongMasterNodesError
}
