package ac001

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
)

func (e *Executor) execute(ctx context.Context) error {
	fmt.Println(e.clients.ControlPlane.RESTConfig().Host)
	fmt.Println("FOOOO")
	crs, npCrs, err := newCRs(e.clients.ControlPlane.RESTConfig().Host)
	if err != nil {
		return microerror.Mask(err)
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

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, npCrs.MachineDeployment)
		if err != nil {
			return microerror.Mask(err)
		}

		err = e.clients.ControlPlane.CtrlClient().Create(ctx, npCrs.AWSMachineDeployment)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
