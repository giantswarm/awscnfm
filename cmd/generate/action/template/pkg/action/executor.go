package action

import (
	"github.com/giantswarm/awscnfm/pkg/key"
)

var ExecutorBase = key.GeneratedWithPrefix("executor.go")

var ExecutorContent = `package {{ .Action }}

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/awscnfm/pkg/action"{{ if eq .Action "ac001" }}
	"github.com/giantswarm/awscnfm/pkg/config"{{ end }}
)

type ExecutorConfig struct {
	Clients *action.Clients
	Logger  micrologger.Logger{{ if ne .Action "ac001" }}

	TenantCluster string{{ end }}
}

type Executor struct {
	clients *action.Clients
	logger  micrologger.Logger{{ if ne .Action "ac001" }}

	tenantCluster string{{ end }}
}

func NewExecutor(config ExecutorConfig) (*Executor, error) {
	if config.Clients == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Clients must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}{{ if ne .Action "ac001" }}

	if config.TenantCluster == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}{{ end }}

	e := &Executor{
		clients: config.Clients,
		logger:  config.Logger,{{ if ne .Action "ac001" }}

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
