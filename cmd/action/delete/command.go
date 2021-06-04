package delete

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/cmd/action/delete/cluster"
	"github.com/giantswarm/awscnfm/v15/cmd/action/delete/ebs"
	"github.com/giantswarm/awscnfm/v15/cmd/action/delete/kiam"
	"github.com/giantswarm/awscnfm/v15/cmd/action/delete/netpol"
	"github.com/giantswarm/awscnfm/v15/cmd/action/delete/networkpool"
)

const (
	name        = "delete"
	description = "Delete resources of conformance tests, e.g. Tenant Cluster CRs."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var clusterCmd *cobra.Command
	{
		c := cluster.Config{
			Logger: config.Logger,
		}

		clusterCmd, err = cluster.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ebsCmd *cobra.Command
	{
		c := ebs.Config{
			Logger: config.Logger,
		}

		ebsCmd, err = ebs.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var kiamCmd *cobra.Command
	{
		c := kiam.Config{
			Logger: config.Logger,
		}

		kiamCmd, err = kiam.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var netpolCmd *cobra.Command
	{
		c := netpol.Config{
			Logger: config.Logger,
		}

		netpolCmd, err = netpol.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var networkpoolCmd *cobra.Command
	{
		c := networkpool.Config{
			Logger: config.Logger,
		}

		networkpoolCmd, err = networkpool.New(c)
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

	c.AddCommand(clusterCmd)
	c.AddCommand(ebsCmd)
	c.AddCommand(kiamCmd)
	c.AddCommand(netpolCmd)
	c.AddCommand(networkpoolCmd)

	return c, nil
}
