package ac001

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
)

func (e *Executor) execute(ctx context.Context) (v1alpha2.ClusterCRs, error) {
	var err error

	var clients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: e.logger,
		}

		clients, err = client.NewControlPlane(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}
	}

	crs, err := newCRs(clients.RESTConfig().Host)
	if err != nil {
		return v1alpha2.ClusterCRs{}, microerror.Mask(err)
	}

	{
		err = clients.CtrlClient().Create(ctx, crs.Cluster)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}

		err = clients.CtrlClient().Create(ctx, crs.AWSCluster)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}

		err = clients.CtrlClient().Create(ctx, crs.G8sControlPlane)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}

		err = clients.CtrlClient().Create(ctx, crs.AWSControlPlane)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}
	}

	return crs, nil
}
