package plan

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/plan/pl001"
	"github.com/giantswarm/awscnfm/v12/cmd/plan/pl002"
	"github.com/giantswarm/awscnfm/v12/cmd/plan/pl003"
	"github.com/giantswarm/awscnfm/v12/cmd/plan/pl004"
	"github.com/giantswarm/awscnfm/v12/cmd/plan/pl005"
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

	var pl003Cmd *cobra.Command
	{
		c := pl003.Config{
			Logger: config.Logger,
		}

		pl003Cmd, err = pl003.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var pl004Cmd *cobra.Command
	{
		c := pl004.Config{
			Logger: config.Logger,
		}

		pl004Cmd, err = pl004.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var pl005Cmd *cobra.Command
	{
		c := pl005.Config{
			Logger: config.Logger,
		}

		pl005Cmd, err = pl005.New(c)
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
	c.AddCommand(pl003Cmd)
	c.AddCommand(pl004Cmd)
	c.AddCommand(pl005Cmd)

	return c, nil
}
