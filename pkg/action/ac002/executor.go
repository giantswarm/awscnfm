package ac002

import (
	"context"
	"fmt"

	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/awscnfm/pkg/action"
	"github.com/giantswarm/awscnfm/pkg/label"
)

type ExecutorConfig struct {
	Clients *action.Clients
	Logger  micrologger.Logger

	TenantCluster string
}

type Executor struct {
	clients *action.Clients
	logger  micrologger.Logger

	tenantCluster string
}

func NewExecutor(config ExecutorConfig) (*Executor, error) {
	if config.Clients == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Clients must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.TenantCluster == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.TenantCluster must not be empty", config)
	}

	e := &Executor{
		clients: config.Clients,
		logger:  config.Logger,

		tenantCluster: config.TenantCluster,
	}

	return e, nil
}

func (e *Executor) Execute(ctx context.Context) error {
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

		if len(list.Items) == 0 {
			return microerror.Mask(notFoundError)
		}
		if len(list.Items) > 1 {
			return microerror.Mask(tooManyCRsError)
		}

		desiredNodesReady = list.Items[0].Spec.Replicas
	}

	if currentNodesReady != desiredNodesReady {
		executionFailedError.Desc = fmt.Sprintf(
			"The Tenant Cluster defines %d master nodes but it only has %d/%d healthy master nodes running.",
			desiredNodesReady,
			currentNodesReady,
			desiredNodesReady,
		)

		return microerror.Mask(executionFailedError)
	}

	return nil
}
