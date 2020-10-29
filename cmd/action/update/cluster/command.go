package cluster

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/action/update/cluster/major"
	"github.com/giantswarm/awscnfm/v12/cmd/action/update/cluster/minor"
	"github.com/giantswarm/awscnfm/v12/cmd/action/update/cluster/patch"
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

	var majorCmd *cobra.Command
	{
		c := major.Config{
			Logger: config.Logger,
		}

		majorCmd, err = major.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var minorCmd *cobra.Command
	{
		c := minor.Config{
			Logger: config.Logger,
		}

		minorCmd, err = minor.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var patchCmd *cobra.Command
	{
		c := patch.Config{
			Logger: config.Logger,
		}

		patchCmd, err = patch.New(c)
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

	c.AddCommand(majorCmd)
	c.AddCommand(minorCmd)
	c.AddCommand(patchCmd)

	return c, nil
}
