package cl005

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/cl005/ac000"
	"github.com/giantswarm/awscnfm/v12/cmd/cl005/ac003"
	"github.com/giantswarm/awscnfm/v12/cmd/cl005/ac004"
	"github.com/giantswarm/awscnfm/v12/cmd/cl005/ac006"
)

const (
	name        = "cl005"
	description = "Conformance tests for cluster scope cl005."
)

type Config struct {
	Logger micrologger.Logger
	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	var err error

	var ac000Cmd *cobra.Command
	{
		c := ac000.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac000Cmd, err = ac000.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ac003Cmd *cobra.Command
	{
		c := ac003.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac003Cmd, err = ac003.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ac004Cmd *cobra.Command
	{
		c := ac004.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac004Cmd, err = ac004.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ac006Cmd *cobra.Command
	{
		c := ac006.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac006Cmd, err = ac006.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
		RunE:  r.Run,
	}

	f.Init(c)

	c.AddCommand(ac000Cmd)
	c.AddCommand(ac003Cmd)
	c.AddCommand(ac004Cmd)
	c.AddCommand(ac006Cmd)

	return c, nil
}
