package cluster

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/action/verify/cluster/created"
	"github.com/giantswarm/awscnfm/v12/cmd/action/verify/cluster/deleted"
	"github.com/giantswarm/awscnfm/v12/cmd/action/verify/cluster/hasetup"
)

const (
	name        = "cluster"
	description = "Verify certain Tenant Cluster aspects within a Control Plane."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var createdCmd *cobra.Command
	{
		c := created.Config{
			Logger: config.Logger,
		}

		createdCmd, err = created.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var deletedCmd *cobra.Command
	{
		c := deleted.Config{
			Logger: config.Logger,
		}

		deletedCmd, err = deleted.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var hasetupCmd *cobra.Command
	{
		c := hasetup.Config{
			Logger: config.Logger,
		}

		hasetupCmd, err = hasetup.New(c)
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

	c.AddCommand(createdCmd)
	c.AddCommand(deletedCmd)
	c.AddCommand(hasetupCmd)

	return c, nil
}
