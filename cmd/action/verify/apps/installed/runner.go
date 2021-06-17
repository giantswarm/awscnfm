package installed

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	applicationv1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/application/v1alpha1"
	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
}

var expectedApps = []string{"aws-ebs-csi-driver", "cert-exporter", "cert-manager", "chart-operator", "cluster-autoscaler", "coredns", "etcd-cluster-migrator", "external-dns", "kiam", "kube-state-metrics", "metrics-server", "net-exporter", "node-exporter"}

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

	var applications []applicationv1alpha1.Chart
	{
		var list applicationv1alpha1.ChartList
		err := tcClients.CtrlClient().List(
			ctx,
			&list,
			k8sruntimeclient.InNamespace("giantswarm"),
		)
		if err != nil {
			return microerror.Mask(err)
		}

		applications = list.Items
	}

	// Contains apps present in TC
	var installedCharts []string
	for _, chart := range applications {
		installedCharts = append(installedCharts, chart.Spec.Name)
	}

	if !reflect.DeepEqual(installedCharts, expectedApps) {
		appsNotInstalled.Desc = fmt.Sprintf(
			"The Tenant Cluster defines %d charts (%s) but it has currently %d charts (%s)",
			len(expectedApps),
			strings.Join(expectedApps, ", "),
			len(installedCharts),
			strings.Join(installedCharts, ", "),
		)

		return microerror.Mask(appsNotInstalled)
	}

	for _, chart := range applications {
		if chart.Status.Release.Status != "deployed" {
			appsNotInstalled.Desc = fmt.Sprintf(
				"The application (%s) is in status (%s), should be deployed.",
				chart.Spec.Name,
				chart.Status.Release.Status,
			)

			return microerror.Mask(appsNotInstalled)
		}
	}

	return nil
}
