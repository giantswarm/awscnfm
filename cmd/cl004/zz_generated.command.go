package cl004

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac000"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac008"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac012"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac014"
)

const (
	name        = "cl004"
	description = "Conformance tests for cluster scope cl004."
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

	var ac008Cmd *cobra.Command
	{
		c := ac008.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac008Cmd, err = ac008.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ac012Cmd *cobra.Command
	{
		c := ac012.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac012Cmd, err = ac012.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ac014Cmd *cobra.Command
	{
		c := ac014.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac014Cmd, err = ac014.New(c)
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
	c.AddCommand(ac008Cmd)
	c.AddCommand(ac012Cmd)
	c.AddCommand(ac014Cmd)

	return c, nil
}
