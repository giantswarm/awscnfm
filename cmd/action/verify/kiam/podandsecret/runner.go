package podandsecret

import (
	"context"
	"fmt"

	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
)

const (
	kubeSystemNamespace = "kube-system"

	kiamServerLabelSelector = "app=kiam,component=kiam-server"
	kiamAgentLabelSelector  = "app=kiam,component=kiam-agent"

	masterNodeLabelSelector = "kubernetes.io/role=master"
	workerNodeLabelSelector = "kubernetes.io/role=worker"
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
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = r.checkTLSCerts(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.checkKiamPods(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) checkTLSCerts(ctx context.Context, tcClient kubernetes.Interface) error {
	for _, secret := range kiamTlSCertSecretNames {
		_, err := tcClient.CoreV1().Secrets(kubeSystemNamespace).Get(ctx, secret, metav1.GetOptions{})

		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

// checkKiamPods ensures kiam-agent and kiam-server pods are alive and running
func (r *runner) checkKiamPods(ctx context.Context, tcClient kubernetes.Interface) error {
	// count expected kiam-server and kiam-agent pods
	var expectedKiamServerPodCount, expectedKiamAgentPodCount int
	{
		masterNodes, err := tcClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{LabelSelector: masterNodeLabelSelector})
		if err != nil {
			return microerror.Mask(err)
		}
		expectedKiamServerPodCount = len(masterNodes.Items)

		workerNodes, err := tcClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{LabelSelector: workerNodeLabelSelector})
		if err != nil {
			return microerror.Mask(err)
		}
		expectedKiamAgentPodCount = len(workerNodes.Items)
	}

	// kiam server
	{
		kiamServerPods, err := tcClient.CoreV1().Pods(kubeSystemNamespace).List(ctx, metav1.ListOptions{LabelSelector: kiamServerLabelSelector})
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
		kiamAgentPods, err := tcClient.CoreV1().Pods(kubeSystemNamespace).List(ctx, metav1.ListOptions{LabelSelector: kiamAgentLabelSelector})
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
