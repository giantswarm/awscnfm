package ac003

import (
	"context"
	"fmt"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/pkg/label"
)

func (e *Executor) execute(ctx context.Context) error {
	var list corev1.NodeList
	{
		err := e.clients.TenantCluster.CtrlClient().List(
			ctx,
			&list,
		)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	var currentNodesReady int
	for _, node := range list.Items {
		_, ok := node.Labels[label.MasterNodeRole]
		if !ok {
			continue
		}

		for _, c := range node.Status.Conditions {
			if c.Type == corev1.NodeReady && c.Status == corev1.ConditionTrue {
				currentNodesReady++
			}
		}
	}

	var desiredNodesReady int
	{
		var list infrastructurev1alpha2.G8sControlPlaneList
		err := e.clients.ControlPlane.CtrlClient().List(
			ctx,
			&list,
			client.MatchingLabels{label.Cluster: e.tenantCluster},
		)
		if err != nil {
			return microerror.Mask(err)
		}

		for _, cr := range list.Items {
			desiredNodesReady += cr.Spec.Replicas
		}
	}

	if currentNodesReady != desiredNodesReady {
		executionFailedError.Desc = fmt.Sprintf(
			"The Tenant Cluster defines %d master nodes but it has only %d/%d healthy master nodes running.",
			desiredNodesReady,
			currentNodesReady,
			desiredNodesReady,
		)

		return microerror.Mask(executionFailedError)
	}

	return nil
}
