package ac005

import (
	"context"

	"github.com/giantswarm/microerror"
)

func (e *Executor) execute(ctx context.Context) error {
	crs, err := newCRs(ctx, e.tenantCluster, e.clients.ControlPlane.RESTConfig().Host)
	if err != nil {
		return microerror.Mask(err)
	}

	{
		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.MachineDeployment)
		if err != nil {
			return microerror.Mask(err)
		}

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, crs.AWSMachineDeployment)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
