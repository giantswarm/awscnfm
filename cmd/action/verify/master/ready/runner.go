package ready

import (
	"context"
	"fmt"

	infrastructurev1alpha3 "github.com/giantswarm/apiextensions/v6/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/k8sclient/v7/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	pkgclient "github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/label"
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
		c := pkgclient.ControlPlaneConfig{
			Logger: r.logger,

			KubeConfig: env.ControlPlaneKubeConfig(),
		}

		cpClients, err = pkgclient.NewControlPlane(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var tcClients k8sclient.Interface
	{
		c := pkgclient.TenantClusterConfig{
			ControlPlane: cpClients,
			Logger:       r.logger,

			TenantCluster: r.flag.TenantCluster,
		}

		tcClients, err = pkgclient.NewTenantCluster(c)
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

	var currentNodesReady int
	for _, node := range list.Items {
		if !label.IsMaster(node) {
			continue
		}

		for _, c := range node.Status.Conditions {
			if c.Type == corev1.NodeReady && c.Status == corev1.ConditionTrue {
				currentNodesReady++
			}
		}
	}

	var desiredNodesReady int
	{
		var list infrastructurev1alpha3.G8sControlPlaneList
		err := cpClients.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: r.flag.TenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		for _, cr := range list.Items {
			desiredNodesReady += cr.Spec.Replicas
		}
	}

	if currentNodesReady != desiredNodesReady {
		executionFailedError.Desc = fmt.Sprintf(
			"The Tenant Cluster defines %d master nodes but it has only %d/%d healthy master nodes running.",
			desiredNodesReady,
			currentNodesReady,
			desiredNodesReady,
		)

		return microerror.Mask(executionFailedError)
	}

	return nil
}
