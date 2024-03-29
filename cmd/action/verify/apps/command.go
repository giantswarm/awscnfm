package apps

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/apps/installed"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify/apps/running"
)

const (
	name        = "apps"
	description = "Verify apps are installed and running correctly in the Tenant Cluster."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var installedCmd *cobra.Command
	{
		c := installed.Config{
			Logger: config.Logger,
		}

		installedCmd, err = installed.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var runningCmd *cobra.Command
	{
		c := running.Config{
			Logger: config.Logger,
		}

		runningCmd, err = running.New(c)
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

	c.AddCommand(installedCmd)
	c.AddCommand(runningCmd)

	return c, nil
}
