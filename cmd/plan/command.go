package plan

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/plan/pl001"
	"github.com/giantswarm/awscnfm/v12/cmd/plan/pl002"
	"github.com/giantswarm/awscnfm/v12/cmd/plan/pl006"
)

const (
	name        = "plan"
	description = "Execute test plans automatically."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var pl001Cmd *cobra.Command
	{
		c := pl001.Config{
			Logger: config.Logger,
		}

		pl001Cmd, err = pl001.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var pl002Cmd *cobra.Command
	{
		c := pl002.Config{
			Logger: config.Logger,
		}

		pl002Cmd, err = pl002.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var pl006Cmd *cobra.Command
	{
		c := pl006.Config{
			Logger: config.Logger,
		}

		pl006Cmd, err = pl006.New(c)
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

	c.AddCommand(pl001Cmd)
	c.AddCommand(pl002Cmd)
	c.AddCommand(pl006Cmd)

	return c, nil
}
