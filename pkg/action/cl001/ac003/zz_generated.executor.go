package ac003

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

type ExecutorConfig struct {
	Command *cobra.Command
	Logger  micrologger.Logger
}

type Executor struct {
	command *cobra.Command
	logger  micrologger.Logger
}

func NewExecutor(config ExecutorConfig) (*Executor, error) {
	if config.Command == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Command must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	e := &Executor{
		command: config.Command,
		logger:  config.Logger,
	}

	return e, nil
}

func (e *Executor) Execute(ctx context.Context) error {
	err := e.execute(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
