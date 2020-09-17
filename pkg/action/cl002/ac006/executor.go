package ac006

import (
	"context"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pkgclient "github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := pkgclient.ControlPlaneConfig{
			Logger: e.logger,

			KubeConfig: env.KubeConfig(),
		}

		cpClients, err = pkgclient.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		var list apiv1alpha2.ClusterList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "Cluster")
		}
	}

	{
		var list infrastructurev1alpha2.AWSClusterList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "AWSCluster")
		}
	}

	{
		var list infrastructurev1alpha2.G8sControlPlaneList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "G8sControlPlane")
		}
	}

	{
		var list infrastructurev1alpha2.AWSControlPlaneList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "AWSControlPlane")
		}
	}

	{
		var list apiv1alpha2.MachineDeploymentList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(list.Items) != 0 {
			return microerror.Maskf(customResourceCleanupError, "MachineDeployment")
		}
	}

	{
		var list infrastructurev1alpha2.AWSMachineDeploymentList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
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
