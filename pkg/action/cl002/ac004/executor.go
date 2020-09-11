package ac004

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: e.logger,
		}

		cpClients, err = client.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var cl v1alpha2.AWSCluster
	{
		err = cpClients.CtrlClient().Get(
			ctx,
			types.NamespacedName{Name: e.tenantCluster, Namespace: v1.NamespaceDefault},
			&cl,
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	if cl.Status.Cluster.LatestCondition() == v1alpha2.ClusterStatusConditionUpdated {
		return nil
	}

	return microerror.Maskf(wrongClusterStatusConditionError, cl.Status.Cluster.LatestCondition())
}
