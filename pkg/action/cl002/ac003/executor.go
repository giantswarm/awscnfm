package ac003

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/apiextensions/v2/pkg/label"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/project"
	"github.com/giantswarm/awscnfm/v12/pkg/release"
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

	var p *release.Patch
	{
		c := release.PatchConfig{
			FromEnv:     env.ReleaseVersion(),
			FromProject: project.Version(),
			Releases:    releases,
		}

		p, err = release.NewPatch(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var cl apiv1alpha2.Cluster
	{
		err = cpClients.CtrlClient().Get(
			ctx,
			types.NamespacedName{Name: e.tenantCluster, Namespace: v1.NamespaceDefault},
			&cl,
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	cl.Labels[label.ClusterOperatorVersion] = p.Components().Latest()["cluster-operator"]
	cl.Labels[label.ReleaseVersion] = p.Version().Latest()

	{
		err = cpClients.CtrlClient().Update(ctx, &cl)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
