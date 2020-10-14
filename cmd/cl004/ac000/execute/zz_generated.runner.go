package execute

import (
	"context"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/action"
	"github.com/giantswarm/awscnfm/v12/pkg/action/cl004/ac000"
	"github.com/giantswarm/awscnfm/v12/pkg/config"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
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

	var e action.Executor
	{
		c := ac000.ExecutorConfig{
			Command: cmd,
			Logger:  r.logger,

			Scope:         "cl004",
			TenantCluster: config.Cluster("cl004", env.TenantCluster()),
		}

		e, err = ac000.NewExecutor(c)
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
