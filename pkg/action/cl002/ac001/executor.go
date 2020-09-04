package ac001

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
)

func (e *Executor) execute(ctx context.Context) (v1alpha2.ClusterCRs, error) {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: e.logger,
		}

		cpClients, err = client.NewControlPlane(c)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}
	}

	var releases []v1alpha1.Release
	{
		var list v1alpha1.ReleaseList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
		)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}

		releases = list.Items
	}

	crs, err := newCRs(releases, cpClients.RESTConfig().Host)
	if err != nil {
		return v1alpha2.ClusterCRs{}, microerror.Mask(err)
	}

	{
		err = cpClients.CtrlClient().Create(ctx, crs.Cluster)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}

		err = cpClients.CtrlClient().Create(ctx, crs.AWSCluster)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}

		err = cpClients.CtrlClient().Create(ctx, crs.G8sControlPlane)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}

		err = cpClients.CtrlClient().Create(ctx, crs.AWSControlPlane)
		if err != nil {
			return v1alpha2.ClusterCRs{}, microerror.Mask(err)
		}
	}

	return crs, nil
}
