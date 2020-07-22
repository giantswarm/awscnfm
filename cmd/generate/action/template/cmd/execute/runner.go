package execute

import (
	"path/filepath"

	"github.com/giantswarm/awscnfm/pkg/key"
)

var RunnerBase = filepath.Join("execute", key.GeneratedWithPrefix("runner.go"))

var RunnerContent = `package execute

import (
	"context"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/pkg/action"
	"github.com/giantswarm/awscnfm/pkg/action/{{ .Cluster }}/{{ .Action }}"
	"github.com/giantswarm/awscnfm/pkg/env"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var clients *action.Clients
	{
		c := action.Config{
			Logger: r.logger,

			KubeConfig:    env.KubeConfig(),
			TenantCluster: env.TenantCluster(),
		}

		clients, err = action.NewClients(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var e action.Executor
	{
		c := {{ .Action }}.ExecutorConfig{
			Clients: clients,
			Logger:  r.logger,

			TenantCluster: env.TenantCluster(),
		}

		e, err = {{ .Action }}.NewExecutor(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = e.Execute(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
`