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
	"github.com/spf13/cobra"{{ if eq .Action "ac001" }}

	"github.com/giantswarm/awscnfm/v12/pkg/config"{{ end }}
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
	{{ if eq .Action "ac001" }}crs, err := e.execute(ctx){{ else }}err := e.execute(ctx){{ end }}
	if err != nil {
		return microerror.Mask(err)
	}{{ if eq .Action "ac001" }}

	config.SetCluster("{{ .Cluster }}", crs.Cluster.GetName()){{ end }}

	return nil
}
`
