package ac005

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
)

type ExecutorConfig struct {
	Command *cobra.Command
	Logger  micrologger.Logger

	Scope         string
	TenantCluster string
}

type Executor struct {
	command *cobra.Command
	logger  micrologger.Logger

	scope         string
	tenantCluster string
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

		scope:         config.Scope,
		tenantCluster: config.TenantCluster,
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
