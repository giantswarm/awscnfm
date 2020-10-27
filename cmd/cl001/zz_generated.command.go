package cl001

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/cl001/ac000"
	"github.com/giantswarm/awscnfm/v12/cmd/cl001/ac007"
	"github.com/giantswarm/awscnfm/v12/cmd/cl001/ac008"
	"github.com/giantswarm/awscnfm/v12/cmd/cl001/ac009"
	"github.com/giantswarm/awscnfm/v12/cmd/cl001/ac010"
	"github.com/giantswarm/awscnfm/v12/cmd/cl001/ac011"
	"github.com/giantswarm/awscnfm/v12/cmd/cl001/ac012"
	"github.com/giantswarm/awscnfm/v12/cmd/cl001/ac013"
)

const (
	name        = "cl001"
	description = "Conformance tests for cluster scope cl001."
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

	var ac007Cmd *cobra.Command
	{
		c := ac007.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac007Cmd, err = ac007.New(c)
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

	var ac009Cmd *cobra.Command
	{
		c := ac009.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac009Cmd, err = ac009.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ac010Cmd *cobra.Command
	{
		c := ac010.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac010Cmd, err = ac010.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ac011Cmd *cobra.Command
	{
		c := ac011.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac011Cmd, err = ac011.New(c)
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

	var ac013Cmd *cobra.Command
	{
		c := ac013.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac013Cmd, err = ac013.New(c)
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
	c.AddCommand(ac007Cmd)
	c.AddCommand(ac008Cmd)
	c.AddCommand(ac009Cmd)
	c.AddCommand(ac010Cmd)
	c.AddCommand(ac011Cmd)
	c.AddCommand(ac012Cmd)
	c.AddCommand(ac013Cmd)

	return c, nil
}
