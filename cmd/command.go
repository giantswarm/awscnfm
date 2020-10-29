package cmd

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/cmd/action"
	"github.com/giantswarm/awscnfm/v12/cmd/cl003"
	"github.com/giantswarm/awscnfm/v12/cmd/cl004"
	"github.com/giantswarm/awscnfm/v12/cmd/completion"
	"github.com/giantswarm/awscnfm/v12/cmd/plan"
	"github.com/giantswarm/awscnfm/v12/cmd/version"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
)

var (
	name        = project.Name()
	description = project.Description()
)

type Config struct {
	Logger micrologger.Logger
	Stderr io.Writer
	Stdout io.Writer

	BinaryName string
	GitCommit  string
	Source     string
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

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	m := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
		RunE:  r.Run,
		// We slience errors because we do not want to see spf13/cobra printing.
		// The errors returned by the commands will be propagated to the main.go
		// anyway, where we have custom error printing for the tool.
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	var err error

	var actionCmd *cobra.Command
	{
		c := action.Config{
			Logger: config.Logger,
		}

		actionCmd, err = action.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var cl003Cmd *cobra.Command
	{
		c := cl003.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		cl003Cmd, err = cl003.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var cl004Cmd *cobra.Command
	{
		c := cl004.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		cl004Cmd, err = cl004.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var completionCmd *cobra.Command
	{
		c := completion.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		completionCmd, err = completion.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var planCmd *cobra.Command
	{
		c := plan.Config{
			Logger: config.Logger,
		}

		planCmd, err = plan.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var versionCmd *cobra.Command
	{
		c := version.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,

			GitCommit: config.GitCommit,
			Source:    config.Source,
		}

		versionCmd, err = version.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	f.Init(m)

	m.AddCommand(actionCmd)
	m.AddCommand(cl003Cmd)
	m.AddCommand(cl004Cmd)
	m.AddCommand(completionCmd)
	m.AddCommand(planCmd)
	m.AddCommand(versionCmd)

	return m, nil
}
