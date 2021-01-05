package defaultnetpol

import (
	"context"

	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/key"
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

	dockerRegistry, err := key.FetchDockerRegistry(ctx, cpClients.CtrlClient())
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.createTestNamespace(ctx, tcClients.CtrlClient())
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.createDefaultNetpol(ctx, tcClients.CtrlClient())
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.createTestPod(ctx, tcClients.CtrlClient(), dockerRegistry)
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.createTestSvc(ctx, tcClients.CtrlClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// createTestNamespace will create a test namespace
func (r *runner) createTestNamespace(ctx context.Context, ctrlClient k8sruntimeclient.Client) error {
	ns := netPolTestNamespace()

	err := ctrlClient.Create(ctx, ns)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// createDefaultNetpol will create a default network policy 'deny-from-all' in the test namespace
func (r *runner) createDefaultNetpol(ctx context.Context, ctrlClient k8sruntimeclient.Client) error {
	networkPolicy := defaultNetworkPolicy()

	err := ctrlClient.Create(ctx, networkPolicy)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// createTestPod will create a pod running and exposing nginx in the test namespace
// the pod will be used to test the network policy
func (r *runner) createTestPod(ctx context.Context, ctrlClient k8sruntimeclient.Client, dockerRegistry string) error {
	pod := nginxTestPod(dockerRegistry)

	err := ctrlClient.Create(ctx, pod)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// createTestSvc will create a service pointing to the nginx test pod
func (r *runner) createTestSvc(ctx context.Context, ctrlClient k8sruntimeclient.Client) error {
	svc := nginxTestPodService()

	err := ctrlClient.Create(ctx, svc)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
