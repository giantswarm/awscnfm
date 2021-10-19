package updated

import (
	"context"

	"github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/types"

	"github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/key"
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

	var cl v1alpha3.AWSCluster
	{
		err = cpClients.CtrlClient().Get(
			ctx,
			types.NamespacedName{Name: r.flag.TenantCluster, Namespace: key.OrganizationNamespaceFromName(key.Organization)},
			&cl,
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	if cl.Status.Cluster.LatestCondition() == v1alpha3.ClusterStatusConditionUpdated {
		return nil
	}

	return microerror.Maskf(wrongClusterStatusConditionError, cl.Status.Cluster.LatestCondition())
}
