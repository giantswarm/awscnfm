package ac000

import (
	"context"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/pkg/plan"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var planExecutor *plan.Executor
	{
		c := plan.ExecutorConfig{
			Commands: e.command.Parent().Parent().Commands(),
			Logger:   e.logger,
			Plan:     Plan,
		}

		planExecutor, err = plan.NewExecutor(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = planExecutor.Execute(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
