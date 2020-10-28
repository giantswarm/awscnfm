package cluster

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "cluster"
	short = "Delete all Tenant Cluster CRs on the Control Plane."
	long  = `Delete all Tenant Cluster CRs on the Control Plane by triggering the
deletion of the Cluster CR. This should ensure the following.

	* Trigger deletion to all other CRs associated with the tenant cluster.
	* Execute cleanup logic in all involved operators.
	* Remove all cloud provider resources.
	* Remove all CRs associated with the tenant cluster.

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
