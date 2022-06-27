package version

import (
	"context"

	"github.com/giantswarm/k8sclient/v7/pkg/k8sclient"
	"github.com/giantswarm/k8smetadata/pkg/label"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/release-operator/v3/api/v1alpha1"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	apiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"

	"github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/release"
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

	var re *release.Release
	{
		c := release.Config{
			FromEnv:  r.flag.ReleaseVersion,
			Releases: releases,
		}

		re, err = release.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var cl apiv1alpha3.Cluster
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

	cl.Labels[label.ClusterOperatorVersion] = re.Components()["cluster-operator"]
	cl.Labels[label.ReleaseVersion] = re.Version()

	{
		err = cpClients.CtrlClient().Update(ctx, &cl)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
