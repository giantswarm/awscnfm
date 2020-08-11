package ac001

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/pkg/action"
	"github.com/giantswarm/awscnfm/pkg/config"
)

type ExecutorConfig struct {
	Clients *action.Clients
	Command *cobra.Command
	Logger  micrologger.Logger
}

type Executor struct {
	clients *action.Clients
	command *cobra.Command
	logger  micrologger.Logger
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

	e := &Executor{
		clients: config.Clients,
		command: config.Command,
		logger:  config.Logger,
	}

	return e, nil
}

func (e *Executor) Execute(ctx context.Context) error {
	crs, err := e.execute(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	config.SetCluster("cl002", crs.Cluster.GetName())

	return nil
}
