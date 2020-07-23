package ac001

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/awscnfm/pkg/action"
)

const (
	// executorCommand is for internal documentation purposes only so that
	// commands can self describe and explain themselves better. This
	// information might be used in different creative ways.
	executorCommand = "awscnfm cl001 ac001 execute"
)

type ExecutorConfig struct {
	Clients *action.Clients
	Logger  micrologger.Logger

	TenantCluster string
}

type Executor struct {
	clients *action.Clients
	logger  micrologger.Logger

	tenantCluster string
}

func NewExecutor(config ExecutorConfig) (*Executor, error) {
	if config.Clients == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Clients must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.TenantCluster == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}

	e := &Executor{
		clients: config.Clients,
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
