package kiam

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/kiam/awsapicall"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/kiam/podandsecret"
)

const (
	name        = "kiam"
	description = "Verify certain kiam app aspects within a Tenant Cluster."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var awsapicallCmd *cobra.Command
	{
		c := awsapicall.Config{
			Logger: config.Logger,
		}

		awsapicallCmd, err = awsapicall.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var podandsecretCmd *cobra.Command
	{
		c := podandsecret.Config{
			Logger: config.Logger,
		}

		podandsecretCmd, err = podandsecret.New(c)
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

	c.AddCommand(awsapicallCmd)
	c.AddCommand(podandsecretCmd)

	return c, nil
}
