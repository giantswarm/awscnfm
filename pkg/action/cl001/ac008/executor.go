package ac008

import (
	"context"
	"fmt"

	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
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

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: e.logger,

			KubeConfig: env.KubeConfig(),
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
			Logger:       e.logger,

			Scope: e.scope,
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	err = e.checkTLSCerts(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	err = e.checkKiamPods(ctx, tcClients.K8sClient())
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// checkTLSCerts Ensures that kiam  tls certs are created.
var kiamTlSCertSecretNames = []string{"kiam-agent-tls", "kiam-ca-tls", "kiam-server-tls"}

func (e *Executor) checkTLSCerts(ctx context.Context, tcClient kubernetes.Interface) error {
	for _, secret := range kiamTlSCertSecretNames {
		_, err := tcClient.CoreV1().Secrets(kubeSystemNamespace).Get(ctx, secret, metav1.GetOptions{})

		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

// checkKiamPods ensures kiam-agent and kiam-server pods are alive and running
func (e *Executor) checkKiamPods(ctx context.Context, tcClient kubernetes.Interface) error {
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
			executionFailedError.Desc = fmt.Sprintf("Execution failed: wrong kiam-server pod count, expected %d but got %d", expectedKiamServerPodCount, len(kiamServerPods.Items))
			return microerror.Mask(executionFailedError)
		}

		for _, kiamServerPod := range kiamServerPods.Items {
			if kiamServerPod.Status.Phase != apiv1.PodRunning {
				executionFailedError.Desc = fmt.Sprintf("Execution failed: pod %s in namespace %s is not running.", kiamServerPod.Name, kiamServerPod.Namespace)
				return microerror.Mask(executionFailedError)
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
			executionFailedError.Desc = fmt.Sprintf("Execution failed: wrong kiam-agent pod count, expected %d but got %d", expectedKiamAgentPodCount, len(kiamAgentPods.Items))
			return microerror.Mask(executionFailedError)
		}

		for _, kiamAgentPod := range kiamAgentPods.Items {
			if kiamAgentPod.Status.Phase != apiv1.PodRunning {
				executionFailedError.Desc = fmt.Sprintf("Execution failed: pod %s in namespace %s is not running.", kiamAgentPod.Name, kiamAgentPod.Namespace)
				return microerror.Mask(executionFailedError)
			}
		}
	}

	return nil
}
