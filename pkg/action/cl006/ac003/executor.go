package ac003

import (
	"context"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"

	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: e.logger,

			KubeConfig: env.ControlPlaneKubeConfig(),
		}

		cpClients, err = client.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var cpl infrastructurev1alpha2.G8sControlPlaneList
	{
		err = cpClients.CtrlClient().List(
			ctx,
			&cpl,
			pkgclient.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var cp infrastructurev1alpha2.G8sControlPlane
	{
		cp = cpl.Items[0]
		cp.Spec.Replicas = 3
	}

	{
		err = cpClients.CtrlClient().Update(ctx, &cp)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
