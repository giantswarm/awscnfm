package worker

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/worker/hostnetworkpod"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/worker/ready"
)

const (
	name        = "worker"
	description = "Verify certain worker node aspects within a Tenant Cluster."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var hostnetworkpodCmd *cobra.Command
	{
		c := hostnetworkpod.Config{
			Logger: config.Logger,
		}

		hostnetworkpodCmd, err = hostnetworkpod.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var readyCmd *cobra.Command
	{
		c := ready.Config{
			Logger: config.Logger,
		}

		readyCmd, err = ready.New(c)
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

	c.AddCommand(hostnetworkpodCmd)
	c.AddCommand(readyCmd)

	return c, nil
}
