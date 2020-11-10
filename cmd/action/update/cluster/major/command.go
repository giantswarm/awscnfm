package major

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

const (
	name  = "major"
	short = "Upgrade the Tenant Cluster to the latest major version."
	long  = `Upgrade the Tenant Cluster to the latest major version.

    * Fetch the Cluster CR.
    * Set the desired cluster-operator version in the CR labels.
    * Set the desired release version in the CR labels.
    * Update the Cluster CR in the Control Plane.

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
