package plan

import (
	"fmt"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

type ExecutorConfig struct {
	Commands []*cobra.Command
	Logger   micrologger.Logger
	Plan     []Step
}

type Executor struct {
	commands []*cobra.Command
	logger   micrologger.Logger
	plan     []Step
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

	e := &Executor{
		commands: config.Commands,
		logger:   config.Logger,
		plan:     config.Plan,
	}

	return e, nil
}

func (e *Executor) Execute(ctx context.Context) error {
	for _, p := range e.plan {
		c, err := commandForAction(p.Action, e.commands)
		if err != nil {
			return microerror.Mask(err)
		}

		e.logger.LogCtx(ctx, "level", "info", "message", fmt.Sprintf("executing action %#q", p.Action))

		o := func() error {
			err := c.RunE(c, nil)
			if err != nil {
				return microerror.Mask(err)
			}

			return nil
		}

		err = backoff.Retry(o, p.Backoff.Wrapped())
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

func commandForAction(action string, commands []*cobra.Command) (*cobra.Command, error) {
	// We get all commands of the registered actions and only want to return the
	// command of the action we want to execute.
	var act *cobra.Command
	for _, c := range commands {
		if c.Name() == action {
			act = c
			break
		}
	}

	if act == nil {
		return nil, microerror.Maskf(commandNotFoundError, action)
	}

	// Once we found the action command we need to find its own execute command,
	// because the execute command wraps the business logic we want to execute
	// for the test plan.
	var exe *cobra.Command
	for _, c := range act.Commands() {
		if c.Name() == "execute" {
			exe = c
			break
		}
	}

	if exe == nil {
		return nil, microerror.Maskf(commandNotFoundError, "%s execute", action)
	}

	return exe, nil
}
