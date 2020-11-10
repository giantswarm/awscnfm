package hacontrolplane

import (
	"context"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/apiextensions/v2/pkg/label"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
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

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: r.logger,

			KubeConfig: env.ControlPlaneKubeConfig(),
		}

		cpClients, err = client.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var cpl infrastructurev1alpha2.G8sControlPlaneList
	{
		err = cpClients.CtrlClient().List(
			ctx,
			&cpl,
			pkgclient.MatchingLabels{label.Cluster: r.flag.TenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var cp infrastructurev1alpha2.G8sControlPlane
	{
		cp = cpl.Items[0]
		cp.Spec.Replicas = 3
	}

	{
		err = cpClients.CtrlClient().Update(ctx, &cp)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
