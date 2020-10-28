package deleted

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "created"
	short = "Check that all relevant CRs of the Tenant Cluster got properly cleaned up."
	long  = `Check that all relevant CRs of the tenant cluster got properly cleaned up
eventually during the transition of cluster deletion. This check considers
the following CRs.

	* Cluster
	* AWSCluster
	* G8sControlPlane
	* AWSControlPlane
	* MachineDeployment
	* AWSMachineDeployment
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
