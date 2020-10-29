package created

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "created"
	short = "Verify if a Tenant Cluster got successfully created."
	long  = `Check if the Tenant Cluster successfully created.

    * List all Tenant Cluster nodes. Doing so without errors means the apiserver is up.
    * Check for the "Created" status condition in the AWSCluster CR.

A cluster creation takes up to 30 minutes. This aligns with our cluster creation metric in cluster-operator,
see https://github.com/giantswarm/cluster-operator/blob/master/service/collector/cluster_transition.go#L135.

More information about cluster transitions: https://intranet.giantswarm.io/docs/monitoring/metrics/cluster-transitions/.
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
