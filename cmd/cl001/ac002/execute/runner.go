package execute

import (
	"context"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/pkg/action/ac002"
	"github.com/giantswarm/awscnfm/pkg/client"
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

	var clients *client.Client
	{
		c := client.Config{
			Logger: r.logger,

			KubeConfig:    env.KubeConfig(),
			TenantCluster: env.TenantCluster(),
		}

		clients, err = client.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = ac002.Execute(ctx, clients)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
