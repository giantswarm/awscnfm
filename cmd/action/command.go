package action

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v15/cmd/action/create"
	"github.com/giantswarm/awscnfm/v15/cmd/action/delete"
	"github.com/giantswarm/awscnfm/v15/cmd/action/update"
	"github.com/giantswarm/awscnfm/v15/cmd/action/verify"
)

const (
	name        = "action"
	description = "Execute actions separately against tenant clusters."
)

type Config struct {
	Logger micrologger.Logger
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var createCmd *cobra.Command
	{
		c := create.Config{
			Logger: config.Logger,
		}

		createCmd, err = create.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var deleteCmd *cobra.Command
	{
		c := delete.Config{
			Logger: config.Logger,
		}

		deleteCmd, err = delete.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var updateCmd *cobra.Command
	{
		c := update.Config{
			Logger: config.Logger,
		}

		updateCmd, err = update.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var verifyCmd *cobra.Command
	{
		c := verify.Config{
			Logger: config.Logger,
		}

		verifyCmd, err = verify.New(c)
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

	c.AddCommand(createCmd)
	c.AddCommand(deleteCmd)
	c.AddCommand(updateCmd)
	c.AddCommand(verifyCmd)

	return c, nil
}
