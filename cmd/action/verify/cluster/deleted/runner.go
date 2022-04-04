package deleted

import (
	"context"

	infrastructurev1alpha3 "github.com/giantswarm/apiextensions/v6/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/k8sclient/v7/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	apiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pkgclient "github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/label"
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
		c := pkgclient.ControlPlaneConfig{
			Logger: r.logger,

			KubeConfig: env.ControlPlaneKubeConfig(),
		}

		cpClients, err = pkgclient.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		var list apiv1alpha3.ClusterList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: r.flag.TenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "Cluster")
		}
	}

	{
		var list infrastructurev1alpha3.AWSClusterList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: r.flag.TenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "AWSCluster")
		}
	}

	{
		var list infrastructurev1alpha3.G8sControlPlaneList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: r.flag.TenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "G8sControlPlane")
		}
	}

	{
		var list infrastructurev1alpha3.AWSControlPlaneList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: r.flag.TenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "AWSControlPlane")
		}
	}

	{
		var list apiv1alpha3.MachineDeploymentList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: r.flag.TenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "MachineDeployment")
		}
	}

	{
		var list infrastructurev1alpha3.AWSMachineDeploymentList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: r.flag.TenantCluster},
		)

		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "AWSMachineDeployment")
		}
	}

	return nil
}
