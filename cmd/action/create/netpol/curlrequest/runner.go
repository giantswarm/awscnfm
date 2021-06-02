package curlrequest

import (
	"context"

	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v14/pkg/client"
	"github.com/giantswarm/awscnfm/v14/pkg/env"
	"github.com/giantswarm/awscnfm/v14/pkg/key"
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

	err = r.createNetPolTestJobs(ctx, tcClients.CtrlClient(), dockerRegistry)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// createNetPolTestJob will spawn 2 jobs to test the network policy
// one job in namespace `test` to ensure that the traffic from pods inside of the namespace `test` is working
// one job in namespace `default` to ensure that the traffic from from pods in other namespace is blocked by the network policy
func (r *runner) createNetPolTestJobs(ctx context.Context, tcClient k8sruntimeclient.Client, dockerRegistry string) error {
	jobSuccess := testNetworkPolicyJob(dockerRegistry, key.NetPolTestNamespaceName)
	err := tcClient.Create(ctx, jobSuccess)
	if err != nil {
		return microerror.Mask(err)
	}

	jobFailure := testNetworkPolicyJob(dockerRegistry, key.NetPolDefaultNamespaceName)
	err = tcClient.Create(ctx, jobFailure)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
