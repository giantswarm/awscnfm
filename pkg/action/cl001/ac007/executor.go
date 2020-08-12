package ac007

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

func (e *Executor) execute(ctx context.Context) error {
	var err error
	var nodeList corev1.NodeList
	{
		err := e.clients.TenantCluster.CtrlClient().List(
			ctx,
			&nodeList,
		)
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
		workerPodList, err = e.clients.TenantCluster.K8sClient().CoreV1().Pods("").List(metav1.ListOptions{
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

	if currentWorkerPodsHostNetwork != 6 {
		executionFailedError.Desc = fmt.Sprintf(
			"The tenant cluster defines 6 pods with host network on a worker node but it has currently %d pods with host network running.\nFound pods:\n%v",
			currentWorkerPodsHostNetwork,
			strings.Join(workerPods, ","),
		)

		return microerror.Mask(executionFailedError)
	}

	return nil
}
