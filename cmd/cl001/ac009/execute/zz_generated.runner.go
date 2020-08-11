package execute

import (
	"context"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/pkg/action"
	"github.com/giantswarm/awscnfm/pkg/action/cl001/ac009"
	"github.com/giantswarm/awscnfm/pkg/config"
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
			TenantCluster: config.Cluster("cl001", env.TenantCluster()),
		}

		clients, err = action.NewClients(c)
		if err != nil {
			return microerror.Mask(err)
		}

		err = clients.InitControlPlane(ctx)
		if err != nil {
			return microerror.Mask(err)
		}

		err = clients.InitTenantCluster(ctx)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var e action.Executor
	{
		c := ac009.ExecutorConfig{
			Clients: clients,
			Command: cmd,
			Logger:  r.logger,

			TenantCluster: config.Cluster("cl001", env.TenantCluster()),
		}

		e, err = ac009.NewExecutor(c)
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
