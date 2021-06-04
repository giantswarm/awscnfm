package nodepool

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/cmd/action/create/nodepool/defaultdataplane"
)

const (
	name        = "nodepool"
	description = "Create Node Pools for a Tenant Cluster."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var defaultdataplaneCmd *cobra.Command
	{
		c := defaultdataplane.Config{
			Logger: config.Logger,
		}

		defaultdataplaneCmd, err = defaultdataplane.New(c)
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

	c.AddCommand(defaultdataplaneCmd)

	return c, nil
}
