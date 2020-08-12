package cmd

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/cmd/cl001"
	"github.com/giantswarm/awscnfm/cmd/completion"
	"github.com/giantswarm/awscnfm/cmd/generate"
	"github.com/giantswarm/awscnfm/cmd/insert"
	"github.com/giantswarm/awscnfm/cmd/version"
	"github.com/giantswarm/awscnfm/pkg/project"
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

	var cl001Cmd *cobra.Command
	{
		c := cl001.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		cl001Cmd, err = cl001.New(c)
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

	var generateCmd *cobra.Command
	{
		c := generate.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		generateCmd, err = generate.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var insertCmd *cobra.Command
	{
		c := insert.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		insertCmd, err = insert.New(c)
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

	m.AddCommand(cl001Cmd)
	m.AddCommand(completionCmd)
	m.AddCommand(generateCmd)
	m.AddCommand(insertCmd)
	m.AddCommand(versionCmd)

	return m, nil
}
