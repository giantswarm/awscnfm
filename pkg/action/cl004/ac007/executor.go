package ac007

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/env"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

// expectedPods are all host network pods which we expect to run on a worker node
var expectedPods = []string{"aws-node", "calico-node", "cert-exporter", "kiam-agent", "kube-proxy", "node-exporter"}

func (e *Executor) execute(ctx context.Context) error {
	var err error

	var cpClients k8sclient.Interface
	{
		c := client.ControlPlaneConfig{
			Logger: e.logger,

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
			Logger:       e.logger,
		}

		tcClients, err = client.NewTenantCluster(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}
	var nodeList *corev1.NodeList
	{
		nodeList, err = tcClients.K8sClient().CoreV1().Nodes().List(ctx, metav1.ListOptions{})
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
		workerPodList, err = tcClients.K8sClient().CoreV1().Pods("").List(ctx, metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + workerNode.Name,
		})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var workerPods []string
	for _, pod := range workerPodList.Items {
		if pod.Spec.HostNetwork {
			workerPods = append(workerPods, pod.Name)
		}
	}

	if len(workerPods) != len(expectedPods) {
		executionFailedError.Desc = fmt.Sprintf(
			"The Tenant Cluster defines %d pods (%s) but it has currently %d pods (%s) with host network running",
			len(expectedPods),
			strings.Join(expectedPods, ", "),
			len(workerPods),
			strings.Join(workerPods, ", "),
		)

		return microerror.Mask(executionFailedError)
	}

	return nil
}
