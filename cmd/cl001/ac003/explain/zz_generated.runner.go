package explain

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/action"
	"github.com/giantswarm/awscnfm/v12/pkg/action/cl001/ac003"
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

	var e action.Explainer
	{
		c := ac003.ExplainerConfig{
			TenantCluster: config.Cluster("cl001", env.TenantCluster()),
		}

		e, err = ac003.NewExplainer(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	s, err := e.Explain(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	fmt.Println(strings.TrimSpace(s))
	fmt.Println()

	return nil
}
