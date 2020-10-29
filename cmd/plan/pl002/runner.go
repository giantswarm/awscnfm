package pl002

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/plan"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	var err error

	ctx := context.Background()

	err = r.flag.Validate()
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

	var planExecutor *plan.Executor
	{
		c := plan.ExecutorConfig{
			Commands: cmd.Root().Commands(),
			Logger:   r.logger,
			Plan:     Plan,
		}

		planExecutor, err = plan.NewExecutor(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = planExecutor.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = planExecutor.Execute(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
