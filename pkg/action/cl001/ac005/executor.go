package ac005

import (
	"context"

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

	crs, err := newCRs(ctx, id, cpClients.RESTConfig().Host)
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
