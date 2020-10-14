package cl004

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac000"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac001"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac002"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac003"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac004"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac005"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac006"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac007"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac008"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac009"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac010"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac011"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac012"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004/ac013"
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

	var ac001Cmd *cobra.Command
	{
		c := ac001.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac001Cmd, err = ac001.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var ac002Cmd *cobra.Command
	{
		c := ac002.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac002Cmd, err = ac002.New(c)
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

	var ac005Cmd *cobra.Command
	{
		c := ac005.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		ac005Cmd, err = ac005.New(c)
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
	c.AddCommand(ac001Cmd)
	c.AddCommand(ac002Cmd)
	c.AddCommand(ac003Cmd)
	c.AddCommand(ac004Cmd)
	c.AddCommand(ac005Cmd)
	c.AddCommand(ac006Cmd)
	c.AddCommand(ac007Cmd)
	c.AddCommand(ac008Cmd)
	c.AddCommand(ac009Cmd)
	c.AddCommand(ac010Cmd)
	c.AddCommand(ac011Cmd)
	c.AddCommand(ac012Cmd)
	c.AddCommand(ac013Cmd)
	c.AddCommand(ac014Cmd)

	return c, nil
}