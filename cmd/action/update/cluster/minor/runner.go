package minor

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/apiextensions/v2/pkg/label"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
	"github.com/giantswarm/awscnfm/v12/pkg/release"
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

	var m *release.Minor
	{
		c := release.MinorConfig{
			FromEnv:     env.ReleaseVersion(),
			FromProject: project.Version(),
			Releases:    releases,
		}

		m, err = release.NewMinor(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var cl apiv1alpha2.Cluster
	{
		err = cpClients.CtrlClient().Get(
			ctx,
			types.NamespacedName{Name: r.flag.TenantCluster, Namespace: v1.NamespaceDefault},
			&cl,
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	cl.Labels[label.ClusterOperatorVersion] = m.Components().Latest()["cluster-operator"]
	cl.Labels[label.ReleaseVersion] = m.Version().Latest()

	{
		err = cpClients.CtrlClient().Update(ctx, &cl)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
