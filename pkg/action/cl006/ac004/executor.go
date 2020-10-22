package ac004

import (
	"context"
	"fmt"

	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"

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

	var tcClients k8sclient.Interface
	{
		c := client.TenantClusterConfig{
			ControlPlane: cpClients,
			Logger:       e.logger,

			KubeConfig: env.TenantClusterKubeConfig(),
			Scope:      e.scope,
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var list corev1.NodeList
	{
		err := tcClients.CtrlClient().List(
			ctx,
			&list,
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var masterNodesReady int
	for _, node := range list.Items {
		_, ok := node.Labels[label.MasterNodeRole]
		if !ok {
			continue
		}

		for _, c := range node.Status.Conditions {
			if c.Type == corev1.NodeReady && c.Status == corev1.ConditionTrue {
				masterNodesReady++
			}
		}
	}

	if masterNodesReady != 3 {
		return microerror.Maskf(wrongMasterNodesError, fmt.Sprintf("%d/3 healthy master nodes running.", masterNodesReady))
	}

	return nil
}
