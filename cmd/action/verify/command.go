package verify

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/apps"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/cluster"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/ebs"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/kiam"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/master"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/netpol"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/worker"
)

const (
	name        = "verify"
	description = "Verify state within conformance tests, e.g. if a Tenant Cluster is created."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var appsCmd *cobra.Command
	{
		c := apps.Config{
			Logger: config.Logger,
		}

		appsCmd, err = apps.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

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

	var masterCmd *cobra.Command
	{
		c := master.Config{
			Logger: config.Logger,
		}

		masterCmd, err = master.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var netpolCommand *cobra.Command
	{
		c := netpol.Config{
			Logger: config.Logger,
		}

		netpolCommand, err = netpol.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var workerCmd *cobra.Command
	{
		c := worker.Config{
			Logger: config.Logger,
		}

		workerCmd, err = worker.New(c)
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

	c.AddCommand(appsCmd)
	c.AddCommand(clusterCmd)
	c.AddCommand(ebsCmd)
	c.AddCommand(kiamCmd)
	c.AddCommand(masterCmd)
	c.AddCommand(netpolCommand)
	c.AddCommand(workerCmd)

	return c, nil
}
