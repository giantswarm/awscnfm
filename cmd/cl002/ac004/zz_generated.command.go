package ac004

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/action"
	"github.com/giantswarm/awscnfm/v12/pkg/action/cl002/ac004"
	"github.com/giantswarm/awscnfm/v12/pkg/config"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
)

const (
	name        = "ac004"
	description = "Execute action ac004 for cluster cl002."
)

type Config struct {
	Logger micrologger.Logger
	Stderr io.Writer
	Stdout io.Writer
}

func New(config Config) (*cobra.Command, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Stderr == nil {
		config.Stderr = os.Stderr
	}
	if config.Stdout == nil {
		config.Stdout = os.Stdout
	}

	f := &flag{}

	r := &runner{
		flag:   f,
		logger: config.Logger,
		stderr: config.Stderr,
		stdout: config.Stdout,
	}

	c := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  mustLong(),
		RunE:  r.Run,
	}

	f.Init(c)

	return c, nil
}

func mustLong() string {
	ctx := context.Background()

	var err error

	var e action.Explainer
	{
		c := ac004.ExplainerConfig{
			Scope:         "cl002",
			TenantCluster: config.Cluster("cl002", env.TenantCluster()),
		}

		e, err = ac004.NewExplainer(c)
		if err != nil {
			panic(err)
		}
	}

	s, err := e.Explain(ctx)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(s)
}
