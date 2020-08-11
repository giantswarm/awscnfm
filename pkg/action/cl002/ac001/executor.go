package ac001

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/microerror"
)

func (e *Executor) execute(ctx context.Context) (v1alpha2.ClusterCRs, v1alpha2.NodePoolCRs, error) {
	crs, npCRs, err := newCRs(e.clients.ControlPlane.RESTConfig().Host)
	if err != nil {
		return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
	}

	{
		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.Cluster)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.AWSCluster)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.G8sControlPlane)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.AWSControlPlane)
		if err != nil {
			return v1alpha2.ClusterCRs{}, v1alpha2.NodePoolCRs{}, microerror.Mask(err)
		}
	}

	return crs, npCRs, nil
}
