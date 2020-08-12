package ac008

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/action"
)

type ExecutorConfig struct {
	Clients *action.Clients
	Command *cobra.Command
	Logger  micrologger.Logger

	TenantCluster string
}

type Executor struct {
	clients *action.Clients
	command *cobra.Command
	logger  micrologger.Logger

	tenantCluster string
}

func NewExecutor(config ExecutorConfig) (*Executor, error) {
	if config.Clients == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Clients must not be empty", config)
	}
	if config.Command == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Command must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.TenantCluster == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}

	e := &Executor{
		clients: config.Clients,
		command: config.Command,
		logger:  config.Logger,

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
