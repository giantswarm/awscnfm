package ac005

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/config"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
)

func (e *Executor) execute(ctx context.Context) error {
	scope := "cl001"
	id := config.Cluster(scope, env.TenantCluster())

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

	var releases []v1alpha1.Release
	{
		var list v1alpha1.ReleaseList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
		)
		if err != nil {
			return microerror.Mask(err)
		}

		releases = list.Items
	}

	crs, err := newCRs(ctx, releases, id, cpClients.RESTConfig().Host)
	if err != nil {
		return microerror.Mask(err)
	}

	{
		err = cpClients.CtrlClient().Create(ctx, crs.MachineDeployment)
		if err != nil {
			return microerror.Mask(err)
		}

		err = cpClients.CtrlClient().Create(ctx, crs.AWSMachineDeployment)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
