package volume

import (
	"context"

	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/key"
)

const (
	kubeSystemNamespace = "kube-system"
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
			ControlPlane: cpClients,
			Logger:       r.logger,

			TenantCluster: r.flag.TenantCluster,
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	// dockerRegistry is needed in order to spawn pod with proper docker image that will execute aws-cli call
	var dockerRegistry string
	{
		dockerRegistry, err = key.FetchDockerRegistry(ctx, cpClients.CtrlClient())
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = r.createEBSVolume(ctx, tcClients.CtrlClient(), dockerRegistry)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// createAWSApiCallJob will spawn a job in k8s tenant cluster to test calling AWS API to ensure kiam works as expected
func (r *runner) createEBSVolume(ctx context.Context, tcClient k8sruntimeclient.Client, dockerRegistry string) error {
	pvClaim := ensurePersistentVolumeClaim()
	err := tcClient.Create(ctx, pvClaim)
	if err != nil {
		return microerror.Mask(err)
	}

	clusterRole := ensureClusterRole()
	err = tcClient.Create(ctx, clusterRole)
	if err != nil {
		return microerror.Mask(err)
	}

	roleBinding := ensureRoleBinding()
	err = tcClient.Create(ctx, roleBinding)
	if err != nil {
		return microerror.Mask(err)
	}

	job := ensureEBSVolume(dockerRegistry, r.flag.TenantCluster)
	err = tcClient.Create(ctx, job)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
