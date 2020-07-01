package cmd

import "github.com/giantswarm/awscnfm/pkg/key"

var CommandBase = key.GeneratedWithPrefix("command.go")

var CommandContent = `package {{ .Action }}

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/cmd/{{ .Cluster }}/{{ .Action }}/execute"
	"github.com/giantswarm/awscnfm/cmd/{{ .Cluster }}/{{ .Action }}/explain"
)

const (
	name        = "{{ .Action }}"
	description = "Action {{ .Action }} for cluster 001."
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

	var executeCmd *cobra.Command
	{
		c := execute.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		executeCmd, err = execute.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var explainCmd *cobra.Command
	{
		c := explain.Config{
			Logger: config.Logger,
			Stderr: config.Stderr,
			Stdout: config.Stdout,
		}

		explainCmd, err = explain.New(c)
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

	c.AddCommand(executeCmd)
	c.AddCommand(explainCmd)

	return c, nil
}
`
