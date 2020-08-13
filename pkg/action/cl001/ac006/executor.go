package ac006

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

	var masterNode corev1.Node
	{
		for _, node := range nodeList.Items {
			_, ok := node.Labels[label.MasterNodeRole]
			if ok {
				masterNode = node
				break
			}

		}
	}

	var masterPodList *corev1.PodList
	{
		masterPodList, err = tcClients.K8sClient().CoreV1().Pods("").List(metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + masterNode.Name,
		})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var currentMasterPodsHostNetwork int
	var masterPods []string
	for _, pod := range masterPodList.Items {
		if pod.Spec.HostNetwork {
			currentMasterPodsHostNetwork++
			masterPods = append(masterPods, pod.Name)
		}
	}

	if currentMasterPodsHostNetwork != 9 {
		executionFailedError.Desc = fmt.Sprintf(
			"The tenant cluster defines 9 pods with host network on a master node but it has currently %d pods with host network running.\nFound pods:\n%v",
			currentMasterPodsHostNetwork,
			strings.Join(masterPods, "\n"),
		)

		return microerror.Mask(executionFailedError)
	}

	return nil
}
