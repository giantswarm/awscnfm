package ac009

import (
	"context"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/microerror"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/pkg/label"
)

func (e *Executor) execute(ctx context.Context) error {
	{
		var list apiv1alpha2.ClusterList
		err := e.clients.ControlPlane.CtrlClient().List(
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
		err := e.clients.ControlPlane.CtrlClient().List(
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
		err := e.clients.ControlPlane.CtrlClient().List(
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
		err := e.clients.ControlPlane.CtrlClient().List(
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
		err := e.clients.ControlPlane.CtrlClient().List(
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
		err := e.clients.ControlPlane.CtrlClient().List(
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
