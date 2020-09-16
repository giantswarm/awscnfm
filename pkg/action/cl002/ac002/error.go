package ac002

import "github.com/giantswarm/microerror"

var wrongClusterStatusConditionError = &microerror.Error{
	Kind: "wrongClusterStatusConditionError",
	Desc: "We want to see the 'Created' cluster status condition in order to verify that the Tenant Cluster creation was finished successfully.",
}

// IswrongClusterStatusCondition asserts wrongClusterStatusConditionError.
func IswrongClusterStatusCondition(err error) bool {
	return microerror.Cause(err) == wrongClusterStatusConditionError
}
