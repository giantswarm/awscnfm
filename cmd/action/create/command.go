package create

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/action/create/cluster"
	"github.com/giantswarm/awscnfm/v12/cmd/action/create/kiam"
	"github.com/giantswarm/awscnfm/v12/cmd/action/create/nodepool"
)

const (
	name        = "create"
	description = "Create resources for conformance tests, e.g. Tenant Cluster CRs."
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

	var nodepoolCmd *cobra.Command
	{
		c := nodepool.Config{
			Logger: config.Logger,
		}

		nodepoolCmd, err = nodepool.New(c)
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
	c.AddCommand(kiamCmd)
	c.AddCommand(nodepoolCmd)

	return c, nil
}
