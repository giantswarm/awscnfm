package running

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/apps/v1"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
}

var expectedDeployments = []string{
	"cert-manager-cainjector",
	"cert-manager-controller",
	"cert-manager-webhook",
	"cluster-autoscaler",
	"coredns",
	"ebs-csi-controller",
	"external-dns",
	"kube-state-metrics",
	"metrics-server",
}

var expectedDaemonSets = []string{
	"aws-node",
	"calico-node",
	"cert-exporter",
	"ebs-csi-node",
	"kiam-agent",
	"kiam-server",
	"kube-proxy",
	"net-exporter",
	"node-exporter",
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
			ControlPlane: cpClients,
			Logger:       r.logger,

			TenantCluster: r.flag.TenantCluster,
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var deployments []v1.Deployment
	{
		var list v1.DeploymentList
		err := tcClients.CtrlClient().List(
			ctx,
			&list,
			k8sruntimeclient.InNamespace("kube-system"),
		)
		if err != nil {
			return microerror.Mask(err)
		}

		deployments = list.Items
	}

	var daemonSets []v1.DaemonSet
	{
		var list v1.DaemonSetList
		err := tcClients.CtrlClient().List(
			ctx,
			&list,
			k8sruntimeclient.InNamespace("kube-system"),
		)
		if err != nil {
			return microerror.Mask(err)
		}

		daemonSets = list.Items
	}

	var failedDaemonSets []string
	for _, expectedDaemonSet := range expectedDaemonSets {
		for _, ds := range daemonSets {
			if expectedDaemonSet == ds.Name {
				r.logger.LogCtx(ctx, "level", "info", "message", fmt.Sprintf("comparing daemonset %s, expected replicas %v, current replicas %v",
					expectedDaemonSet,
					ds.Status.DesiredNumberScheduled,
					ds.Status.CurrentNumberScheduled,
				))
				if ds.Status.DesiredNumberScheduled != ds.Status.CurrentNumberScheduled {
					failedDaemonSets = append(failedDaemonSets, ds.Name)
				}
			}
		}
	}

	var failedDeployments []string
	for _, expectedDeployment := range expectedDeployments {
		for _, deployment := range deployments {
			if expectedDeployment == deployment.Name {
				r.logger.LogCtx(ctx, "level", "info", "message", fmt.Sprintf("comparing deployment %s, expected replicas %v, current replicas %v",
					expectedDeployment,
					deployment.Status.Replicas,
					deployment.Status.ReadyReplicas,
				))
				if deployment.Status.Replicas != deployment.Status.ReadyReplicas {
					failedDaemonSets = append(failedDeployments, deployment.Name)
				}
			}
		}
	}

	if len(failedDaemonSets) > 0 {
		appsNotRunning.Desc = fmt.Sprintf(
			"The Tenant Cluster failed to ensure %d daemonsets (%s)",
			len(failedDaemonSets),
			strings.Join(failedDaemonSets, ", "),
		)

		return microerror.Mask(appsNotRunning)
	}

	if len(failedDeployments) > 0 {
		appsNotRunning.Desc = fmt.Sprintf(
			"The Tenant Cluster failed to ensure %d deployments (%s)",
			len(failedDeployments),
			strings.Join(failedDeployments, ", "),
		)

		return microerror.Mask(appsNotRunning)
	}

	return nil
}
