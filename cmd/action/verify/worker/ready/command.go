package ready

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "ready"
	short = "Verify if all worker nodes are ready."
	long  = `
Check if the desired amount of Tenant Cluster worker nodes are up and ready.

	* Fetch all AWSMachineDeployment CRs spec.scaling.min so that we know how many workers the Tenant Cluster is supposed to have.
	* Fetch the Tenant Cluster worker nodes.
	* Compare the current and desired amount of worker nodes.
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
