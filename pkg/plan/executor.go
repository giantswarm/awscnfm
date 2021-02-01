package plan

import (
	"fmt"
	"time"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

type ExecutorConfig struct {
	Commands []*cobra.Command
	Logger   micrologger.Logger

	Plan          []Step
	TenantCluster string
}

type Executor struct {
	commands []*cobra.Command
	logger   micrologger.Logger

	plan          []Step
	tenantCluster string
}

func NewExecutor(config ExecutorConfig) (*Executor, error) {
	if len(config.Commands) == 0 {
		return nil, microerror.Maskf(invalidConfigError, "%T.Commands must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.Plan == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Plan must not be empty", config)
	}
	if config.TenantCluster == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}

	e := &Executor{
		commands: config.Commands,
		logger:   config.Logger,

		plan:          config.Plan,
		tenantCluster: config.TenantCluster,
	}

	return e, nil
}

func (e *Executor) Execute(ctx context.Context) error {
	var err error
	var failed bool
	var cleanup bool

	for _, p := range e.plan {
		if p.Condition == ConditionAlwaysExecute {
			cleanup = true
		}
	}

	for _, p := range e.plan {
		e.logger.LogCtx(ctx, "level", "info", "message", fmt.Sprintf("executing action %#q", p.Action))

		o := func() error {
			cmds, err := commandsForAction("action", e.commands)
			if err != nil {
				return microerror.Mask(err)
			}

			c, err := commandForAction(p.Action, cmds)
			if err != nil {
				return microerror.Mask(err)
			}

			// Here we inject the generated cluster ID to each command. That
			// mechanism ensures that all commands throughout plan exection operate
			// on the same cluster from cluster creation to cluster deletion. That
			// implies that all commands must provide the same flag to allow the
			// propagation of the cluster ID.
			f := c.LocalNonPersistentFlags()
			err = f.Set("tenant-cluster", e.tenantCluster)
			if err != nil {
				return microerror.Mask(err)
			}

			err = c.RunE(c, nil)
			if err != nil {
				return microerror.Mask(err)
			}

			return nil
		}

		if cleanup {
			if !failed || failed && p.Condition == ConditionAlwaysExecute {
				err = backoff.Retry(o, p.Backoff.Wrapped())
				if err != nil {
					failed = true
					e.logger.LogCtx(ctx, "level", "error", "message", fmt.Sprintf("failed executing action %#q", p.Action))
				}
			}
		} else {
			err = backoff.Retry(o, p.Backoff.Wrapped())
			if err != nil {
				return microerror.Mask(err)
			}
		}

		if !failed && p.CoolDown != 0 {
			e.logger.LogCtx(ctx, "level", "info", "message", fmt.Sprintf("cooling down for %s", p.CoolDown))
			time.Sleep(p.CoolDown)
		}
	}

	return nil
}

func (e *Executor) Validate() error {
	cmds, err := commandsForAction("action", e.commands)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, p := range e.plan {
		_, err := commandForAction(p.Action, cmds)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

func commandForAction(action StepAction, commands []*cobra.Command) (*cobra.Command, error) {
	var cmd *cobra.Command

	cmds := append([]*cobra.Command{}, commands...)
	acts := action.Split()

	for _, a := range acts {
		for _, c := range cmds {
			if c.Name() == a {
				cmd = c
				cmds = c.Commands()
				break
			}
		}
	}

	if cmd == nil {
		return nil, microerror.Maskf(commandNotFoundError, string(action))
	}

	return cmd, nil
}

func commandsForAction(action string, commands []*cobra.Command) ([]*cobra.Command, error) {
	var cmds []*cobra.Command
	{
		for _, c := range commands {
			if c.Name() == action {
				cmds = c.Commands()
				break
			}
		}

		if cmds == nil {
			return nil, microerror.Maskf(commandNotFoundError, action)
		}
	}

	return cmds, nil
}
