package ac004

import "github.com/giantswarm/microerror"

var wrongClusterStatusConditionError = &microerror.Error{
	Kind: "wrongClusterStatusConditionError",
	Desc: "We want to see the 'Updated' cluster status condition in order to verify that the Tenant Cluster upgrade was finished successfully.",
}

// IswrongClusterStatusCondition asserts wrongClusterStatusConditionError.
func IswrongClusterStatusCondition(err error) bool {
	return microerror.Cause(err) == wrongClusterStatusConditionError
}
