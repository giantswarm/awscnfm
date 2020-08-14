package ac007

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/k8sclient/v3/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: e.logger,
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

			Scope: "cl001",
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}
	var nodeList *corev1.NodeList
	{
		nodeList, err = tcClients.K8sClient().CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var workerNode corev1.Node
	{
		for _, node := range nodeList.Items {
			_, ok := node.Labels[label.WorkerNodeRole]
			if ok {
				workerNode = node
				break
			}

		}
	}

	var workerPodList *corev1.PodList
	{
		workerPodList, err = tcClients.K8sClient().CoreV1().Pods("").List(metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + workerNode.Name,
		})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var currentWorkerPodsHostNetwork int
	var workerPods []string
	for _, pod := range workerPodList.Items {
		if pod.Spec.HostNetwork {
			currentWorkerPodsHostNetwork++
			workerPods = append(workerPods, pod.Name)
		}
	}

	expectedPods := "aws-node, calico-node, cert-exporter, kiam-agent, kube-proxy, node-exporter"
	if currentWorkerPodsHostNetwork != 6 {
		executionFailedError.Desc = fmt.Sprintf(
			"The tenant cluster defines 6 pods (%s) with host network on a worker node but it has currently %d pods with host network running.\nFound pods:\n%v",
			expectedPods,
			currentWorkerPodsHostNetwork,
			strings.Join(workerPods, "\n"),
		)

		return microerror.Mask(executionFailedError)
	}

	return nil
}
