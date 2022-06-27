package hostnetworkpod

import (
	"context"
	"fmt"
	"strings"

	"github.com/giantswarm/k8sclient/v7/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/awscnfm/v15/pkg/client"
	"github.com/giantswarm/awscnfm/v15/pkg/env"
	"github.com/giantswarm/awscnfm/v15/pkg/label"
)

// expectedPods are all host network pods which we expect to run on a worker node
var expectedPods = []string{"aws-node", "calico-node", "ebs-csi-node", "kiam-agent", "kube-proxy", "node-exporter"}

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
