package podandsecret

import (
	"context"
	"fmt"

	"github.com/giantswarm/k8sclient/v5/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

const (
	kubeSystemNamespace = "kube-system"

	kiamApp             = "kiam"
	componentKiamAgent  = "kiam-agent"
	componentKiamServer = "kiam-server"
)

// checkTLSCerts Ensures that kiam  tls certs are created.
var kiamTlSCertSecretNames = []string{"kiam-agent-tls", "kiam-ca-tls", "kiam-server-tls"}

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

	err = r.checkTLSCerts(ctx, tcClients.CtrlClient())
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.checkKiamPods(ctx, tcClients.CtrlClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) checkTLSCerts(ctx context.Context, tcClient k8sruntimeclient.Client) error {
	for _, secret := range kiamTlSCertSecretNames {

		s := &apiv1.Secret{}

		err := tcClient.Get(ctx,
			k8sruntimeclient.ObjectKey{
				Namespace: kubeSystemNamespace,
				Name:      secret,
			},
			s)

		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

// checkKiamPods ensures kiam-agent and kiam-server pods are alive and running
func (r *runner) checkKiamPods(ctx context.Context, tcClient k8sruntimeclient.Client) error {
	// count expected kiam-server and kiam-agent pods
	var expectedKiamServerPodCount, expectedKiamAgentPodCount int
	{
		masterNodes := &apiv1.NodeList{}

		err := tcClient.List(
			ctx,
			masterNodes,
			k8sruntimeclient.MatchingLabels{label.MasterNodeRole: ""})
		if err != nil {
			return microerror.Mask(err)
		}
		expectedKiamServerPodCount = len(masterNodes.Items)

		workerNodes := &apiv1.NodeList{}

		err = tcClient.List(
			ctx,
			workerNodes,
			k8sruntimeclient.MatchingLabels{label.WorkerNodeRole: ""})
		if err != nil {
			return microerror.Mask(err)
		}
		expectedKiamAgentPodCount = len(workerNodes.Items)
	}

	// kiam server
	{
		kiamServerPods := &apiv1.PodList{}

		err := tcClient.List(
			ctx,
			kiamServerPods,
			k8sruntimeclient.InNamespace(kubeSystemNamespace),
			k8sruntimeclient.MatchingLabels{
				label.App:       kiamApp,
				label.Component: componentKiamServer,
			})
		if err != nil {
			return microerror.Mask(err)
		}

		if len(kiamServerPods.Items) != expectedKiamServerPodCount {
			return microerror.Maskf(executionFailedError, fmt.Sprintf("wrong kiam-server pod count, expected %d but got %d", expectedKiamServerPodCount, len(kiamServerPods.Items)))
		}

		for _, kiamServerPod := range kiamServerPods.Items {
			if kiamServerPod.Status.Phase != apiv1.PodRunning {
				return microerror.Maskf(executionFailedError, fmt.Sprintf("pod %s in namespace %s is not running.", kiamServerPod.Name, kiamServerPod.Namespace))
			}
		}
	}

	// kiam agent
	{
		kiamAgentPods := &apiv1.PodList{}

		err := tcClient.List(
			ctx,
			kiamAgentPods,
			k8sruntimeclient.InNamespace(kubeSystemNamespace),
			k8sruntimeclient.MatchingLabels{
				label.App:       kiamApp,
				label.Component: componentKiamAgent,
			})

		if err != nil {
			return microerror.Mask(err)
		}

		if len(kiamAgentPods.Items) != expectedKiamAgentPodCount {
			return microerror.Maskf(executionFailedError, fmt.Sprintf("wrong kiam-agent pod count, expected %d but got %d", expectedKiamAgentPodCount, len(kiamAgentPods.Items)))
		}

		for _, kiamAgentPod := range kiamAgentPods.Items {
			if kiamAgentPod.Status.Phase != apiv1.PodRunning {
				return microerror.Maskf(executionFailedError, fmt.Sprintf("pod %s in namespace %s is not running.", kiamAgentPod.Name, kiamAgentPod.Namespace))
			}
		}
	}

	return nil
}
