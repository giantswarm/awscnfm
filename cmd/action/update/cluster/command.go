package cluster

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/action/update/cluster/hacontrolplane"
	"github.com/giantswarm/awscnfm/v12/cmd/action/update/cluster/version"
)

const (
	name        = "cluster"
	description = "Update Tenant Clusters within a Control Plane."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var versionCmd *cobra.Command
	{
		c := version.Config{
			Logger: config.Logger,
		}

		versionCmd, err = version.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var haCmd *cobra.Command
	{
		c := hacontrolplane.Config{
			Logger: config.Logger,
		}

		haCmd, err = hacontrolplane.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
	}

	c := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
		RunE:  r.Run,
	}

	f.Init(c)

	c.AddCommand(haCmd)
	c.AddCommand(versionCmd)

	return c, nil
}
