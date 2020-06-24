package completion

import (
	"io"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/cmd/completion/zsh"
)

const (
	name        = "completion"
	description = "Generate shell completions for zsh."
)

type Config struct {
	Logger  micrologger.Logger
	MainCmd *cobra.Command
	Stderr  io.Writer
	Stdout  io.Writer
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

	var zshCmd *cobra.Command
	{
		c := zsh.Config{
			Logger:  config.Logger,
			MainCmd: config.MainCmd,
			Stderr:  config.Stderr,
			Stdout:  config.Stdout,
		}

		zshCmd, err = zsh.New(c)
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

	c.AddCommand(zshCmd)

	return c, nil
}
