package ha

import (
	"context"
	"fmt"

	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	corev1 "k8s.io/api/core/v1"

	"github.com/giantswarm/awscnfm/v14/pkg/client"
	"github.com/giantswarm/awscnfm/v14/pkg/env"
	"github.com/giantswarm/awscnfm/v14/pkg/label"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: r.logger,

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
			ControlPlane:  cpClients,
			Logger:        r.logger,
			TenantCluster: r.flag.TenantCluster,
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
