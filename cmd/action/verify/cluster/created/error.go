package created

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

var wrongClusterStatusConditionError = &microerror.Error{
	Kind: "wrongClusterStatusConditionError",
	Desc: "We want to see the 'Created' cluster status condition in order to verify that the Tenant Cluster creation was finished successfully.",
}

// IswrongClusterStatusCondition asserts wrongClusterStatusConditionError.
func IswrongClusterStatusCondition(err error) bool {
	return microerror.Cause(err) == wrongClusterStatusConditionError
}
