package ac006

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v12/pkg/client"
	"github.com/giantswarm/awscnfm/v12/pkg/label"
)

// expectedPods are all host network pods which we expect to run on a master node
var expectedPods = "aws-node, calico-node, cert-exporter, k8s-api-healthz, k8s-api-server, k8s-controller-manager, k8s-scheduler, kube-proxy, node-exporter"

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
		nodeList, err = tcClients.K8sClient().CoreV1().Nodes().List(ctx, metav1.ListOptions{})
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
		masterPodList, err = tcClients.K8sClient().CoreV1().Pods("").List(ctx, metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + masterNode.Name,
		})
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var masterPods []string
	for _, pod := range masterPodList.Items {
		// this cronjob pod currently counts against the pods with hostnetwork on master nodes
		// it will be dropped with the fix for aws-cni v1.7.0, see issue https://github.com/giantswarm/giantswarm/issues/11077
		if strings.Contains(pod.Name, "aws-cni-restarter") {
			continue
		}
		if pod.Spec.HostNetwork {
			masterPods = append(masterPods, pod.Name)
		}
	}

	if len(masterPods) != len(expectedPods) {
		executionFailedError.Desc = fmt.Sprintf(
			"The Tenant Cluster defines %d pods (%s) but it has currently %d pods (%s) with host network running",
			len(expectedPods),
			expectedPods,
			len(masterPods),
			strings.Join(masterPods, ", "),
		)

		return microerror.Mask(executionFailedError)
	}

	return nil
}
