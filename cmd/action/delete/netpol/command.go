package netpol

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/action/delete/netpol/curlrequest"
	"github.com/giantswarm/awscnfm/v12/cmd/action/delete/netpol/defaultnetpol"
)

const (
	name        = "netpol"
	description = "Delete certain netpol test resources."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var curlrequestCmd *cobra.Command
	{
		c := curlrequest.Config{
			Logger: config.Logger,
		}

		curlrequestCmd, err = curlrequest.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var defaultnetpolCmd *cobra.Command
	{
		c := defaultnetpol.Config{
			Logger: config.Logger,
		}

		defaultnetpolCmd, err = defaultnetpol.New(c)
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

	c.AddCommand(curlrequestCmd)
	c.AddCommand(defaultnetpolCmd)

	return c, nil
}
