package action

import (
	"github.com/giantswarm/awscnfm/v12/pkg/key"
)

var ExecutorBase = key.GeneratedWithPrefix("executor.go")

var ExecutorContent = `package {{ .Action }}

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/action"{{ if eq .Action "ac001" }}
	"github.com/giantswarm/awscnfm/v12/pkg/config"{{ end }}
)

type ExecutorConfig struct {
	Clients *action.Clients
	Command *cobra.Command
	Logger  micrologger.Logger{{ if and (ne .Action "ac000") (ne .Action "ac001") }}

	TenantCluster string{{ end }}
}

type Executor struct {
	clients *action.Clients
	command *cobra.Command
	logger  micrologger.Logger{{ if and (ne .Action "ac000") (ne .Action "ac001") }}

	tenantCluster string{{ end }}
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
	}{{ if and (ne .Action "ac000") (ne .Action "ac001") }}

	if config.TenantCluster == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}{{ end }}

	e := &Executor{
		clients: config.Clients,
		command: config.Command,
		logger:  config.Logger,{{ if and (ne .Action "ac000") (ne .Action "ac001") }}

		tenantCluster: config.TenantCluster,{{ end }}
	}

	return e, nil
}

func (e *Executor) Execute(ctx context.Context) error {
	{{ if eq .Action "ac001" }}crs, err := e.execute(ctx){{ else }}err := e.execute(ctx){{ end }}
	if err != nil {
		return microerror.Mask(err)
	}{{ if eq .Action "ac001" }}

	config.SetCluster("{{ .Cluster }}", crs.Cluster.GetName()){{ end }}

	return nil
}
`
