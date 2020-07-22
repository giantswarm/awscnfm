package ac001

import (
	"context"

	"github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/pkg/key"
	"github.com/giantswarm/awscnfm/pkg/release"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var releaseComponents map[string]string
	{
		c := release.Config{}

		releaseCollection, err := release.New(c)
		if err != nil {
			return microerror.Mask(err)
		}

		releaseComponents = releaseCollection.ReleaseComponents(key.Release)
	}

	var crs v1alpha2.ClusterCRs
	{
		c := v1alpha2.ClusterCRsConfig{
			ClusterID:         "cl001",
			Domain:            key.DomainFromHost(e.clients.ControlPlane.RESTConfig().Host),
			Description:       "awscnfm cluster cl001",
			Owner:             "giantswarm",
			Region:            key.RegionFromHost(e.clients.ControlPlane.RESTConfig().Host),
			ReleaseComponents: releaseComponents,
			ReleaseVersion:    key.Release,
		}

		crs, err = v1alpha2.NewClusterCRs(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.Cluster)
		if err != nil {
			return microerror.Mask(err)
		}

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.AWSCluster)
		if err != nil {
			return microerror.Mask(err)
		}

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.G8sControlPlane)
		if err != nil {
			return microerror.Mask(err)
		}

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.AWSControlPlane)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
