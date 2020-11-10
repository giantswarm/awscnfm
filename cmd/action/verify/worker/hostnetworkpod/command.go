package hostnetworkpod

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "hostnetworkpod"
	short = "Verify the number of host network pods on worker nodes."
	long  = `Check if the number of host network pods on tenant cluster worker nodes
matches the number we expect from k8scloudconfig.

	* Fetch all Tenant Cluster nodes and take the the first worker node by label.
	* Compare the current pods with host network set with the expected amount of pods on worker node.
	* See also https://github.com/giantswarm/k8scloudconfig/blob/529491d591e039da1ffde03fef070101c8d4a95c/files/conf/setup-kubelet-environment#L26.
`
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
	}

	c := &cobra.Command{
		Use:   name,
		Short: short,
		Long:  long,
		RunE:  r.Run,
	}

	f.Init(c)

	return c, nil
}
