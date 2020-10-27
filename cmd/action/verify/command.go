package verify

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/action/verify/cluster"
	"github.com/giantswarm/awscnfm/v12/cmd/action/verify/kiam"
	"github.com/giantswarm/awscnfm/v12/cmd/action/verify/master"
	"github.com/giantswarm/awscnfm/v12/cmd/action/verify/worker"
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

	c.AddCommand(clusterCmd)
	c.AddCommand(kiamCmd)
	c.AddCommand(masterCmd)
	c.AddCommand(workerCmd)

	return c, nil
}
