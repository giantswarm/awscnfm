package cluster

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/action/create/cluster/defaultcontrolplane"
)

const (
	name        = "cluster"
	description = "Create Tenant Clusters within a Control Plane."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var defaultcontrolplaneCmd *cobra.Command
	{
		c := defaultcontrolplane.Config{
			Logger: config.Logger,
		}

		defaultcontrolplaneCmd, err = defaultcontrolplane.New(c)
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

	c.AddCommand(defaultcontrolplaneCmd)

	return c, nil
}
