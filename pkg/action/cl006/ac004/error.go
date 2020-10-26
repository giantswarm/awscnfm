package ac004

import "github.com/giantswarm/microerror"

var wrongMasterNodesError = &microerror.Error{
	Kind: "wrongMasterNodesError",
	Desc: "We want to see the 3 master nodes to verify that the Tenant Cluster HA upgrade was finished successfully.",
}

// IswrongMasterNodesError asserts wrongMasterNodesError.
func IswrongWrongMasterNodes(err error) bool {
	return microerror.Cause(err) == wrongMasterNodesError
}
