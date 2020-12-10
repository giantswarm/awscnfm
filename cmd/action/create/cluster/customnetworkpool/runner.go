package customnetworkpool

import (
	"context"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
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

	crs, npcrs, err := r.newCRs(releases, cpClients.RESTConfig().Host)
	if err != nil {
		return microerror.Mask(err)
	}

	{
		err = cpClients.CtrlClient().Create(ctx, npcrs.NetworkPool)
		if err != nil {
			return microerror.Mask(err)
		}

		err = cpClients.CtrlClient().Create(ctx, crs.Cluster)
		if err != nil {
			return microerror.Mask(err)
		}

		err = cpClients.CtrlClient().Create(ctx, crs.AWSCluster)
		if err != nil {
			return microerror.Mask(err)
		}

		err = cpClients.CtrlClient().Create(ctx, crs.G8sControlPlane)
		if err != nil {
			return microerror.Mask(err)
		}

		err = cpClients.CtrlClient().Create(ctx, crs.AWSControlPlane)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
